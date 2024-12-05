package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func parseRules(lines []string) []Rule {
	var rules []Rule
	for _, line := range lines {
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{Before: before, After: after})
	}
	return rules
}

func parseUpdates(lines []string) [][]int {
	var updates [][]int
	for _, line := range lines {
		parts := strings.Split(line, ",")
		var update []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return updates
}

func check(update []int, rules []Rule) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}
	for _, rule := range rules {
		beforeIdx, beforeExists := position[rule.Before]
		afterIdx, afterExists := position[rule.After]
		if beforeExists && afterExists && beforeIdx > afterIdx {
			return false
		}
	}
	return true
}

func getMiddleElement(update []int) int {
	return update[len(update)/2]
}

func reorderedSum(updates [][]int, rules []Rule) int {
	reorderUpdate := func(update []int, rules []Rule) []int {
		less := func(a, b int) bool {
			for _, rule := range rules {
				if rule.Before == a && rule.After == b {
					return true
				}
				if rule.Before == b && rule.After == a {
					return false
				}
			}
			return a < b
		}

		sorted := append([]int(nil), update...)

		for i := 0; i < len(sorted); i++ {
			for j := 0; j < len(sorted)-i-1; j++ {
				if !less(sorted[j], sorted[j+1]) {
					sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				}
			}
		}
		return sorted
	}

	var middleElementsSum int
	for _, update := range updates {
		if !check(update, rules) {
			reordered := reorderUpdate(update, rules)
			middleElementsSum += getMiddleElement(reordered)
		}
	}

	return middleElementsSum
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var rulesLines, updatesLines []string
	scanner := bufio.NewScanner(file)
	parsingRules := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parsingRules = false
			continue
		}
		if parsingRules {
			rulesLines = append(rulesLines, line)
		} else {
			updatesLines = append(updatesLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	rules := parseRules(rulesLines)
	updates := parseUpdates(updatesLines)

	var totalMiddleSum int
	for _, update := range updates {
		if check(update, rules) {
			totalMiddleSum += getMiddleElement(update)
		}
	}

	part2Sum := reorderedSum(updates, rules)

	fmt.Println(totalMiddleSum)
	fmt.Println(part2Sum)
}
