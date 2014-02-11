package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var flagPathToTextfile = flag.String("path", "./text.txt", "The textfile you want to analyse.")

func init() {
	flag.Parse()
}

func main() {
	dat, err := ioutil.ReadFile(*flagPathToTextfile)
	if err != nil {
		log.Fatal("Could not read file.")
	}
	text := string(dat)
	//words := strings.Fields(text)

	words := strings.FieldsFunc(text, func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ' ', '"', '\'', ':', ';', '(', ')', '\n', '\r',
			'\t', '\v', '\\', '/', '\f', '\a', '\b':
			return true
		}
		return false
	})

	counts := make(map[string]int)
	for _, word := range words {
		_, present := counts[word]
		if present {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	fmt.Println("Info about:", *flagPathToTextfile)
	for key, val := range counts {
		fmt.Printf("%5v %v\n", val, key)
	}
}
