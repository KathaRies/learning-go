package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)

	byteFlag := flag.Bool("c", false, "Count bytes of a file")

	flag.Parse()

	filePath := flag.Arg(0)

	if filePath == "" {
		log.Fatal("No file given")
	}

	//log.Printf("Analysing file '%v'\n", filePath)

	dat, err := os.ReadFile(filePath)
	checkForErrors(err)
	if *byteFlag {
		log.Printf("%v %v", countBytes(dat), filepath.Base(filePath))
	}
}

func countBytes(text []byte) int {
	return len(text)
}

func checkForErrors(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
