package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ChronosX88/vala-parser/scanner"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "path", "", "Path to the file which need to read")
	flag.Parse()
	if filePath == "" {
		panic(fmt.Errorf("file path isn't specified"))
	}
	f, err := os.Open(filePath)
	check(err)
	fileInfo, err := f.Stat()
	check(err)
	if fileInfo.IsDir() {
		panic(fmt.Errorf("file is a dir, not a file"))
	}
	s := scanner.NewScanner(f)
	for {
		tok := s.Scan()
		if tok.Kind == scanner.EOF {
			os.Exit(0)
		} else if tok.Kind == scanner.Whitespace {
			continue
		}
		fmt.Println(tok)
	}
}
