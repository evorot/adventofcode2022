package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	stacks := make([][]string, 9, 9)
	for reader.Scan() {
		stringData := reader.Text()
		fmt.Println(stringData)
		if strings.HasPrefix(stringData, "[") {
			for i := 0; i < 9; i++ {
				stacks[i] = append(stacks[i], stringData[:3])
				if i == 8 {
					break
				}
				stringData = stringData[4:]
			}
		}
	}
	fmt.Println(stacks)
	//fmt.Println("Day4. Puzzle1. Matched pairs:", matchedPairs)
}
