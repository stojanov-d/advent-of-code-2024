package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("wrong input fmt: %s", line)
		}

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func calculateTotalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := math.Abs(float64(left[i] - right[i]))
		totalDistance += int(distance)
	}

	return totalDistance
}

func calculateSimilarityScore(left, right []int) int {
	rightCount := make(map[int]int)
	for _, i := range right {
		rightCount[i]++
	}

	similarityScore := 0
	for _, i := range left {
		similarityScore += i * rightCount[i]
	}

	return similarityScore
}

func main() {

	left, right, err := parseInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	totalDistance := calculateTotalDistance(left, right)
	fmt.Println("Total distance:", totalDistance)

	similarityScore := calculateSimilarityScore(left, right)
	fmt.Println("Similarity score:", similarityScore)
}
