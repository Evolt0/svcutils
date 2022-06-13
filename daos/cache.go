package daos

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

func LoadOrSaveCache(id uint64, mu *sync.Mutex, loader func(uint64) (interface{}, error), instance *cache.Cache) (interface{}, error) {
	mu.Lock()
	defer mu.Unlock()

	var ok bool
	var err error
	var value interface{}

	idStr := fmt.Sprintf("%d", id)
	value, ok = instance.Get(idStr)
	if !ok {
		value, err = loader(id)
		if err != nil {
			return nil, err
		}
		instance.Set(idStr, value, 1*time.Hour)
	}

	return value, nil
}
