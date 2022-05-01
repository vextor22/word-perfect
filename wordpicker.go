package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func validString(key_rune rune, runes map[rune]bool, s string) (string, error) {
	var err error
	if len(s) < 4 {
		return "", errors.New("under 4 letters")
	}
	letter_encountered := make(map[rune]bool)
	key_encountered := false
	for _, r := range s {
		if key_rune == r {
			key_encountered = true
		} else if _, in_runes := runes[r]; !in_runes {
			err = errors.New("invalid letter")
		}
		letter_encountered[r] = true
	}
	if key_encountered && err == nil {
		return s, nil
	} else {
		err = errors.New("missing central letter")
		return "", err
	}

}

func generate(args []string) {

	stat, _ := os.Stdin.Stat()

	if len(args) != 1 || (stat.Mode()&os.ModeCharDevice) != 0 {
		fmt.Println("Provide a sequence of 7 letters and pipe text into stdin: abcdefg")
		return
	}
	runes := []rune(args[0])
	runemap := make(map[rune]bool)
	for i := range runes {
		runemap[runes[i]] = true
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		possible_world := scanner.Text()
		result, err := validString(runes[0], runemap, possible_world)
		if err == nil {
			fmt.Println(result)
		}
	}
}
