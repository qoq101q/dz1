package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	numOfGo    int64 = 50
	numOfNotes int64 = 500
)

type Cache struct {
	ch []int64
}

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	var mutex sync.Mutex

	ch := Cache{
		ch: []int64{},
	}

	wg.Add(int(numOfGo))

	for i := int64(0); i < numOfGo; i++ {
		go writer(&ch, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println("Должно быть записано:", numOfNotes*numOfGo, "значений")
	fmt.Println("В итоге записано", len(ch.ch), "значений")
}

func writer(ch *Cache, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	for i := int64(0); i < numOfNotes; i++ {
		n := rand.Intn(500000)
		ch.Add(int64(n))
	}
	mutex.Unlock()
}

func (ch *Cache) Add(data int64) {
	ch.ch = append(ch.ch, data)
}
