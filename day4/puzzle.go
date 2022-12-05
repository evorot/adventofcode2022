package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var matchedPairs int
	for reader.Scan() {
		pairs := strings.Split(reader.Text(), ",")
		separatedPairs := make([][]string, 0, len(pairs))
		for _, pair := range pairs {
			separatedPairs = append(separatedPairs, strings.Split(pair, "-"))
		}

		firstElementFirstPair, _ := strconv.Atoi(separatedPairs[0][0])
		secondElementFirstPair, _ := strconv.Atoi(separatedPairs[0][1])

		firstElementSecondPair, _ := strconv.Atoi(separatedPairs[1][0])
		secondElementSecondPair, _ := strconv.Atoi(separatedPairs[1][1])
		var count int
		for i := firstElementFirstPair; i <= secondElementFirstPair; i++ {

			for j := firstElementSecondPair; j <= secondElementSecondPair; j++ {
				if i == j {
					count++
				}
			}

		}
		if count == secondElementSecondPair-firstElementSecondPair+1 {
			matchedPairs++
			continue
		}

		count = 0
		for i := firstElementSecondPair; i <= secondElementSecondPair; i++ {
			for j := firstElementFirstPair; j <= secondElementFirstPair; j++ {
				if i == j {
					count++
				}
			}
		}
		if count == secondElementFirstPair-firstElementFirstPair+1 {
			matchedPairs++
		}

	}
	fmt.Println("Day4. Puzzle1. Matched pairs:", matchedPairs)
}

func Puzzle2() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var matchedPairs int
	for reader.Scan() {
		pairs := strings.Split(reader.Text(), ",")
		separatedPairs := make([][]string, 0, len(pairs))
		for _, pair := range pairs {
			separatedPairs = append(separatedPairs, strings.Split(pair, "-"))
		}

		firstElementFirstPair, _ := strconv.Atoi(separatedPairs[0][0])
		secondElementFirstPair, _ := strconv.Atoi(separatedPairs[0][1])

		firstElementSecondPair, _ := strconv.Atoi(separatedPairs[1][0])
		secondElementSecondPair, _ := strconv.Atoi(separatedPairs[1][1])
		var count int
		for i := firstElementFirstPair; i <= secondElementFirstPair; i++ {
			for j := firstElementSecondPair; j <= secondElementSecondPair; j++ {
				if i == j {
					count++
				}
			}
			if count > 0 {
				matchedPairs++
				break
			}
		}
	}
	fmt.Println("Day4. Puzzle2. Matched pairs:", matchedPairs, "\n")
}
