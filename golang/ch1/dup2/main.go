package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesByLine := make(map[string]map[string]struct{}) //1.4 вроде правильно

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, filesByLine)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, filesByLine)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var fileList []string
			for file := range filesByLine[line] {
				fileList = append(fileList, file)
			}
			fmt.Printf("%d\t%s\t%v\n", n, line, fileList)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, filesByLine map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		counts[line]++

		if filesByLine[line] == nil {
			filesByLine[line] = make(map[string]struct{})
		}
		filesByLine[line][filename] = struct{}{}
	}
}
