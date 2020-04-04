package main

import (
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
)

type isbnCache struct {
	ISBNCacheMap map[string]*string
	List         *sll.List
	LruLength    int
}

func newISBNCache(length int) *isbnCache {
	return &isbnCache{
		ISBNCacheMap: map[string]*string{},
		List:         sll.New(),
		LruLength:    length,
	}
}

func (cache *isbnCache) put(key, value string) {
	cache.List.Insert(0, key)
	cache.ISBNCacheMap[key] = &value
	if cache.List.Size() == cache.LruLength+1 {
		cache.List.Remove(cache.LruLength)
	}
	return
}

func (cache *isbnCache) get(key string) (string, bool) {
	val, ok := cache.ISBNCacheMap[key]
	if ok {
		ind := cache.List.IndexOf(key)
		cache.List.Remove(ind)
		cache.List.Insert(0, key)
		return *val, true
	}
	return "", false
}
