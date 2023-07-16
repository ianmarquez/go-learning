package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func runWorkers(wg *sync.WaitGroup) {
	const workerPool = 10
	for i := 1; i <= workerPool; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
}

func main() {
	var wg sync.WaitGroup
	runWorkers(&wg)
	wg.Wait()
}
