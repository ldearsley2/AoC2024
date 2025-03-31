package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partOne("input.txt"))
	fmt.Println(partTwo("input.txt"))
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var first, second []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		num_one, _ := strconv.Atoi(nums[0])
		num_two, _ := strconv.Atoi(nums[1])

		first = append(first, num_one)
		second = append(second, num_two)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(first)
	sort.Ints(second)

	total := 0
	for i, _ := range first {
		diff := abs(second[i] - first[i])
		total += diff
	}

	return total
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var first []int
	countMap := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		num_one, _ := strconv.Atoi(nums[0])
		num_two, _ := strconv.Atoi(nums[1])

		first = append(first, num_one)
		countMap[num_two]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, num := range first {
		total += num * countMap[num]
	}

	return total
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
