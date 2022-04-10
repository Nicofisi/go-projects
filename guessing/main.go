package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const MaxValue = 10

func getRandomNumber() int {
	return rand.Intn(MaxValue + 1)
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error | reading standard input:", err)
	}

	return input
}

func main() {
	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("welcome | Hello human! You will be guessing numbers.")
		maxValueStr := strconv.Itoa(MaxValue)
		fmt.Print("next unknown " + maxValueStr + " | Proceed with a number between 0 and " + maxValueStr + ": ")

		target := getRandomNumber()
		tries := 0

		for true {
			tries += 1

			input := readLine(scanner)

			if input == "koniec" {
				os.Exit(0)
			}

			// string to int
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "error | your input is not 'koniec' nor a valid integer: ", err)
			}

			if inputInt > target {
				fmt.Print("next lower 0 " + maxValueStr + " | The number you guessed is too high! Try again: ")
			} else if inputInt < target {
				fmt.Print("next higher 0 " + maxValueStr + " | The number you guessed is too low! Try again: ")
			} else {
				triesStr := strconv.Itoa(tries)
				fmt.Println("success " + triesStr + " | You got it in " + triesStr + " tries!")
				break
			}
		}

		fmt.Print("again | Play again? [Y/n] ")

		input := readLine(scanner)

		if input == "" || strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
			continue
		}

		os.Exit(0)
	}

	dt := time.Now().AddDate(0, 0, 15)
	dt.Format("2006-01-02")
	fmt.Println()

	//json, _ := json.Marshal()
}
