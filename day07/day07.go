package day07

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	positions, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	sort.Ints(positions)
	meetingPoint := median(positions)
	var cost int
	for _, start := range positions {
		delta := start - meetingPoint
		if delta < 0 {
			delta = -delta
		}
		cost += delta
	}
	return cost, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	positions, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	trueMean := mean(positions)
	var costToFloor int
	for _, start := range positions {
		delta := start - int(math.Floor(trueMean))
		if delta < 0 {
			delta = -delta
		}
		costToFloor += (delta + delta*delta) / 2
	}
	var costToCiel int
	for _, start := range positions {
		delta := start - int(math.Ceil(trueMean))
		if delta < 0 {
			delta = -delta
		}
		costToCiel += (delta + delta*delta) / 2
	}
	minCost := costToFloor
	if costToCiel < minCost {
		minCost = costToCiel
	}
	return minCost, nil
}

func parse(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected only one line, got %d", len(lines))
	}
	var nums []int
	for _, val := range strings.Split(lines[0], ",") {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("expected int, got %s: %w", val, err)
		}
		nums = append(nums, n)
	}
	return nums, nil
}

// median returns the median of a pre-sorted slice of ints.
func median(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums)%2 == 1 {
		return nums[len(nums)/2+1]
	}
	left := nums[len(nums)/2-1]
	right := nums[len(nums)/2]
	return (left + right) / 2
}

func mean(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	var total int
	for _, n := range nums {
		total += n
	}
	return float64(total) / float64(len(nums))
}
