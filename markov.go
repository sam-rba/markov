// This program implements a Markov chain algorithm. It generates pseudo-random
// prose based on the given input text. It works by breaking each phrase into two
// parts: a multi-word prefix, and a single suffix word. It generates output by randomly
// choosing a suffix that follows the prefix.
//
// The program is based on Chapter 3 of "The Practice of Programming" by
// Brian Kernighan and Rob Pike.
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const (
	nPref  = 2     // number of prefix words
	maxGen = 10000 // maximum number of words to generate
)

type Prefix struct {
	words [nPref]string
}

// push advances the prefix to contain `word'.
func (p *Prefix) push(word string) {
	if nPref < 1 {
		return
	} else if nPref == 1 {
		p.words[0] = word
	} else {
		p.words = [nPref]string(append(p.words[1:], word))
	}
}

// StateMap associates each multi-word Prefix with a list of suffix words.  The starting
// state is an empty Prefix{} with the suffix list containing the first word of the
// input text. The ending state has the last few words of the input as prefix, and
// the suffix list contains an empty string.
type StateMap map[Prefix][]string

func main() {
	states := build(os.Stdin)
	generate(states, maxGen)
}

// build populates a StateMap with words read from `in'.
func build(in io.Reader) StateMap {
	states := make(StateMap)
	prefix := Prefix{}

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		suffix := scanner.Text()
		states[prefix] = append(states[prefix], suffix)
		prefix.push(suffix)
	}

	states[prefix] = append(states[prefix], "") // ending state

	return states
}

// generate produces at most `nWords' of output, one word per line.
func generate(states StateMap, nWords int) {
	prefix := Prefix{}
	for i := 0; i < maxGen; i++ {
		suffixes := states[prefix]
		suffix := suffixes[rand.Intn(len(suffixes))]
		if suffix == "" { // ending state
			return
		}
		fmt.Println(suffix)
		prefix.push(suffix)
	}
}
