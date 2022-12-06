package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Puzzle1() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var result int
	distChar := 4
	for reader.Scan() {
		inputText := reader.Text()
	loop:
		for i, j := 0, distChar; j < len(inputText); i, j = i+1, j+1 {
			checkText := inputText[i:j]
			checkMap := make(map[int32]bool)
			for _, letter := range checkText {
				if !checkMap[letter] {
					checkMap[letter] = true
				} else {
					continue loop
				}

			}
			if len(checkMap) == distChar {
				result = j
				break loop
			}
		}
	}
	fmt.Println("Day6. Puzzle1. Answer:", result)
}

func Puzzle2() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var result int
	distChar := 14
	for reader.Scan() {
		inputText := reader.Text()
	loop:
		for i, j := 0, distChar; j < len(inputText); i, j = i+1, j+1 {
			checkText := inputText[i:j]
			checkMap := make(map[int32]bool)
			for _, letter := range checkText {
				if !checkMap[letter] {
					checkMap[letter] = true
				} else {
					continue loop
				}

			}
			if len(checkMap) == distChar {
				result = j
				break loop
			}
		}
	}
	fmt.Println("Day6. Puzzle2. Answer:", result)
}
