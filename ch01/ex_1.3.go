package main

import (
	"fmt"
	"strings"
	"time"
)

func ItoB(i int) bool {
	if i > 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var s0, sep string
	start := time.Now()
	repetition := 10000
	for i := 0; i < repetition; i++ {
		s0 += "a" + sep
		sep = " "
	}
	end := time.Now()
	t0 := end.Sub(start)
	fmt.Printf("consumed time for +=: %v\n", t0)

	start = time.Now()
	var str []string
	for i := 0; i < repetition; i++ {
		str = append(str, "a")
	}
	end = time.Now()
	s1 := strings.Join(str, " ")
	t1 := end.Sub(start)
	fmt.Printf("consumed time for join: %v\n", t1)
	fmt.Printf("the results of += and join are the same: %v\n", ItoB(strings.Compare(s0, s1)))
	fmt.Printf("join is %v times faster than\n", int64(t0/time.Microsecond)*1.0/int64(t1/time.Microsecond))
}
