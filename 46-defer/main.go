package main

import (
	"fmt"
	"os"
)

func main() {
	testDefer()

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func testDefer() (x int) {
	// https://stackoverflow.com/questions/52718143/is-golang-defer-statement-execute-before-or-after-return-statement
	// Whenever a program executes, it execute line by line.
	defer func(n int) {
		fmt.Printf("n: %d, x: %d\n", n, x)
	}(x)

	x = 7

	return 9
}
