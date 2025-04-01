package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partTwo("test_input.txt"))
	fmt.Println(partTwo("input.txt"))

}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reportList [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		var report []int
		for _, v := range nums {
			n, error := strconv.Atoi(v)
			if error != nil {
				fmt.Println("Conversion error:", error)
			}
			report = append(report, n)

		}
		reportList = append(reportList, report)
	}

	total := 0
	for _, report := range reportList {
		if reportSafe(report) {
			total++
		}
	}
	return total
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reportList [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		var report []int
		for _, v := range nums {
			n, error := strconv.Atoi(v)
			if error != nil {
				fmt.Println("Conversion error:", error)
			}
			report = append(report, n)
		}
		reportList = append(reportList, report)
	}

	total := 0
	for _, report := range reportList {
		if reportSafeTolerance(report) {
			total++
		}
	}

	return total
}

func reportSafeTolerance(report []int) bool {

	isSafe := reportSafe(report)

	if !isSafe {
		for i := 0; i < len(report); i++ {
			newReport := []int{}
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			isSafe = reportSafe(newReport)

			if isSafe {
				break
			}
		}
	}

	return isSafe
}

func reportSafe(report []int) bool {

	var increasing bool

	if report[1] > report[0] {
		increasing = true
	} else {
		increasing = false
	}

	for i := 1; i < len(report); i++ {
		if increasing {
			diff := report[i] - report[i-1]
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		} else {
			diff := report[i] - report[i-1]
			if !(diff <= -1 && diff >= -3) {
				return false
			}
		}
	}

	return true
}
