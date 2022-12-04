package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Puzzle1() {
	alphabetMap := make(map[int32]int)
	for i, letter := range alphabet {
		alphabetMap[letter] = i + 1
	}
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	sum := 0
	for reader.Scan() {
		similarMapFirstHalf := make(map[int32]bool)
		similarMapSecondHalf := make(map[int32]bool)
		for i, letter := range reader.Text() {
			if i < len(reader.Text())/2 {
				similarMapFirstHalf[letter] = true
			} else if similarMapFirstHalf[letter] && !similarMapSecondHalf[letter] {
				similarMapSecondHalf[letter] = true
				sum += alphabetMap[letter]
			}
		}

	}
	fmt.Println("Day3. Puzzle1. Answer: ", sum)
}

func Puzzle2() {
	alphabetMap := make(map[int32]int)
	for i, letter := range alphabet {
		alphabetMap[letter] = i + 1
	}
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	sum := 0
	count := 0
	var similarMapFirstElf, similarMapSecondElf, similarMapThirdElf map[int32]bool

	for reader.Scan() {
		if count == 0 {
			similarMapFirstElf = make(map[int32]bool)
			similarMapSecondElf = make(map[int32]bool)
			similarMapThirdElf = make(map[int32]bool)
			for _, letter := range reader.Text() {
				similarMapFirstElf[letter] = true
			}
			count++
			continue
		}

		if count == 1 {
			for _, letter := range reader.Text() {
				similarMapSecondElf[letter] = true
			}
			count++
			continue
		}

		if count == 2 {
			for _, letter := range reader.Text() {
				if similarMapFirstElf[letter] && similarMapSecondElf[letter] && !similarMapThirdElf[letter] {
					similarMapThirdElf[letter] = true
					sum += alphabetMap[letter]
				}

			}
			count = 0
		}
	}
	fmt.Println("Day3. Puzzle2. Answer: ", sum, "\n")
}
