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

// if a WaitGroup is explicitly passed into functions, it should be done by pointer.

func main() {
	var wg sync.WaitGroup
	// for i := 1; i <= 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		worker(i)
	// 	}(i)
	// }
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
