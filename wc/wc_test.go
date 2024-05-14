package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"testing"
)

func TestByteCount(t *testing.T) {
	filePath := "./test.txt"

	dat, _ := os.ReadFile(filePath)
	res := countBytes(dat)

	if res != 342190 {
		t.Errorf("Expected 342190 but got %v", res)
	}
}

func TestMainFunc(t *testing.T) {
	file := "./test.txt"

	os.Args = append(os.Args, file)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	// can only be called once because of flag registration
	main()

	// Test file path argument
	expected, err := regexp.Match(`.`+file+`.`, buf.Bytes())
	if !expected || err != nil {
		t.Fatalf("Argument not read correctly, %v instead of %v", &buf, file)
	}
}
