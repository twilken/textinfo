package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var flagPath = flag.String("path", "./text.txt", "The textfile you want to analyse.")

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

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
		case '.', ',', '!', '?', ' ', '"', ':', ';', '(', ')', '\n', '\r',
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
	sorted := sortMapByValue(*counts)

	for _, pair := range sorted {
		fmt.Printf("%5v %v\n", pair.Value, pair.Key)
	}
}
