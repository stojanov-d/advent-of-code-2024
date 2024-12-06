package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x, y int
}

var directions = []Position{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func parseInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func findGuard(grid []string) (Position, int) {
	height := len(grid)
	width := len(grid[0])
	var guardPos Position
	var directionIdx int

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch grid[i][j] {
			case '^':
				guardPos = Position{i, j}
				directionIdx = 0
			case '>':
				guardPos = Position{i, j}
				directionIdx = 1
			case 'v':
				guardPos = Position{i, j}
				directionIdx = 2
			case '<':
				guardPos = Position{i, j}
				directionIdx = 3
			}
		}
	}
	return guardPos, directionIdx
}

func simulatePatrol(grid []string, guardPos Position, directionIdx int) int {
	height := len(grid)
	width := len(grid[0])
	visited := make(map[Position]bool)
	visited[guardPos] = true

	for {
		frontPos := Position{
			x: guardPos.x + directions[directionIdx].x,
			y: guardPos.y + directions[directionIdx].y,
		}

		if frontPos.x < 0 || frontPos.y < 0 || frontPos.x >= height || frontPos.y >= width {
			break
		}

		if grid[frontPos.x][frontPos.y] == '#' {
			directionIdx = (directionIdx + 1) % 4
		} else {
			guardPos = frontPos
			visited[guardPos] = true
		}
	}

	return len(visited)
}

func main() {
	grid, err := parseInputFile("./input.txt")
	if err != nil {
		panic(err)
	}

	guardPos, directionIdx := findGuard(grid)
	result := simulatePatrol(grid, guardPos, directionIdx)
	fmt.Println(result)
}
