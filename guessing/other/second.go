package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "main.go")
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(stdout))
}
