package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber() int {
	min := 0
	max := 50
	random := rand.Intn(max-min+1) + min
	return random
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rd := getRandomNumber()

	fmt.Println(rd)

	dt := time.Now().AddDate(0, 0, 15)
	fmt.Println(dt.Format("01-02-2006"))

	//json, _ := json.Marshal()
}
