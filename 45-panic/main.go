package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/a.txt")
	if err != nil {
		panic(err)
	}
}
