package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MaxValue = 100

func getRandomNumber() int {
	return rand.Intn(MaxValue + 1)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("welcome | Hello human! You will be guessing numbers.")
	fmt.Print("next unknown 0 1000 | Proceed with a number between 0 and 1000: ")

	target := getRandomNumber()
	tries := 0

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		tries += 1

		scanner.Scan()
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error | reading standard input:", err)
		}

		if input == "koniec" {
			os.Exit(0)
		}

		// string to int
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error | your input is not 'koniec' nor a valid integer: ", err)
		}

		if inputInt > target {
			fmt.Print("next lower 0 1000 | The number you guessed is too high! Try again: ")
		} else if inputInt < target {
			fmt.Print("next higher 0 1000 | The number you guessed is too low! Try again: ")
		} else {
			triesStr := strconv.Itoa(tries)
			fmt.Println("success " + triesStr + " | You got it in " + triesStr + " tries!")
			break
		}
	}

	dt := time.Now().AddDate(0, 0, 15)
	dt.Format("2006-01-02")
	fmt.Println()

	//json, _ := json.Marshal()
}
