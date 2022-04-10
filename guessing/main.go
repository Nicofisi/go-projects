package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getRandomNumber() int {
	min := 0
	max := 1000
	random := rand.Intn(max-min+1) + min
	return random
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("welcome | Hello human! You will be guessing numbers.")
	fmt.Print("next unknown 0 1000 | Proceed with a number between 0 and 1000: ")

	target := getRandomNumber()
	tries := 0

	for true {
		tries += 1
		var input int

		if _, err := fmt.Scanf("%d\n", &input); err != nil {
		}
		if input > target {
			fmt.Print("next lower 0 1000 | The number you guessed is too high! Try again: ")
		} else if input < target {
			fmt.Print("next higher 0 1000 | The number you guessed is too low! Try again: ")
		} else {
			triesStr := strconv.Itoa(tries)
			fmt.Println("success " + triesStr + " | You got it in " + triesStr + " tries!")
		}
	}

	dt := time.Now().AddDate(0, 0, 15)
	dt.Format("2006-01-02")
	fmt.Println()

	//json, _ := json.Marshal()
}
