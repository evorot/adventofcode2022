package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		calories, err := strconv.Atoi(reader.Text())
		if err != nil {
			log.Println(err)
		}
		log.Println(calories)

	}
}
