package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const usage = "Usage: textinfo textfile [numOfMostFrequentWordsToShow]"
const defaultNumOfWordsToShow = 50

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort. Interface to sort by Value. Used as a sortable map.
type PairList []Pair

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

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

// Read text from file at path and return the whole text as a string reference.
func readText(path *string) *string {
	dat, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal("Could not read file " + *path + "\n" + usage)
	}
	text := string(dat)
	return &text
}

// Split a string into an array of words.
func extractWords(text *string) *[]string {
	lowercase := strings.ToLower(*text)
	words := strings.FieldsFunc(lowercase, func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ' ', '"', ':', ';', '(', ')', '\n', '\r',
			'\t', '\v', '\\', '/', '\f', '\a', '\b':
			return true
		}
		return false
	})
	return &words
}

// Return a map of each word and it's frequency.
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

func getArgs() (string, int) {
	numArgs := flag.NArg()
	path := ""
	numOfWordsToShow := defaultNumOfWordsToShow
	if numArgs == 1 {
		path = flag.Arg(0)
	} else if numArgs == 2 {
		path = flag.Arg(0)
		arg2, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			log.Fatal("Second argument has to be an integer.\n" + usage)
		}
		if arg2 < 0 {
			log.Fatal("Second argument has to be a positive integer.\n" + usage)
		}
		numOfWordsToShow = arg2
	} else {
		log.Fatal("Wrong number of arguments.\n" + usage)
	}
	return path, numOfWordsToShow
}

func main() {
	flag.Parse()
	path, numOfWordsToShow := getArgs()
	text := readText(&path)
	words := extractWords(text)
	totalWordCount := len(*words)
	counts := countWords(words)
	sorted := sortMapByValue(*counts)
	fmt.Println("Total number of words:", totalWordCount)
	for i := 0; i < numOfWordsToShow && i < len(sorted); i++ {
		fmt.Printf("%5v %v\n", sorted[i].Value, sorted[i].Key)
	}
}
