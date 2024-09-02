// This program implements a Markov chain algorithm. It generates pseudo-random
// prose based on the input text. It works by breaking each phrase into two parts:
// a multi-word prefix, and a single word suffix. It generates output by randomly
// choosing a suffix that follows the prefix.
//
// The program is based on Chapter 3 of "The Practice of Programming" by Brian Kernighan
// and Rob Pike.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

const (
	nPrefDefault  = 2     // number of prefix words
	maxGenDefault = 10000 // maximum number of words to generate
)

func main() {
	nPref := flag.Uint("p", nPrefDefault, "number of prefix words")
	maxGen := flag.Uint("g", maxGenDefault, "maximum number of words to generate")
	flag.Parse()

	states := build(os.Stdin, *nPref)
	generate(states, *maxGen)
}

// build populates a StateMap with words read from `in'. States are keyed by
// `nPref'-word prefixes.
func build(in io.Reader, nPref uint) StateMap {
	states := StateMap{make(map[string][]string), nPref}
	prefix := make([]string, nPref)

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		suffix := scanner.Text()
		states.addSuffix(prefix, suffix)
		prefix = advance(prefix, suffix)
	}

	states.addSuffix(prefix, "") // ending state
	return states
}

// generate produces at most `maxWords' of output, one word per line.
func generate(states StateMap, maxWords uint) {
	prefix := make([]string, states.nPref)
	var i uint
	for i = 0; i < maxWords; i++ {
		suffixes, ok := states.get(prefix)
		if !ok {
			log.Fatal("state map missing prefix:", prefix)
		}
		suffix := suffixes[rand.Intn(len(suffixes))]
		if suffix == "" { // ending state
			return
		}
		fmt.Println(suffix)
		prefix = advance(prefix, suffix)
	}
}

// StateMap associates each multi-word prefix with a list of suffix words.  The
// start-state prefix is a slice of empty strings, with the suffix list containing the
// first word of the input text. The end-state has the last few words of the input as
// prefix, and the suffix list contains an empty string.
type StateMap struct {
	m     map[string][]string // prefix -> []suffix
	nPref uint                // number of prefix words
}

// addSuffix appends `suffix' to the list of suffixes associated with `prefix'.
func (sm StateMap) addSuffix(prefix []string, suffix string) {
	key := key(prefix)
	sm.m[key] = append(sm.m[key], suffix)
}

// get retrieves the list of suffixes associated with `prefix', or returns false if
// prefix is not a valid key to the state map.
func (sm StateMap) get(prefix []string) (suffixes []string, ok bool) {
	suffixes, ok = sm.m[key(prefix)]
	return
}

// key converts a multi-word prefix into a hashable string that can be used as a map key.
func key(prefix []string) string {
	return fmt.Sprintf("%q", prefix)
}

// advance shifts the prefix forward to contain `word', returning the modified prefix.
func advance(prefix []string, word string) []string {
	if len(prefix) == 1 {
		prefix[0] = word
	} else if len(prefix) > 1 {
		prefix = append(prefix[1:], word)
	}
	return prefix
}
