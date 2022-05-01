package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isCleanString(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func filterFile(filePath string) {
	fileHandle, err := os.Open(filePath)
	checkError(err)

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Split(bufio.ScanLines)
	//wordSet := make(map[string]bool)
	for fileScanner.Scan() {
		scanText(fileScanner.Text())
	}

}

func scanText(line string) {
	currentText := strings.TrimSpace(line)
	//words := strings.Fields(currentText)
	if currentText != "" {
		isGood := isCleanString(currentText)
		if isGood {
			fmt.Printf("%v\n", strings.ToLower(currentText))
		}
	}

}

func scanStdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		scanText(scanner.Text())
	}

}

func filter(args []string) {
	no_input := true

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		no_input = false
		scanStdin()
	} else if len(args) == 1 {
		filterFile(args[0])

	}

	if len(args) != 1 && no_input {
		fmt.Println("Give me a file path, or input on stdin")
	}
}
