package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>")
		return
	}
	Run(args[0])
}
