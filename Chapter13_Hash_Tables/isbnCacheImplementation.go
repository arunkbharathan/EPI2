package main

import "fmt"

func isbnBookCache() {
	lruCache := newISBNCache(5)
	lruCache.put("a", "bookA")
	lruCache.put("b", "bookB")
	lruCache.put("c", "bookC")
	lruCache.put("d", "bookD")
	lruCache.put("e", "bookE")
	fmt.Println(lruCache.List)
	lruCache.put("f", "bookF")
	fmt.Println(lruCache.List)
	lruCache.put("g", "bookG")
	fmt.Println(lruCache.List)
	fmt.Println(lruCache.get("d"))
	fmt.Println(lruCache.List)
}
