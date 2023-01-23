package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "+12345.456"
	point := strings.LastIndex(s, ".")
	converted := s[:1] + comma(s[1:point]) + s[point:]
	fmt.Println(converted)
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
