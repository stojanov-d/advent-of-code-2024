package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (map[int][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[int][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("error parsing target: %v", err)
		}

		numberStrings := strings.Fields(parts[1])
		numbers := []int{}
		for _, numStr := range numberStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing number: %v", err)
			}
			numbers = append(numbers, num)
		}

		data[target] = numbers
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func evaluate(numbers []int, target int, current int) bool {
	if len(numbers) == 0 {
		return current == target
	}

	return evaluate(numbers[1:], target, current+numbers[0]) ||
		evaluate(numbers[1:], target, current*numbers[0]) ||
		evaluate(numbers[1:], target, concat(current, numbers[0]))

}

func concat(a, b int) int {
	concatStr := fmt.Sprintf("%d%d", a, b)
	concat, _ := strconv.Atoi(concatStr)
	return concat
}

func calculate(data map[int][]int) int {
	res := 0

	for target, numbers := range data {
		if evaluate(numbers[1:], target, numbers[0]) {
			res += target
		}
	}

	return res
}

func main() {

	data, err := parseInput("./input.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	result := calculate(data)
	fmt.Println(result)
}
