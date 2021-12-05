package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielmmetz/adventofcode2021/solution"
)

type Solution struct {
	solution.Solution
}

func (s Solution) Part1(lines []string) (string, error) {
	commands, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}
	var x, y int
	for i, c := range commands {
		switch c.direction {
		case "up":
			y -= c.magnitude
		case "down":
			y += c.magnitude
		case "forward":
			x += c.magnitude
		default:
			return "", fmt.Errorf("unexpected direction for command %d: %s", i, c.direction)
		}
	}
	return fmt.Sprint(x * y), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	commands, err := parse(lines)
	if err != nil {
		return "", fmt.Errorf("parse error: %w", err)
	}
	var x, y, aim int
	for i, c := range commands {
		switch c.direction {
		case "up":
			aim -= c.magnitude
		case "down":
			aim += c.magnitude
		case "forward":
			x += c.magnitude
			y += c.magnitude * aim
		default:
			return "", fmt.Errorf("unexpected direction for command %d: %s", i, c.direction)
		}
	}
	return fmt.Sprint(x * y), nil
}

type command struct {
	direction string
	magnitude int
}

func parse(lines []string) ([]command, error) {
	commands := make([]command, 0, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected format for line %d: %s", i, line)
		}
		magnitude, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("unexpected magnitude for line %d: %s: %w", i, parts[1], err)
		}
		commands = append(commands, command{parts[0], magnitude})
	}
	return commands, nil
}
