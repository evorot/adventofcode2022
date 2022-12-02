package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Puzzle() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var sumCalories []int
	var sum int
	for reader.Scan() {
		calories, err := strconv.Atoi(reader.Text())
		if err != nil {
			sumCalories = append(sumCalories, sum)
			sum = 0
		} else {
			sum += calories
		}
	}
	sort.SliceStable(sumCalories, func(i, j int) bool { return sumCalories[i] > sumCalories[j] })
	fmt.Println("Day1. Puzzle1. Самое больше кол-во калорий, которые несёт 1 эльф =", sumCalories[0])
	fmt.Println("Day1. Puzzle2. 3 лучших эльфа несут =", sumCalories[0]+sumCalories[1]+sumCalories[2], "калорий\n")
}
