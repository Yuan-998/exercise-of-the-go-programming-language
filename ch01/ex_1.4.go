// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdIn", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for line, m := range counts {
		cnt := 0
		for _, arg := range files {
			cnt += m[arg]
		}
		if cnt > 1 {
			fmt.Printf("%d\t%s\n", cnt, line)
			fmt.Printf("files: \n")
			for _, arg := range files {
				if m[arg] >= 1 {
					fmt.Println(arg)
				}
			}
			fmt.Println()
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][arg]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
