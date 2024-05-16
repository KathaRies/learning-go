package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"testing"
)

var filePath = "./test.txt"

func getContent(t *testing.T) []byte {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	return dat
}

func TestByteCount(t *testing.T) {
	dat := getContent(t)
	res := countBytes(dat)

	if res != 342190 {
		t.Errorf("Expected 342190 but got %v", res)
	}
}

func TestLineCount(t *testing.T) {
	dat := getContent(t)
	res := countLines(dat)

	if res != 7145 {
		t.Errorf("Expected 7145  but got %v", res)
	}
}

func TestWordCount(t *testing.T) {
	dat := getContent(t)
	res := countWords(dat)

	if res != 58164 {
		t.Errorf("Expected 58164 but got %v", res)
	}
}

func TestCharCount(t *testing.T) {
	dat := getContent(t)
	res := countChars(dat)

	if res != 327898 {
		t.Errorf("Expected 327898 but got %v", res)
	}
}

func TestMainFunc(t *testing.T) {
	os.Args = append(os.Args, filePath)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	// can only be called once because of flag registration
	main()

	// Test file path argument
	expected, err := regexp.Match(`.`+filePath+`.`, buf.Bytes())
	if !expected || err != nil {
		t.Fatalf("Argument not read correctly, %v instead of %v", &buf, filePath)
	}
}
