package main

import (
	"container/list"
)


const CACHE_SIZE int = 3


var m = make(map[int]int)
var lists = list.New()

func Set(key int, value int) {
	if lists.Len() == CACHE_SIZE { 
		deletes() 
	}
	m[key] = value  
	lists.PushBack(key) 

}

func Get(key int) int {
	value, yes := m[key]
	if yes {
		for a := lists.Front(); a != nil; a = a.Next() {
			if a.Value == key {
				lists.Remove(a)
				lists.PushBack(key)
				
			}
		}
		return value
	} else { 
		return -1
	}
}

func deletes() {
	var key interface{} = lists.Front().Value
	k := key.(int)      
	delete(m, k)        
	lists.Remove(lists.Front()) 

}

