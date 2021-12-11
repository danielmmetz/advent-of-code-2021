package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	drawings, boards, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	seen := make(map[int]bool)
	for _, n := range drawings {
		seen[n] = true
		for _, b := range boards {
			if wins(b, seen) {
				return score(b, seen) * n, nil
			}
		}
	}
	return 0, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	drawings, boards, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	seen := make(map[int]bool)
	var losingBoard board
	for _, n := range drawings {
		seen[n] = true
		var remaining []board
		for _, b := range boards {
			if !wins(b, seen) {
				remaining = append(remaining, b)
			}
		}
		if len(remaining) == 1 {
			losingBoard = remaining[0]
			break
		}
		boards = remaining
	}
	for _, n := range drawings {
		seen[n] = true
		if wins(losingBoard, seen) {
			return score(losingBoard, seen) * n, nil
		}
	}
	return 0, fmt.Errorf("exhausted input without finding sole losing board")
}

func parse(lines []string) ([]int, []board, error) {
	parseBoard := func(lines []string) (board, error) {
		var b [5][5]int
		if len(lines) != 5 {
			return b, fmt.Errorf("expected 5 lines, got %d", len(lines))
		}
		for i, line := range lines {
			nums := strings.Fields(line)
			if len(nums) != 5 {
				return b, fmt.Errorf("line %d: expected 5 numbers", i)
			}
			for j, num := range nums {
				cast, err := strconv.Atoi(num)
				if err != nil {
					return b, fmt.Errorf("cannot convert line to []int: %s: %w", line, err)
				}
				b[i][j] = cast
			}
		}
		return b, nil
	}
	var drawings []int
	var boards []board
	for i := 0; i < len(lines); {
		if i == 0 {
			fields := strings.Split(lines[i], ",")
			drawings = make([]int, 0, len(fields))
			for _, s := range fields {
				n, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, fmt.Errorf("unable to convert %s to int: %w", s, err)
				}
				drawings = append(drawings, n)
			}
			i++
			continue
		}
		if len(strings.TrimSpace(lines[i])) == 0 {
			i++
			continue
		}
		if i+5 > len(lines) {
			return nil, nil, fmt.Errorf("ran out of input to consume")
		}
		board, err := parseBoard(lines[i : i+5])
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing board starting at line %d: %w", i, err)
		}
		boards = append(boards, board)
		i += 5
	}
	return drawings, boards, nil
}

type board [5][5]int

func wins(b board, seen map[int]bool) bool {
	for _, row := range b {
		rowSuccess := true
		for _, value := range row {
			if !seen[value] {
				rowSuccess = false
				break
			}
		}
		if rowSuccess {
			return true
		}
	}
	for col := 0; col < 5; col++ {
		colSuccess := true
		for row := 0; row < 5; row++ {
			if !seen[b[row][col]] {
				colSuccess = false
				break
			}
		}
		if colSuccess {
			return true
		}
	}
	return false
}

func score(b board, seen map[int]bool) int {
	var sum int
	for _, row := range b {
		for _, value := range row {
			if !seen[value] {
				sum += value
			}
		}
	}
	return sum
}
