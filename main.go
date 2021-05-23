package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Root struct {
	sync.Mutex
	// index is the key.
	NK []*Edge
}

type Edge struct {
	NK   map[int16]*Edge
	Data interface{}
	Sync sync.RWMutex
}

func Insert(tags []string, value interface{}) {
	currentEdge := ROOT
	for i, v := range tags {
		if len(tags) == i+1 {
			currentEdge = currentEdge.Add(GetKey(v), value)
		} else {
			currentEdge = currentEdge.Add(GetKey(v), nil)
		}
	}
}

func (e *Edge) Add(key int16, data interface{}) *Edge {
	e.Sync.Lock()
	defer e.Sync.Unlock()
	if e.NK == nil {
		e.NK = make(map[int16]*Edge)
		if data != nil {
			e.NK[key] = &Edge{Data: data}
		} else {
			e.NK[key] = &Edge{NK: make(map[int16]*Edge)}
		}
	} else if e.NK[key] == nil {
		if e.NK[key] == nil {
			if data != nil {
				e.NK[key] = &Edge{Data: data}
			} else {
				e.NK[key] = &Edge{NK: make(map[int16]*Edge)}
			}
		}
	}
	return e.NK[key]
}

func (e *Edge) Get(key int16) *Edge {
	e.Sync.RLock()
	defer e.Sync.RUnlock()
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
		newEdge := currentEdge.Get(int16(v))
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

var ROOT = &Edge{NK: make(map[int16]*Edge)}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func main() {

	dat, err := ioutil.ReadFile("gg/gg/gg/xx/final")
	if err != nil {
		log.Println(err)
	}
	log.Println(dat)
	dat, err = ioutil.ReadFile("gg/gg/gg/xx/final")
	if err != nil {
		log.Println(err)
	}
	log.Println(dat)
	os.Exit(1)
	// k1 := GetKey("Sveinn")
	// k2 := GetKey("sveinn")
	// k3 := GetKey("svenin")
	// log.Println(k1, k2, k3)

	// os.Exit(0)
	// x := make([]*Edge, math.MaxUint16)
	// log.Println(len(x))
	// for {
	// 	time.Sleep(10 * time.Second)
	// }
	// record1 := "stuff1"
	// record2 := "stuff2"
	// Insert("abcd1234", &record1)
	// Insert("abcd4321", &record2)
	for i := 0; i < 4000000; i++ {
		go Insert([]string{RandStringRunes(10), RandStringRunes(10), RandStringRunes(10), RandStringRunes(10), RandStringRunes(10)}, string(i))
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
