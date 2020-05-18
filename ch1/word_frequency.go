package main

import (
    "bufio"
    "fmt"
    "os"
)

func countWords (f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
	counts[input.Text()]++
    }
}

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
    	countWords(os.Stdin, counts)
    } else {
	for _, file := range files {
	   f, err := os.Open(file)
	   if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		continue
	   }
	   countWords(f, counts)
	   f.Close()
	}
    }
    for line, count := range counts {
	fmt.Printf("%d\t%s\n", count, line)
    }
}
