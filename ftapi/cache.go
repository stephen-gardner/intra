package ftapi

import (
	"encoding/json"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

type (
	cachedObject struct {
		data      json.RawMessage
		updatedAt time.Time
	}
	objectCache struct {
		sync.RWMutex
		types   map[string]map[int]cachedObject
		timeout time.Duration
		enabled bool
	}
)

var intraCache = objectCache{
	types:   make(map[string]map[int]cachedObject),
	timeout: 30 * time.Minute,
	enabled: true,
}

func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			now := time.Now()
			intraCache.Lock()
			for typeName, objects := range intraCache.types {
				for key, cached := range objects {
					if now.Sub(cached.updatedAt) >= intraCache.timeout {
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
			prev = p.data
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
		_ = json.Unmarshal(cached.data, obj)
	}
	return
}

func (cache objectCache) put(obj interface{}) (prev interface{}) {
	data, _ := json.Marshal(obj)
	toCache := cachedObject{
		data:      data,
		updatedAt: time.Now(),
	}
	value := reflect.Indirect(reflect.ValueOf(obj))
	typeName := value.Type().String()
	ID := value.FieldByName("ID").Interface().(int)
	intraCache.Lock()
	objects, present := intraCache.types[typeName]
	if !present {
		objects = make(map[int]cachedObject)
		intraCache.types[typeName] = objects
	}
	if cached, present := objects[ID]; present {
		prev = cached.data
	}
	objects[ID] = toCache
	intraCache.Unlock()
	return
}

func (cache objectCache) isEnabled() bool {
	intraCache.RLock()
	enabled := intraCache.enabled
	intraCache.RUnlock()
	return enabled
}

// Puts object into cache if cache writes are enabled in RequestData
func CacheObject(obj interface{}) (prev interface{}) {
	if !intraCache.isEnabled() {
		return
	}
	field := reflect.Indirect(reflect.ValueOf(obj)).Field(0)
	req := reflect.NewAt(
		field.Type(),
		unsafe.Pointer(field.UnsafeAddr()),
	).Elem().Interface().(RequestData)
	if req.bypassCacheWrite {
		return
	}
	return intraCache.put(obj)
}

func SetCacheEnabled(enabled bool) {
	intraCache.Lock()
	intraCache.enabled = enabled
	intraCache.Unlock()
	if !enabled {
		FlushCache()
	}
}

func SetCacheTimeout(minutes int) {
	intraCache.Lock()
	intraCache.timeout = time.Duration(minutes) * time.Minute
	intraCache.Unlock()
}

func FlushCache() {
	intraCache.Lock()
	intraCache.types = make(map[string]map[int]cachedObject)
	intraCache.Unlock()
}
