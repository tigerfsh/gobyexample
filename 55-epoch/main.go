package main

import (
	"fmt"
	"time"
)

var p = fmt.Println

func main() {
	now := time.Now()
	p(now)

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())

	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))
}
