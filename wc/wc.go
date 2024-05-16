package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	log.SetFlags(0)

	flags := map[string]*bool{
		"bytes": flag.Bool("c", false, "Count bytes of a file"),
		"lines": flag.Bool("l", false, "Count lines of a file"),
		"words": flag.Bool("w", false, "Count words of a file"),
		"chars": flag.Bool("m", false, "Count characters of a file"),
	}

	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		log.Fatal("No file given")
	}

	printHeader(flags)

	dat, err := os.ReadFile(filePath)
	checkForErrors(err)
	stats := []string{}
	for k, v := range flags {
		if *v {
			switch k {
			case "bytes":
				stats = append(stats, fmt.Sprint(countBytes(dat)))
			case "lines":
				stats = append(stats, fmt.Sprint(countLines(dat)))
			case "words":
				stats = append(stats, fmt.Sprint(countWords(dat)))
			case "chars":
				stats = append(stats, fmt.Sprint(countChars(dat)))
			}
		}
	}
	stats = append(stats, filepath.Base(filePath))

	log.Print(strings.Join(stats, " "))
}

func printHeader(flags map[string]*bool) {
	none := true
	for _, v := range flags {
		none = none && !*v
	}
	header := []string{}
	for k, v := range flags {
		if none {
			*v = true
		}
		if *v {
			header = append(header, k)
		}
	}
	log.Print(strings.Join(header, " | "))
}

func countBytes(text []byte) int {
	return len(text)
}

func countChars(text []byte) int {
	return len(string(text))
}

func countLines(text []byte) int {
	all := string(text)
	lines := strings.Split(all, "\n")
	return len(lines)
}

func countWords(text []byte) int {
	all := string(text)
	wordCount := 0
	lines := strings.Split(all, "\n")
	for _, l := range lines {
		words := strings.Split(l, " ")
		for _, w := range words {
			if isWord(w) {
				wordCount += 1
			}
		}
	}
	return wordCount
}

// Fix this as it is the nice code solution
func countWordsRegex(text []byte) int {
	matcher, _ := regexp.Compile(`\W+\w+\W+`)
	words := matcher.FindAll(text, -1)
	log.Print(string(words[0]))
	log.Print(string(words[1]))
	return len(words)
}

func isWord(word string) bool {
	return word != "" && word != "\n" && word != "\r"
}

func checkForErrors(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
