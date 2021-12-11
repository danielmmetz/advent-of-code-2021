package day01

import (
	"fmt"
	"strconv"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	depths := make([]int, 0, len(lines))
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("cannot convert %s to int: %w", line, err)
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
	return count, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	depths := make([]int, 0, len(lines))
	for _, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("cannot convert %s to int: %w", line, err)
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
	return count, nil
}
