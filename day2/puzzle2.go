package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Puzzle2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var points int
	for reader.Scan() {
		slice := strings.Split(reader.Text(), " ")
		switch slice[1] {
		case "X":
			switch slice[0] {
			case "A":
				points += lose + scissors
			case "B":
				points += lose + rock
			case "C":
				points += lose + paper
			}
		case "Y":
			switch slice[0] {
			case "A":
				points += draw + rock
			case "B":
				points += draw + paper
			case "C":
				points += draw + scissors
			}
		case "Z":
			switch slice[0] {
			case "A":
				points += win + paper
			case "B":
				points += win + scissors
			case "C":
				points += win + rock
			}
		}
	}
	fmt.Println("Day2. Puzzle2. В камень ножницы бумага ты выиграешь по стратегии:", points, "\n")
}
