package day14

import (
	"fmt"
	"math"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	return solve(lines, 10)
}

func (s Solution) Part2(lines []string) (int, error) {
	return solve(lines, 40)
}

func solve(lines []string, iterations int) (int, error) {
	template, rules, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}

	pairCount := make(map[pair]int)
	for _, p := range pairs(template) {
		pairCount[p]++
	}

	for i := 0; i < iterations; i++ {
		pairCount = step(pairCount, rules)
	}
	return maxMinDifference(template, pairCount), nil
}

func parse(lines []string) (string, map[pair]rune, error) {
	var template string
	rules := make(map[pair]rune)
	for _, line := range lines {
		if template == "" {
			template = line
			continue
		}
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return "", nil, fmt.Errorf("unexpected number of fields in line: expected 3, got %d: %s", len(line), line)
		}
		lhs := []rune(parts[0])
		if len(lhs) != 2 {
			return "", nil, fmt.Errorf("unexpected number of runes in field: %s", parts[0])
		}
		rhs := []rune(parts[2])
		if len(rhs) != 1 {
			return "", nil, fmt.Errorf("unexpected rhs for rule: %s", line)
		}
		rules[pair{lhs[0], lhs[1]}] = rhs[0]
	}
	return template, rules, nil
}

func step(pairCounts map[pair]int, rules map[pair]rune) map[pair]int {
	result := make(map[pair]int)
	for p, freq := range pairCounts {
		middle, ok := rules[p]
		if !ok {
			result[p] += freq
			continue
		}
		result[pair{p.left, middle}] += freq
		result[pair{middle, p.right}] += freq
	}
	return result
}

func maxMinDifference(template string, pairCount map[pair]int) int {
	templateRunes := []rune(template)
	if len(templateRunes) < 2 {
		return 0
	}
	counts := make(map[rune]int)
	for p, freq := range pairCount {
		counts[p.left] += freq
		counts[p.right] += freq
	}
	// Except for the first and the last runes, each rune will be double counted
	// by belonging to a left pair and a right pair.
	first, last := templateRunes[0], templateRunes[len(templateRunes)-1]
	counts[first]++
	counts[last]++
	minC, maxC := int(math.MaxInt64), int(math.MinInt64)
	for _, c := range counts {
		if c < minC {
			minC = c
		}
		if c > maxC {
			maxC = c
		}
	}
	return (maxC - minC) / 2
}

func pairs(input string) []pair {
	if len(input) < 2 {
		return nil
	}
	var results []pair
	for i := 0; i < len(input)-1; i++ {
		runes := []rune(input[i : i+2])
		results = append(results, pair{runes[0], runes[1]})
	}
	return results
}

type pair struct {
	left, right rune
}

func (p pair) String() string {
	return string([]rune{p.left, p.right})
}
