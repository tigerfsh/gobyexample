package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")
	close(done)
}

func main() {
	done := make(chan struct{})
	go worker(done)
	<-done
}
