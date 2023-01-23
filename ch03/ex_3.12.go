package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "cab"
	s2 := "bac"
	fmt.Printf("%v == %v: %v\n", s1, s2, outOfOrderCompare(s1, s2))
}

func reorder(s string) string {
	tmp := strings.Split(s, "")
	sort.Strings(tmp)
	var buffer bytes.Buffer
	for _, s := range tmp {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func outOfOrderCompare(s1, s2 string) bool {
	return strings.EqualFold(reorder(s1), reorder(s2))
}
