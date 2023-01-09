// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchTwice(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ch) // receive from channel ch
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchTwice(url string, ch chan<- string) {
	b := make([][]byte, 2)
	for i := 0; i < 2; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}
		file := url[8:] + "_" + strconv.Itoa(i)
		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		nbytes, err := io.Copy(f, resp.Body)
		b[i], err = ioutil.ReadAll(resp.Body)
		f.Close()
		resp.Body.Close() // don't leak resources
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	}
	ch <- fmt.Sprintf("results of %v are the same: %v\n", url, bytes.Equal(b[0], b[1]))
}
