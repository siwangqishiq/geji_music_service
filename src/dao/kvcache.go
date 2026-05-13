package dao

import (
	"time"
)

type KvCache struct {
	ID         int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Key        string    `gorm:"not null;uniqueIndex;column:key" json:"key"`
	Value      string    `gorm:"column:value" json:"value"`
	ExpireTime int64     `gorm:"column:expire_time;default:-1" json:"expire_time"`
	Status     int       `gorm:"column:status;default:0" json:"status"`
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`
}

type KvCacheDAO struct {
}

var KvcDao *KvCacheDAO

func init() {
	KvcDao = &KvCacheDAO{}
}

func (KvCache) TableName() string {
	return "kv_cache"
}

func (k *KvCacheDAO) FindAllKvCaches() ([]KvCache, error) {
	var caches []KvCache
	var result = DB.Find(&caches)
	return caches, result.Error
}

// Create 创建记录
func (k *KvCacheDAO) Create(cache *KvCache) error {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	result := DB.Create(cache)
	return result.Error
}

// CreateOrUpdate 创建或更新（基于 key）
func (k *KvCacheDAO) CreateOrUpdate(key string, value string, expireTime int64) error {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	cache := &KvCache{
		Key:        key,
		Value:      value,
		ExpireTime: expireTime,
		Status:     0,
		UpdateTime: time.Now(),
	}

	// UPSERT: 如果 key 存在则更新，否则插入
	result := DB.Where(KvCache{Key: key}).
		Assign(KvCache{
			Value:      value,
			ExpireTime: expireTime,
			UpdateTime: time.Now(),
		}).
		FirstOrCreate(cache)

	return result.Error
}

// UpdateValue 更新值
func (k *KvCacheDAO) UpdateValue(key string, value string) error {
	result := DB.Model(&KvCache{}).
		Where("key = ?", key).
		Updates(map[string]any{
			"value":       value,
			"update_time": time.Now(),
		})
	return result.Error
}

// Delete 删除记录
func (k *KvCacheDAO) Delete(id int64) error {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	result := DB.Delete(&KvCache{}, id)
	return result.Error
}

// DeleteByKey 根据 Key 删除
func (k *KvCacheDAO) DeleteByKey(key string) error {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	result := DB.Where("key = ?", key).Delete(&KvCache{})
	return result.Error
}
