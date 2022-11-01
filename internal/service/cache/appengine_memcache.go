package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/maxheckel/campfirereads/internal/domain"
	"google.golang.org/appengine/v2/memcache"
	"time"
)

type appengineMemcache struct {
}

func NewAppEngineMemcache() Cache {
	return &appengineMemcache{}
}

func (a appengineMemcache) Read(key string) (interface{}, error) {
	res, err := memcache.Get(context.Background(), key)
	if err != nil {
		fmt.Printf("ERROR RECEIVED %s", err.Error())
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

func (a appengineMemcache) Write(key string, obj interface{}, timeoutSeconds int32) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	item := &memcache.Item{
		Key:        key,
		Value:      bytes,
		Flags:      0,
		Expiration: time.Duration(timeoutSeconds) * time.Second,
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
	return memcache.Set(context.Background(), item)
}
