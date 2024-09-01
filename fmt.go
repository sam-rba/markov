// This program wraps its input into short lines. It is based on fmt.awk from "The
// Practice of Programming" by Brian Kernighan and Rob Pike (pp. 229).
package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxWidth = 80 // maximum number of characters per line

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	width := 0
	for input.Scan() {
		word := input.Text()

		if width+1+len(word) > maxWidth {
			fmt.Println()
			width = 0
		}

		if width == 0 {
			fmt.Print(word)
			width += len(word)
		} else {
			fmt.Printf(" %s", word)
			width += 1 + len(word)
		}
	}
	fmt.Println()
}
