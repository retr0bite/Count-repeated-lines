package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	counts := make(map[string]int)
	file := os.Args[1:]
	if len(file) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range file {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error : %v", err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}
	for _, filename := range os.Args[1:] {
		_, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		fmt.Println(filename)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
