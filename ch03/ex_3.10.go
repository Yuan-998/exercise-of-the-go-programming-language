package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12323143251234123345"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	for i := 1; i <= length; i++ {
		buf.WriteByte(s[i-1])
		if (length-i)%3 == 0 && i != length {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
