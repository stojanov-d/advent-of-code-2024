package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data [][]int

	for scanner.Scan() {
		line := scanner.Text()

		stringNumbers := strings.Fields(line)

		var numbers []int
		for _, str := range stringNumbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("error converting string to int: %w", err)
			}
			numbers = append(numbers, num)
		}

		data = append(data, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return data, nil
}

func isSafeReport(report []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff < 0 {
			increasing = false
		}

		if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func isSafeWithDampener(report []int) bool {
	if isSafeReport(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		if isSafeAfterSkipping(report, i) {
			return true
		}
	}

	return false
}

func isSafeAfterSkipping(report []int, skipIndex int) bool {
	prev := -1
	increasing := true
	decreasing := true

	for i := 0; i < len(report); i++ {
		if i == skipIndex {
			continue
		}

		if prev != -1 {
			diff := report[i] - prev

			if diff < -3 || diff > 3 || diff == 0 {
				return false
			}

			if diff < 0 {
				increasing = false
			}
			if diff > 0 {
				decreasing = false
			}
		}

		prev = report[i]
	}

	return increasing || decreasing
}

func main() {
	data, err := parseFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	safeCount := 0
	for _, report := range data {
		if isSafeWithDampener(report) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}
