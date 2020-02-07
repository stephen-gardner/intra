package intra

import (
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
		objects map[string]cachedObject
	}
)

var (
	intraCache   = objectCache{objects: make(map[string]cachedObject)}
	cacheTimeout = 30 * time.Minute
)

func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			now := time.Now()
			intraCache.Lock()
			for key, cached := range intraCache.objects {
				if now.Sub(cached.updatedAt) >= cacheTimeout {
					delete(intraCache.objects, key)
				}
			}
			intraCache.Unlock()
		}
	}()
}

func (cache objectCache) get(endpoint string) (obj interface{}, present bool) {
	var cached cachedObject
	intraCache.RLock()
	cached, present = intraCache.objects[endpoint]
	intraCache.RUnlock()
	if !present {
		return
	}
	return cached.object, true
}

func (cache objectCache) put(endpoint string, obj interface{}) (prev interface{}) {
	cached := cachedObject{
		object:    obj,
		updatedAt: time.Now(),
	}
	intraCache.Lock()
	if p, present := intraCache.objects[endpoint]; present {
		prev = p.object
	}
	intraCache.objects[endpoint] = cached
	intraCache.Unlock()
	return
}

func GetCacheTimeout() int {
	return int(cacheTimeout / time.Minute)
}

func SetCacheTimeout(minutes int) {
	cacheTimeout = time.Duration(minutes) * time.Minute
}
