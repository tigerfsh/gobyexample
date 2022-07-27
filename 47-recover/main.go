package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered. error: %v\n", r)
		}
	}()

	mayPanic()

	fmt.Println("after mayPanic")
}
