package day01

import (
	"fmt"
	"strconv"

	"github.com/danielmmetz/adventofcode2021/solution"
)

type Solution struct {
	solution.Solution
}

func (s Solution) Part1(lines []string) (string, error) {
	depths := make([]int, 0, len(lines))
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("cannot convert %s to int: %w", line, err)
		}
		depths = append(depths, depth)
	}
	var count int
	for i, depth := range depths {
		if i == 0 {
			continue
		}
		if depth > depths[i-1] {
			count++
		}
	}
	return strconv.Itoa(count), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	depths := make([]int, 0, len(lines))
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return "", fmt.Errorf("cannot convert %s to int: %w", line, err)
		}
		depths = append(depths, depth)
	}
	var count int
	var prev int
	for i, depth := range depths {
		if i < 3 {
			prev += depth
			continue
		}
		current := prev - depths[i-3] + depth
		if current > prev {
			count++
		}
		prev = current
	}
	return strconv.Itoa(count), nil
}
