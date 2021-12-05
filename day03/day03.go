package day03

import (
	"fmt"
	"strconv"

	"github.com/danielmmetz/adventofcode2021/solution"
)

type Solution struct {
	solution.Solution
}

func (s Solution) Part1(lines []string) (string, error) {
	var numBits int
	for i, line := range lines {
		if i == 0 {
			numBits = len(line)
			continue
		}
		if len(line) != numBits {
			return "", fmt.Errorf("expected lines of equal bit lengths: not true for line %d", i)
		}
	}
	var gamma, epsilon int
	for pos := 0; pos < numBits; pos++ {
		gamma = gamma << 1
		epsilon = epsilon << 1

		var zeros, ones int
		for i, line := range lines {
			switch line[pos] {
			case '0':
				zeros++
			case '1':
				ones++
			default:
				return "", fmt.Errorf("invalid bit in line %d: %s", i, line)
			}
		}
		switch {
		case zeros > ones:
			epsilon += 1
		case ones > zeros:
			gamma += 1
		default:
			return "", fmt.Errorf("equal number of zeros and ones for bit position %d", pos)
		}
	}

	return fmt.Sprint(gamma * epsilon), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	var numBits int
	for i, line := range lines {
		if i == 0 {
			numBits = len(line)
			continue
		}
		if len(line) != numBits {
			return "", fmt.Errorf("expected lines of equal bit lengths: not true for line %d", i)
		}
	}
	oxygenLines := lines
	for pos := 0; pos < numBits && len(oxygenLines) > 1; pos++ {
		var zeros, ones int
		var zeroLines, oneLines []string
		for i, line := range oxygenLines {
			switch line[pos] {
			case '0':
				zeros++
				zeroLines = append(zeroLines, line)
			case '1':
				ones++
				oneLines = append(oneLines, line)
			default:
				return "", fmt.Errorf("invalid bit in line %d: %s", i, line)
			}
		}
		if zeros > ones {
			oxygenLines = zeroLines
		} else {
			oxygenLines = oneLines
		}
	}

	c02Lines := lines
	for pos := 0; pos < numBits && len(c02Lines) > 1; pos++ {
		var zeros, ones int
		var zeroLines, oneLines []string
		for i, line := range c02Lines {
			switch line[pos] {
			case '0':
				zeros++
				zeroLines = append(zeroLines, line)
			case '1':
				ones++
				oneLines = append(oneLines, line)
			default:
				return "", fmt.Errorf("invalid bit in line %d: %s", i, line)
			}
		}
		if zeros <= ones {
			c02Lines = zeroLines
		} else {
			c02Lines = oneLines
		}
	}

	if len(oxygenLines) != 1 {
		return "", fmt.Errorf("expected to have whittled down to a single oxygen line: %v", oxygenLines)
	}
	if len(c02Lines) != 1 {
		return "", fmt.Errorf("expected to have whittled down to a single oxygen line: %v", c02Lines)
	}
	oxygen, err := strconv.ParseInt(oxygenLines[0], 2, 64)
	if err != nil {
		return "", fmt.Errorf("error converting %s to int: %w", oxygenLines[0], err)
	}
	c02, err := strconv.ParseInt(c02Lines[0], 2, 64)
	if err != nil {
		return "", fmt.Errorf("error converting %s to int: %w", c02Lines[0], err)
	}

	return fmt.Sprint(oxygen * c02), nil
}
