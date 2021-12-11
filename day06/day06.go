package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	ages, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	population := make(map[int]int)
	for _, age := range ages {
		population[age]++
	}
	total := grow(population, 80)
	return total, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	ages, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	population := make(map[int]int)
	for _, age := range ages {
		population[age]++
	}
	total := grow(population, 256)
	return total, nil
}

func parse(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line, got %d", len(lines))
	}
	var result []int
	for _, val := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("expected int, got %s: %w", val, err)
		}
		result = append(result, n)
	}
	return result, nil
}

func grow(population map[int]int, days int) int {
	for day := 0; day < days; day++ {
		nextPopulation := make(map[int]int)
		nextPopulation[8] = population[0]
		nextPopulation[6] = population[0]
		for i := 1; i <= 8; i++ {
			nextPopulation[i-1] += population[i]
		}
		population = nextPopulation
	}
	var total int
	for i := 0; i <= 8; i++ {
		total += population[i]
	}
	return total
}
