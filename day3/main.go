package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(partOne("test_input.txt"))
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input string
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = line
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	matches := re.FindAllStringSubmatch(input, -1)
	total := 0
	for _, match := range matches {
		first_num, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		second_num, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		total += first_num * second_num
	}

	return total
}
