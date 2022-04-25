// wykonano etapy od 1 do 6

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const MaxValue = 10
const ScoresFileName = "scores.json"

type Score struct {
	Name  string
	Value int
	Data  time.Time
}

//type BestScoreEntry struct {
//	Name  string
//	Value int
//}

func getRandomNumber() int {
	return rand.Intn(MaxValue + 1)
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error | reading standard input: ", err)
	}

	return input
}

func endGame(scores []Score) {
	// myślałem z jakiegoś powodu, że tylko jeden najlepszy wynik każdej osoby ma się wyświetlać, ups

	//nameToBestScore := map[string]int{}
	//for _, score := range scores {
	//	if nameToBestScore[score.Name] < score.Value {
	//		nameToBestScore[score.Name] = score.Value
	//	}
	//}
	//var bestScoreEntries []BestScoreEntry
	//for Name, bestScore := range nameToBestScore {
	//	bestScoreEntries = append(bestScoreEntries, BestScoreEntry{Name, bestScore})
	//}
	//sort.Slice(bestScoreEntries, func(i, j int) bool {
	//	return bestScoreEntries[i].Value < bestScoreEntries[j].Value
	//})
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Value < scores[j].Value
	})

	fmt.Println("standings | BEST SCORES")
	for _, score := range scores {
		fmt.Println("standings | " + score.Name + " - " + strconv.Itoa(score.Value))
	}
	os.Exit(0)
}

func main() {
	var scores []Score
	data, err := ioutil.ReadFile(ScoresFileName)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr,
			"welcome | Failed to read the file with scores, I assume it's your first game! Welcome! "+
				"The error was: ", err)
	} else {
		err := json.Unmarshal(data, &scores)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr,
				"error | Failed to parse JSON from the file with scores. Try deleting it to recreate it.", err)
			return
		}
	}

	rand.Seed(time.Now().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("welcome | Hello human! You will be guessing numbers.")
		fmt.Print("Name | Type your Name: ")
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
				_, _ = fmt.Fprintln(os.Stderr, "error | Your input is not 'koniec' nor a valid integer: ", err)
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

		newScore := Score{name, tries, time.Now()}
		scores = append(scores, newScore)

		fmt.Println(scores)
		jsonScores, err := json.Marshal(scores)
		fmt.Println(string(jsonScores))
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error | Error has somehow occurred while marshalling JSON!?: ", err)
			return
		}

		// You can also write it to a file as a whole
		file, err := os.Create(ScoresFileName)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error | Error has occurred while creating your scores file: ", err)
			return
		}
		_, err = file.WriteString(string(jsonScores))
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "error | Error has occurred while saving your score to a file: ", err)
			return
		}
		_ = file.Close()

		fmt.Print("again | Play again? [Y/n] ")

		input := readLine(scanner)

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
