package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, iteracion int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Println("iteracion ....", iteracion)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines pendientes")
				counter -= 1
				wg.Done()
			}()
			duracion := time.Duration(rand.Intn(500) * int(time.Millisecond))
			fmt.Println("esperando ... ", duracion)
			time.Sleep(duracion)
		}()

	}

	hitCounter := Hits{}
	for i := 0; i < 20000; i++ {
		iteracion := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteracion)
	}
	fmt.Println("esperando las goroutines ")
	wg.Wait()

	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()
	fmt.Println("golpes totales: ", totalHits)
	fmt.Println("terminando...")

}
