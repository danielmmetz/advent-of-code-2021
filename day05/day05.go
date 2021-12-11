package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	ventLines, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	hits := make(map[point]int)
	for i, line := range ventLines {
		ps, err := pointsBetween(line.start, line.end, false)
		if err != nil {
			return 0, fmt.Errorf("error calculating points on vent line %d: %w", i, err)
		}
		for _, p := range ps {
			hits[p]++
		}
	}
	var dangerSpots int
	for _, count := range hits {
		if count >= 2 {
			dangerSpots++
		}
	}
	return dangerSpots, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	ventLines, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	hits := make(map[point]int)
	for i, line := range ventLines {
		ps, err := pointsBetween(line.start, line.end, true)
		if err != nil {
			return 0, fmt.Errorf("error calculating points on vent line %d: %w", i, err)
		}
		for _, p := range ps {
			hits[p]++
		}
	}
	var dangerSpots int
	for _, count := range hits {
		if count >= 2 {
			dangerSpots++
		}
	}
	return dangerSpots, nil
}

func parse(lines []string) ([]ventLine, error) {
	var ventLines []ventLine
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, fmt.Errorf("unexpected content at line %d: %s", i, line)
		}
		parts = strings.Split(fmt.Sprintf("%s,%s", parts[0], parts[2]), ",")
		if len(parts) != 4 {
			return nil, fmt.Errorf("expected 4 comma separated numbers for line %d: %s", i, line)
		}
		nums := make([]int, 0, 4)
		for _, s := range parts {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("error converting %s to int: %w", s, err)
			}
			nums = append(nums, n)
		}
		ventLines = append(ventLines, ventLine{
			start: point{x: nums[0], y: nums[1]},
			end:   point{x: nums[2], y: nums[3]},
		})
	}
	return ventLines, nil
}

type ventLine struct {
	start, end point
}

type point struct {
	x, y int
}

func pointsBetween(p1, p2 point, includeDiagonals bool) ([]point, error) {
	var points []point
	switch {
	case p1.x == p2.x:
		if p1.y >= p2.y {
			p1, p2 = p2, p1
		}
		for y := p1.y; y <= p2.y; y++ {
			points = append(points, point{x: p1.x, y: y})
		}
	case p1.y == p2.y:
		if p1.x >= p2.x {
			p1, p2 = p2, p1
		}
		for x := p1.x; x <= p2.x; x++ {
			points = append(points, point{x: x, y: p1.y})
		}
	case includeDiagonals:
		return diagonalPointsBetween(p1, p2)
	default:
		return nil, nil
	}
	return points, nil
}

func diagonalPointsBetween(p1, p2 point) ([]point, error) {
	xDelta := p2.x - p1.x
	yDelta := p2.y - p1.y
	absXDelta, absYDelta := xDelta, yDelta
	if xDelta < 0 {
		absXDelta = -xDelta
	}
	if yDelta < 0 {
		absYDelta = -yDelta
	}
	if absXDelta != absYDelta {
		return nil, fmt.Errorf("line between points is not at 45°")
	}

	xDir, yDir := 1, 1
	if xDelta < 0 {
		xDir = -1
	}
	if yDelta < 0 {
		yDir = -1
	}

	var points []point
	for i := 0; i <= absXDelta; i++ {
		points = append(points, point{
			x: p1.x + xDir*i,
			y: p1.y + yDir*i,
		})
	}
	return points, nil
}
