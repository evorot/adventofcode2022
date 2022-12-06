package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	re := regexp.MustCompile("[0-9]+")
	stacks := make([][]string, 9, 9)
	for reader.Scan() {
		stringData := reader.Text()
		if strings.HasPrefix(stringData, "[") {
			for i := 0; i < 9; i++ {
				stacks[i] = append([]string{stringData[:3]}, stacks[i]...)
				if i == 8 {
					break
				}
				stringData = stringData[4:]
			}
		}
		if strings.HasPrefix(stringData, "") {
			for index, stack := range stacks {
				for i, s := range stack {
					if s == "   " {
						stacks[index] = stack[:i]
					}
				}
			}

		}

		if strings.HasPrefix(stringData, "move") {
			regexStrings := re.FindAllString(stringData, -1)
			inputNum := make([]int, 0, len(regexStrings))
			for _, regexString := range regexStrings {
				num, err := strconv.Atoi(regexString)
				if err != nil {
					log.Println(err)
				}
				inputNum = append(inputNum, num)

			}
			stacks = Tusk(stacks, inputNum[0], inputNum[1], inputNum[2])
		}
	}
	result := make([]string, 0, len(stacks))
	for _, stack := range stacks {
		result = append(result, strings.Trim(stack[len(stack)-1], "[]"))
	}
	fmt.Println("Day5. Puzzle1. Answer:", strings.Join(result, ""))
}

func Puzzle2() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	re := regexp.MustCompile("[0-9]+")
	stacks := make([][]string, 9, 9)
	for reader.Scan() {
		stringData := reader.Text()
		if strings.HasPrefix(stringData, "[") {
			for i := 0; i < 9; i++ {
				stacks[i] = append([]string{stringData[:3]}, stacks[i]...)
				if i == 8 {
					break
				}
				stringData = stringData[4:]
			}
		}
		if strings.HasPrefix(stringData, "") {
			for index, stack := range stacks {
				for i, s := range stack {
					if s == "   " {
						stacks[index] = stack[:i]
					}
				}
			}

		}

		if strings.HasPrefix(stringData, "move") {
			regexStrings := re.FindAllString(stringData, -1)
			inputNum := make([]int, 0, len(regexStrings))
			for _, regexString := range regexStrings {
				num, err := strconv.Atoi(regexString)
				if err != nil {
					log.Println(err)
				}
				inputNum = append(inputNum, num)

			}
			stacks = TuskWithoutReverse(stacks, inputNum[0], inputNum[1], inputNum[2])
		}
	}
	result := make([]string, 0, len(stacks))
	for _, stack := range stacks {
		result = append(result, strings.Trim(stack[len(stack)-1], "[]"))
	}
	fmt.Println("Day5. Puzzle2. Answer:", strings.Join(result, ""))
}

func Tusk(stacks [][]string, quantity, from, to int) [][]string {
	moveStack := stacks[from-1][len(stacks[from-1])-quantity:]
	stacks[from-1] = stacks[from-1][:len(stacks[from-1])-quantity]
	for i, j := 0, len(moveStack)-1; i < j; i, j = i+1, j-1 {
		moveStack[i], moveStack[j] = moveStack[j], moveStack[i]
	}
	stacks[to-1] = append(stacks[to-1], moveStack...)
	return stacks
}

func TuskWithoutReverse(stacks [][]string, quantity, from, to int) [][]string {
	moveStack := stacks[from-1][len(stacks[from-1])-quantity:]
	stacks[from-1] = stacks[from-1][:len(stacks[from-1])-quantity]
	stacks[to-1] = append(stacks[to-1], moveStack...)
	return stacks
}
