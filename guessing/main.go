package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const MaxValue = 10

type Score struct {
	name  string
	value int
	date  time.Time
}

//type BestScoreEntry struct {
//	name  string
//	value int
//}

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

func endGame(scores []Score) {
	// myślałem z jakiegoś powodu, że tylko jeden najlepszy wynik każdej osoby ma się wyświetlać, ups

	//nameToBestScore := map[string]int{}
	//for _, score := range scores {
	//	if nameToBestScore[score.name] < score.value {
	//		nameToBestScore[score.name] = score.value
	//	}
	//}
	//var bestScoreEntries []BestScoreEntry
	//for name, bestScore := range nameToBestScore {
	//	bestScoreEntries = append(bestScoreEntries, BestScoreEntry{name, bestScore})
	//}
	//sort.Slice(bestScoreEntries, func(i, j int) bool {
	//	return bestScoreEntries[i].value < bestScoreEntries[j].value
	//})
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].value < scores[j].value
	})

	fmt.Println("standings | BEST SCORES")
	for _, score := range scores {
		fmt.Println("standings | " + score.name + " - " + strconv.Itoa(score.value))
	}
	os.Exit(0)
}

func main() {
	var scores []Score
	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("welcome | Hello human! You will be guessing numbers.")
		fmt.Print("name | Type your name: ")
		name := readLine(scanner)
		maxValueStr := strconv.Itoa(MaxValue)
		fmt.Print("next unknown " + maxValueStr + " | Proceed with a number between 0 and " + maxValueStr + ": ")

		target := getRandomNumber()
		tries := 0

		for true {
			tries += 1

			input := readLine(scanner)

			if input == "koniec" {
				endGame(scores)
			}

			// string to int
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "error | your input is not 'koniec' nor a valid integer: ", err)
			}

			if inputInt > target {
				fmt.Print("next lower 0 " + maxValueStr + " | The correct number is LOWER! Try again: ")
			} else if inputInt < target {
				fmt.Print("next higher 0 " + maxValueStr + " | The correct number is HIGHER! Try again: ")
			} else {
				triesStr := strconv.Itoa(tries)
				fmt.Println("success " + triesStr + " | You got it in " + triesStr + " tries!")
				break
			}
		}

		fmt.Print("again | Play again? [Y/n] ")

		input := readLine(scanner)
		newScore := Score{name, tries, time.Now()}
		scores = append(scores, newScore)

		if input == "" || strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
			continue
		}

		endGame(scores)
	}

	dt := time.Now().AddDate(0, 0, 15)
	dt.Format("2006-01-02")
	fmt.Println()

	//json, _ := json.Marshal()
}
