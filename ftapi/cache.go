package ftapi

import (
	"reflect"
	"sync"
	"time"
)

type (
	cachedObject struct {
		object    interface{}
		updatedAt time.Time
	}
	objectCache struct {
		sync.RWMutex
		types map[string]map[int]cachedObject
	}
)

var (
	intraCache   = objectCache{types: make(map[string]map[int]cachedObject)}
	cacheTimeout = 30 * time.Minute
)

func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			now := time.Now()
			intraCache.Lock()
			for typeName, objects := range intraCache.types {
				for key, cached := range objects {
					if now.Sub(cached.updatedAt) >= cacheTimeout {
						delete(objects, key)
					}
				}
				if len(objects) == 0 {
					delete(intraCache.types, typeName)
				}
			}
			intraCache.Unlock()
		}
	}()
}

func (cache objectCache) delete(obj interface{}) (prev interface{}) {
	value := reflect.Indirect(reflect.ValueOf(obj))
	ID := value.FieldByName("ID").Interface().(int)
	intraCache.Lock()
	if objects, present := intraCache.types[value.Type().String()]; present {
		if p, present := objects[ID]; present {
			prev = p.object
			delete(objects, ID)
		}
	}
	intraCache.Unlock()
	return
}

func (cache objectCache) get(obj interface{}) (present bool) {
	var objects map[int]cachedObject
	var cached cachedObject
	value := reflect.Indirect(reflect.ValueOf(obj))
	ID := value.FieldByName("ID").Interface().(int)
	intraCache.RLock()
	if objects, present = intraCache.types[value.Type().String()]; present {
		cached, present = objects[ID]
	}
	intraCache.RUnlock()
	if present {
		// Load object from cache without overwriting RequestData
		dup := reflect.ValueOf(cached.object)
		for i := 1; i < value.NumField(); i++ {
			value.Field(i).Set(dup.Field(i))
		}
	}
	return
}

func (cache objectCache) put(obj interface{}) (prev interface{}) {
	value := reflect.Indirect(reflect.ValueOf(obj))
	// Copy object, without hanging on to RequestData
	dup := reflect.New(value.Type()).Elem()
	for i := 1; i < value.NumField(); i++ {
		dup.Field(i).Set(value.Field(i))
	}
	toCache := cachedObject{
		object:    dup.Interface(),
		updatedAt: time.Now(),
	}
	typeName := dup.Type().String()
	ID := dup.FieldByName("ID").Interface().(int)
	intraCache.Lock()
	objects, present := intraCache.types[typeName]
	if !present {
		objects = make(map[int]cachedObject)
		intraCache.types[typeName] = objects
	}
	if cached, present := objects[ID]; present {
		prev = cached.object
	}
	objects[ID] = toCache
	intraCache.Unlock()
	return
}

func SetCacheTimeout(minutes int) {
	cacheTimeout = time.Duration(minutes) * time.Minute
}
