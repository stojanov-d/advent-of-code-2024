package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func extractByRegex(input string) ([]string, error) {
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := reg.FindAllStringSubmatch(input, -1)

	var results []string
	for _, match := range matches {
		if len(match) == 3 {
			results = append(results, fmt.Sprintf("%s %s", match[1], match[2]))
		}
	}
	return results, nil
}

func writeToFile(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func calculateSum(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		if len(numbers) == 2 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			totalSum += num1 * num2
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return totalSum, nil
}

func main() {
	input, err := parseInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	validInstructions, err := extractByRegex(input)
	if err != nil {
		fmt.Println("Error while extracting:", err)
		return
	}

	err = writeToFile("numbers.txt", validInstructions)
	if err != nil {
		fmt.Println("Error generating numbers file:", err)
		return
	}

	totalSum, err := calculateSum("numbers.txt")
	if err != nil {
		fmt.Println("Error calculating sum:", err)
		return
	}
	fmt.Printf("Sum = %d\n", totalSum)
}
