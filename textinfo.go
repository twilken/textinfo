package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var flagPath = flag.String("path", "./text.txt", "The textfile you want to analyse.")

func init() {
	flag.Parse()
}

func readText(path *string) *string {
	dat, err := ioutil.ReadFile(*flagPath)
	if err != nil {
		log.Fatal("Could not read file at " + *path)
	}
	text := string(dat)
	return &text
}

func extractWords(text *string) *[]string {
	words := strings.FieldsFunc(*text, func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ' ', '"', '\'', ':', ';', '(', ')', '\n', '\r',
			'\t', '\v', '\\', '/', '\f', '\a', '\b':
			return true
		}
		return false
	})
	return &words
}

func countWords(words *[]string) *map[string]int {
	counts := make(map[string]int)
	for _, word := range *words {
		_, present := counts[word]
		if present {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return &counts
}

func main() {
	text := readText(flagPath)
	words := extractWords(text)
	counts := countWords(words)

	fmt.Println("Info about:", *flagPath)
	for key, val := range *counts {
		fmt.Printf("%5v %v\n", val, key)
	}
}
