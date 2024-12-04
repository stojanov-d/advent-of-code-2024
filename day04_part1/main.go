package main

import (
	"bufio"
	"fmt"
	"os"
)

func countWordOccurrences(matrix [][]rune, word string) int {
	directions := [][2]int{
		{0, 1},   // gore
		{1, 0},   // dole
		{1, 1},   // dole desno
		{1, -1},  // dole levo
		{0, -1},  // levo
		{-1, 0},  // gore
		{-1, -1}, // gore levo
		{-1, 1},  // gore desno
	}

	wordRunes := []rune(word)
	wordLen := len(wordRunes)
	rows := len(matrix)
	cols := len(matrix[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if matches(matrix, r, c, dir, wordRunes, wordLen, rows, cols) {
					count++
				}
			}
		}
	}

	return count
}

func matches(matrix [][]rune, r, c int, dir [2]int, wordRunes []rune, wordLen, rows, cols int) bool {
	for i := 0; i < wordLen; i++ {
		nr, nc := r+i*dir[0], c+i*dir[1]
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols || matrix[nr][nc] != wordRunes[i] {
			return false
		}
	}
	return true
}

func parseInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func main() {
	matrix, err := parseInput("./input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	word := "XMAS"
	fmt.Println(countWordOccurrences(matrix, word))
}
