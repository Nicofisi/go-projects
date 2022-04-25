package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error reading standard input: ", err)
	}

	return input
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		line := readLine(scanner)
		noop(line)
	}
}

func noop(_ interface{}) {}
