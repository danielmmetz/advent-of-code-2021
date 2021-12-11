package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	grid, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	if len(grid) != 10 {
		return 0, fmt.Errorf("expected 10 rows, got %d", len(grid))
	}
	if len(grid[0]) != 10 {
		return 0, fmt.Errorf("expected 10 columns, got %d", len(grid[0]))
	}
	var flashes int
	for i := 0; i < 100; i++ {
		grid = step(grid)
		flashes += len(flashing(grid))
		grid = reset(grid)
	}
	return flashes, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	grid, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	if len(grid) != 10 {
		return 0, fmt.Errorf("expected 10 rows, got %d", len(grid))
	}
	if len(grid[0]) != 10 {
		return 0, fmt.Errorf("expected 10 columns, got %d", len(grid[0]))
	}
	for i := 1; ; i++ {
		grid = step(grid)
		if allFlashing(grid) {
			return i, nil
		}
		grid = reset(grid)
	}
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

func step(grid [][]int) [][]int {
	var clone [][]int
	for x := 0; x < len(grid); x++ {
		var row []int
		for y := 0; y < len(grid[x]); y++ {
			row = append(row, grid[x][y]+1)
		}
		clone = append(clone, row)
	}
	flashed := map[coordinate]bool{}
	queue := flashing(clone)
	for _, p := range queue {
		flashed[p] = true
	}
	for len(queue) > 0 {
		p := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, q := range neighbors(clone, p.x, p.y) {
			clone[q.x][q.y]++
			if !flashed[q] && clone[q.x][q.y] >= 10 {
				flashed[q] = true
				queue = append(queue, q)
			}
		}
	}
	return clone
}

func reset(grid [][]int) [][]int {
	var result [][]int
	for x := 0; x < len(grid); x++ {
		var row = make([]int, 0, len(grid[0]))
		for y := 0; y < len(grid[x]); y++ {
			value := grid[x][y]
			if value >= 10 {
				value = 0
			}
			row = append(row, value)
		}
		result = append(result, row)
	}
	return result

}

func display(grid [][]int) string {
	var rows []string
	for _, row := range grid {
		var line []string
		for _, col := range row {
			if col >= 10 {
				col = 0
			}
			line = append(line, strconv.Itoa(col))
		}
		rows = append(rows, strings.Join(line, ""))
	}
	return strings.Join(rows, "\n")
}

func flashing(grid [][]int) []coordinate {
	var coords []coordinate
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] >= 10 {
				coords = append(coords, coordinate{x, y})
			}
		}
	}
	return coords
}

func allFlashing(grid [][]int) bool {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] < 10 {
				return false
			}
		}
	}
	return true
}

func neighbors(grid [][]int, x, y int) []coordinate {
	var results []coordinate
	for _, xdelta := range []int{-1, 0, 1} {
		for _, ydelta := range []int{-1, 0, 1} {
			nx, ny := x+xdelta, y+ydelta
			switch {
			case nx == x && ny == y:
				continue
			case nx < 0 || nx >= len(grid):
				continue
			case ny < 0 || ny >= len(grid[nx]):
				continue
			}
			results = append(results, coordinate{nx, ny})
		}
	}
	return results

}

type coordinate struct {
	x, y int
}

type point struct {
	coordinate
	value int
}
