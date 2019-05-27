package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Deps map[string][]string

func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}
	filename := os.Args[1]
	search := os.Args[2]
	d, err := ParseFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(FindDependencies(d, search))
}

func FindDependencies(src Deps, file string) []string {
	els, ok := src[file]
	if !ok {
		return nil
	}
	var deps []string
	for _, el := range els {
		deps = append(deps, el)
		deps = append(deps, FindDependencies(src, el)...)
	}
	return deps
}

func ParseFile(filepath string) (Deps, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	d := make(Deps)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, ":")
		dependencies := strings.Split(strings.TrimSpace(vals[1]), " ")
		d[strings.TrimSpace(vals[0])] = dependencies
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return d, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
%s
Usage:
	%s [filepath] [search name]
Example:
	%s input.txt a.class
`, os.Args[0], os.Args[0], os.Args[0])
}
