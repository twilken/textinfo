package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "sort"
    "strings"
)

// Command line arguments.
var flagPath = flag.String("path", "./text.txt", "The textfile you want to analyse.")
var flagNumOfWordsToPrint = flag.Int("n", 50, "Number of most frequent words to show")

// A data structure to hold a key/value pair.
type Pair struct {
    Key   string
    Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value. Used as a sortable map.
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
    dat, err := ioutil.ReadFile(*flagPath)
    if err != nil {
        log.Fatal("Could not read file at " + *path)
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

func main() {
    flag.Parse()
    text := readText(flagPath)
    words := extractWords(text)
    totalWordCount := len(*words)
    counts := countWords(words)
    sorted := sortMapByValue(*counts)
    fmt.Println("Total number of words:", totalWordCount)
    for i := 0; i < *flagNumOfWordsToPrint; i++ {
        fmt.Printf("%5v %v\n", sorted[i].Value, sorted[i].Key)
    }
}
