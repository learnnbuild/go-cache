package lru

import "container/list"

type Cache struct {
	//max memory could be used, or capacity
	maxBytes int64
	// current used memory
	nBytes int64
	/* core data structure, all the entries will be added into this list, the most recently used entry should be
	the tail. it's used to keep track of the least recently used entry
	*/
	ll *list.List
	/*
		core data structure, save the key-value pairs for lookups
	*/
	cache map[string]*list.Element
	// callback function when an entry is purged
	onEvicted func(key string, value Value)
}

type Value interface {
	Len() int
}
