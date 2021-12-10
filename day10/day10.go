package day10

import (
	"fmt"
	"sort"

	"github.com/danielmmetz/adventofcode2021/solution"
)

type Solution struct {
	solution.Solution
}

func (s Solution) Part1(lines []string) (string, error) {
	var total int
	for _, line := range lines {
		var corrupted bool
		var stack []rune
		for _, r := range line {
			if corrupted {
				break
			}
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, r)
			case ')', ']', '}', '>':
				if !closes(stack[len(stack)-1], r) {
					total += corruptionPoints(r)
					corrupted = true
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return fmt.Sprint(total), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	var scores []int
	for _, line := range lines {
		var corrupted bool
		var stack []rune
		for _, r := range line {
			if corrupted {
				break
			}
			switch r {
			case '(', '[', '{', '<':
				stack = append(stack, r)
			case ')', ']', '}', '>':
				if !closes(stack[len(stack)-1], r) {
					corrupted = true
				}
				stack = stack[:len(stack)-1]
			}
		}
		if corrupted || len(stack) == 0 {
			continue
		}
		var score int
		for len(stack) > 0 {
			score = 5*score + completionPoints(complement(stack[len(stack)-1]))
			stack = stack[:len(stack)-1]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	middleIndex := len(scores) / 2
	return fmt.Sprint(scores[middleIndex]), nil
}

func closes(left, right rune) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	if left == '<' && right == '>' {
		return true
	}
	return false
}

func corruptionPoints(r rune) int {
	return map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}[r]
}

func completionPoints(r rune) int {
	return map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}[r]
}

func complement(r rune) rune {
	return map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}[r]
}
