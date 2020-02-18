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
		categories map[string]map[int]cachedObject
	}
)

const (
	catCloses          = "closes"
	catCursusUsers     = "cursus_users"
	catExperiences     = "experiences"
	catLocations       = "locations"
	catProjectSessions = "project_sessions"
	catProjects        = "projects"
	catTeams           = "teams"
)

var (
	intraCache   = objectCache{categories: make(map[string]map[int]cachedObject)}
	cacheTimeout = 30 * time.Minute
)

func init() {
	go func() {
		for {
			time.Sleep(time.Minute)
			now := time.Now()
			intraCache.Lock()
			for category, objects := range intraCache.categories {
				for key, cached := range objects {
					if now.Sub(cached.updatedAt) >= cacheTimeout {
						delete(objects, key)
					}
				}
				if len(objects) == 0 {
					delete(intraCache.categories, category)
				}
			}
			intraCache.Unlock()
		}
	}()
}

func (cache objectCache) delete(category string, ID int) (prev interface{}) {
	intraCache.Lock()
	if objects, present := intraCache.categories[category]; present {
		if p, present := objects[ID]; present {
			prev = p.object
			delete(objects, ID)
		}
	}
	intraCache.Unlock()
	return
}

func (cache objectCache) get(category string, ID int) (obj interface{}, present bool) {
	var objects map[int]cachedObject
	var cached cachedObject
	intraCache.RLock()
	if objects, present = intraCache.categories[category]; present {
		cached, present = objects[ID]
	}
	intraCache.RUnlock()
	if present {
		obj = cached.object
	}
	return
}

func (cache objectCache) put(category string, ID int, obj interface{}) (prev interface{}) {
	cached := cachedObject{
		object:    obj,
		updatedAt: time.Now(),
	}
	intraCache.Lock()
	objects, present := intraCache.categories[category]
	if !present {
		objects = make(map[int]cachedObject)
		intraCache.categories[category] = objects
	}
	if p, present := objects[ID]; present {
		prev = p.object
	}
	objects[ID] = cached
	intraCache.Unlock()
	return
}

func GetCacheTimeout() int {
	return int(cacheTimeout / time.Minute)
}

func SetCacheTimeout(minutes int) {
	cacheTimeout = time.Duration(minutes) * time.Minute
}
