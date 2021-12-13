package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	points, err := solve(lines, 1)
	return len(points), err
}

func (s Solution) Part2(lines []string) (int, error) {
	ps, err := solve(lines, 0)
	return len(ps), err
}

func solve(lines []string, maxFolds int) (points, error) {
	ps, folds, err := parse(lines)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}
	if len(folds) < maxFolds {
		return nil, fmt.Errorf("expected at least %d folds, got %d", maxFolds, len(folds))
	}
	if maxFolds > 0 {
		folds = folds[:maxFolds]
	}
	for _, fold := range folds {
		switch fold.axis {
		case "x":
			ps = foldXAt(ps, fold.value)
		case "y":
			ps = foldYAt(ps, fold.value)
		}
	}
	return ps, nil
}

func parse(lines []string) (points, []fold, error) {
	ps := make(map[point]bool)
	var folds []fold

	var encounteredBlank bool
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			encounteredBlank = true
			continue
		}
		switch encounteredBlank {
		case false:
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("expected 2 parts for line: %s", line)
			}
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, fmt.Errorf("expected int, got %s: %w", parts[0], err)
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, fmt.Errorf("expected int, got %s: %w", parts[1], err)
			}
			ps[point{x, y}] = true
		case true:
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, nil, fmt.Errorf("expected 3 parts for line: %s", line)
			}
			parts = strings.Split(parts[2], "=")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, fmt.Errorf("expected int, got %s: %w", parts[1], err)
			}
			switch parts[0] {
			case "x", "y":
				folds = append(folds, fold{axis: parts[0], value: value})
			default:
				return nil, nil, fmt.Errorf("expected axis x or y, got %s", parts[0])
			}
		}
	}
	return ps, folds, nil
}

type points map[point]bool

func (ps points) String() string {
	var maxX, maxY int
	for p := range ps {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	var grid [][]rune
	for i := 0; i < maxY+1; i++ {
		grid = append(grid, []rune{})
		for j := 0; j < maxX+1; j++ {
			grid[i] = append(grid[i], '.')
		}
	}
	for p := range ps {
		grid[p.y][p.x] = '#'
	}
	var s strings.Builder
	for _, row := range grid {
		for _, c := range row {
			s.WriteRune(c)
		}
		s.WriteRune('\n')
	}
	return s.String()
}

type point struct {
	x, y int
}

type fold struct {
	axis  string
	value int
}

func foldYAt(ps points, value int) points {
	results := make(map[point]bool)
	for p := range ps {
		if p.y < value {
			results[p] = true
			continue
		}
		results[point{x: p.x, y: value - (p.y - value)}] = true
	}
	return results
}

func foldXAt(ps points, value int) points {
	results := make(map[point]bool)
	for p := range ps {
		if p.x < value {
			results[p] = true
			continue
		}
		results[point{x: value - (p.x - value), y: p.y}] = true
	}
	return results
}
