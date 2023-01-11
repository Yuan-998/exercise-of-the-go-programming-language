package main

import (
	"ch02/popcount"
	"fmt"
	"time"
)

func main() {
	const num = 10000
	start := time.Now()
	for i := 0; i < num; i++ {
		popcount.PopCount(uint64(i))
	}
	end := time.Now()
	fmt.Printf("popcount: %v\n", end.Sub(start))

	start = time.Now()
	for i := 0; i < num; i++ {
		popcount.PopCountLoop(uint64(i))
	}
	end = time.Now()
	fmt.Printf("popcountloop: %v\n", end.Sub(start))
}
