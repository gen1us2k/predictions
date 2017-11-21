package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	length = 32768
)

var (
	sorted   [length]int
	unsorted [length]int
)

func init() {
	for i := 0; i < length; i++ {
		sorted[i] = i
	}
	for i := 0; i < length; i++ {
		unsorted[i] = rand.Int()
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	start := time.Now()
	var sum int

	end := time.Now()

	end = time.Now()
	sum = 0
	// start = time.Now()
	// for j := 0; j < 100000; j++ {
	// 	for i := 0; i < length; i++ {
	// 		if sorted[i] >= 128 {
	// 			sum += sorted[i]
	// 		}
	// 	}
	// }
	// end = time.Now()
	// fmt.Printf("Sorted %f seconds\n", end.Sub(start).Seconds())
	// sum = 0
	// start = time.Now()
	// for j := 0; j < 100000; j++ {
	// 	for i := 0; i < length; i++ {
	// 		if unsorted[i] >= 1000 {
	// 			sum += unsorted[i]
	// 		}
	// 	}
	// }
	// end = time.Now()
	// fmt.Printf("Unsorted %f  seconds\n", end.Sub(start).Seconds())
	// sum = 0
	start = time.Now()
	for j := 0; j < 100000; j++ {
		for i := 0; i < length; i++ {
			t := (unsorted[i] - 1000) >> 31
			sum += ^t & unsorted[i]
		}
	}
	end = time.Now()
	fmt.Printf("Unsorted branchless %f  seconds\n", end.Sub(start).Seconds())
}
