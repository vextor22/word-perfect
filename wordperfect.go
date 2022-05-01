package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pick a function!")
		return
	}

	command := os.Args[1]
	switch command {
	case "filter":
		filter(args[1:])
	case "generate":
		generate(args[1:])
	}
}
