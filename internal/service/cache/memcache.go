package cache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/maxheckel/campfirereads/internal/domain"
)

type memcacheDriver struct {
	memcache *memcache.Client
}

const (
	GetBestSellerList = iota
	AllListsBestSellers
	Book
	AmazonListings
)

func NewMemcache(address string) Cache {
	mc := memcache.New(address)
	return &memcacheDriver{memcache: mc}
}

func (m *memcacheDriver) Read(key string) (interface{}, error) {
	res, err := m.memcache.Get(key)
	if err != nil && err.Error() != "memcache: cache miss" {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	var obj interface{}
	switch res.Flags {
	case Book:
		obj = &domain.Book{}
	case AllListsBestSellers:
		obj = &domain.AllListsBestSellers{}
	case GetBestSellerList:
		obj = &domain.GetBestSellerList{}
	case AmazonListings:
		obj = &domain.AmazonListings{}
	}
	err = json.Unmarshal(res.Value, obj)
	return obj, err
}

func (m *memcacheDriver) Write(key string, obj interface{}, timeoutSeconds int32) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	item := &memcache.Item{
		Key:        key,
		Value:      bytes,
		Flags:      0,
		Expiration: timeoutSeconds,
	}
	if _, ok := obj.(*domain.Book); ok {
		item.Flags = Book
	}
	if _, ok := obj.(*domain.AllListsBestSellers); ok {
		item.Flags = AllListsBestSellers
	}
	if _, ok := obj.(*domain.GetBestSellerList); ok {
		item.Flags = GetBestSellerList
	}
	if _, ok := obj.(*domain.AmazonListings); ok {
		item.Flags = AmazonListings
	}
	return m.memcache.Set(item)
}
