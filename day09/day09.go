package day09

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/danielmmetz/adventofcode2021/solution"
)

type Solution struct {
	solution.Solution
}

func (s Solution) Part1(lines []string) (string, error) {
	grid, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}
	var lowPoints []int
	for i := range grid {
		for j := range grid[i] {
			candidate := grid[i][j]
			lowest := true
			for _, n := range neighbors(grid, i, j) {
				if candidate >= n.value {
					lowest = false
				}
			}
			if lowest {
				lowPoints = append(lowPoints, candidate)
			}
		}
	}
	var total int
	for _, p := range lowPoints {
		total += p + 1
	}
	return fmt.Sprint(total), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	grid, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}
	var lowPoints []point
	for i := range grid {
		for j := range grid[i] {
			candidate := point{i, j, grid[i][j]}
			lowest := true
			for _, n := range neighbors(grid, i, j) {
				if candidate.value >= n.value {
					lowest = false
				}
			}
			if lowest {
				lowPoints = append(lowPoints, candidate)
			}
		}
	}
	notNine := func(p point) bool { return p.value != 9 }
	basins := make(map[point]map[point]bool)
	for _, lp := range lowPoints {
		basins[lp] = make(map[point]bool)
		queue := filter(notNine, neighbors(grid, lp.x, lp.y))
		for _, p := range queue {
			basins[lp][p] = true
		}
		for len(queue) > 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			for _, candidate := range filter(notNine, neighbors(grid, p.x, p.y)) {
				if !basins[lp][candidate] {
					basins[lp][candidate] = true
					queue = append(queue, candidate)
				}
			}
		}
	}
	if len(basins) < 3 {
		return "", fmt.Errorf("expected at least 3 basins, found %d", len(basins))
	}

	var basinSizes []int
	for _, points := range basins {
		basinSizes = append(basinSizes, len(points))
	}
	sort.Ints(basinSizes)
	n := len(basinSizes)
	total := basinSizes[n-1] * basinSizes[n-2] * basinSizes[n-3]
	return fmt.Sprint(total), nil
}

func parse(lines []string) ([][]int, error) {
	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		row := make([]int, 0, len(line))
		for _, s := range line {
			d, err := strconv.Atoi(string(s))
			if err != nil {
				return nil, fmt.Errorf("error converting %s to int: %w", string(s), err)
			}
			row = append(row, d)
		}
		grid = append(grid, row)
	}
	return grid, nil
}

func neighbors(grid [][]int, x, y int) []point {
	var results []point
	// above
	if x > 0 {
		results = append(results, point{x - 1, y, grid[x-1][y]})
	}
	// below
	if x < len(grid)-1 {
		results = append(results, point{x + 1, y, grid[x+1][y]})
	}
	// left
	if y > 0 {
		results = append(results, point{x, y - 1, grid[x][y-1]})
	}
	// right
	if y < len(grid[0])-1 {
		results = append(results, point{x, y + 1, grid[x][y+1]})
	}
	return results

}

func filter(predicate func(point) bool, points []point) []point {
	var valid []point
	for _, point := range points {
		if predicate(point) {
			valid = append(valid, point)
		}
	}
	return valid
}

type point struct {
	x, y  int
	value int
}
