package service

import (
	"geji/dao"
	"geji/util"
	"sync"
	"time"
)

type KVCacheService struct {
	cache      map[string]dao.KvCache
	cacheMutex sync.Mutex
}

var KVSvr KVCacheService = KVCacheService{
	cache: make(map[string]dao.KvCache),
}

func init() {
	util.Logi("KVCache service init")
	KVSvr.load()
}

func (k *KVCacheService) load() {
	util.Logi("KVCache service load all keys")
	k.cacheMutex.Lock()
	defer k.cacheMutex.Unlock()

	kvCaches, err := dao.KvcDao.FindAllKvCaches()
	if err != nil {
		util.Loge("KVCache FindAllKvCaches error %v", err)
		return
	}

	for _, cache := range kvCaches {
		k.cache[cache.Key] = cache
	} //end for each

	util.Logi("KVCache load data %d", len(k.cache))
}

func (k *KVCacheService) Put(key string, value string) bool {
	util.Logi("KVCache put key:%s value:%s", key, value)

	if !k.cacheMutex.TryLock() {
		return false
	}
	defer k.cacheMutex.Unlock()

	var cache = dao.KvCache{
		Key:        key,
		Value:      value,
		ExpireTime: -1,
		Status:     0,
		UpdateTime: time.Now(),
	}
	k.cache[key] = cache
	go dao.KvcDao.CreateOrUpdate(key, value, -1)
	return true
}

func (k *KVCacheService) Get(key string) string {
	util.Logi("KVCache get key:%s", key)

	if !k.cacheMutex.TryLock() {
		return ""
	}
	defer k.cacheMutex.Unlock()

	cache, exist := k.cache[key]
	if !exist {
		return ""
	}
	return cache.Value
}

func (k *KVCacheService) Del(key string) bool {
	util.Logi("KVCache del key:%s", key)
	if !k.cacheMutex.TryLock() {
		return false
	}
	defer k.cacheMutex.Unlock()

	if _, exist := k.cache[key]; exist {
		delete(k.cache, key)
		go dao.KvcDao.DeleteByKey(key)
		return true
	}
	return false
}
