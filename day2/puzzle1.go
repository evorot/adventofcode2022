package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	win      = 6
	draw     = 3
	lose     = 0
	rock     = 1
	paper    = 2
	scissors = 3
)

func Puzzle1() {
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
			points += rock
			switch slice[0] {
			case "A":
				points += draw
			case "B":
				points += lose
			case "C":
				points += win
			}
		case "Y":
			points += paper
			switch slice[0] {
			case "A":
				points += win
			case "B":
				points += draw
			case "C":
				points += lose
			}
		case "Z":
			points += scissors
			switch slice[0] {
			case "A":
				points += lose
			case "B":
				points += win
			case "C":
				points += draw
			}
		}
	}
	fmt.Println("Day2. Puzzle1. В камень ножницы бумага ты выиграешь:", points)
}
