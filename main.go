package main

import (
	"log"
	"math"
	"sync"
	"time"

	"github.com/google/uuid"
)

// func (e *Edge) Lock() bool {
// 	if atomic.AddUint32(e.L, 1) == 1 {
// 		return true
// 	} else {
// 		return false
// 	}

// }
// func (e *Edge) Unlock() {
// 	atomic.AddUint32(e.L, 0)
// }

type Root struct {
	sync.Mutex
	// index is the key.
	NK []*Edge
}

type Edge struct {
	// Data *SomeData
	NK []*Edge
	// Drop *[]Edge
	Data interface{}
	Sync sync.Mutex
	L    *uint32
}

func Insert(tag string, value interface{}) {
	currentEdge := ROOT
	for i, v := range []byte(tag) {
		if len(tag) == i+1 {
			currentEdge = currentEdge.Add(int8(v), value)
		} else {
			currentEdge = currentEdge.Add(int8(v), nil)
		}
	}
}

func (e *Edge) Add(key int8, data interface{}) *Edge {
	if e.NK[key] == nil {
		e.Sync.Lock()
		// we do a second nil check because two runtimes can get to
		// the lock at the same time despite the first nil check.
		if e.NK[key] == nil {
			e.NK[key] = &Edge{NK: make([]*Edge, math.MaxInt8), Data: data}
		}
		e.Sync.Unlock()
	}
	return e.NK[key]
}
func (e *Edge) Get(key int8) *Edge {
	if e.NK[key] != nil {
		return e.NK[key]
	}
	return nil
}
func Find(tag string) (*Edge, bool) {
	currentEdge := ROOT
	totalMatches := len(tag)
	currentMatches := 0
	for _, v := range []byte(tag) {
		newEdge := currentEdge.Get(int8(v))
		if newEdge != nil {
			currentMatches++
			log.Println(currentMatches, newEdge.Data, v)
			currentEdge = newEdge
		} else {
			break
		}
	}
	if totalMatches == currentMatches {
		return currentEdge, true
	}
	return currentEdge, false
}

var ROOT = &Edge{NK: make([]*Edge, math.MaxInt8)}

func main() {

	record1 := "stuff1"
	record2 := "stuff2"
	Insert("abcd1234", &record1)
	Insert("abcd4321", &record2)
	for i := 0; i < 500000; i++ {
		Insert(uuid.New().String(), string(i))
	}
	// Insert("abcdddsa")
	// Insert("abcdcxxxxxx")
	// level := 0
	// IterateRoot(ROOT, level, 0)
	e, found := Find("abcd1234")
	log.Println(e.Data, found)
	e, found = Find("abcd")
	log.Println(e.Data, found)
	for {
		time.Sleep(10 * time.Second)
	}
}
func IterateRoot(root *Edge, level int, parent int8) {
	currentLevel := level
	for i, v := range root.NK {
		if v != nil {
			log.Println(currentLevel, string(parent), " >> ", string(i))
			currentLevel++
			IterateRoot(v, currentLevel, int8(i))
			currentLevel = level
		}
	}
}
