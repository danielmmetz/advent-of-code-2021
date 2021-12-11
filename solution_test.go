package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/danielmmetz/adventofcode2021/day01"
	"github.com/danielmmetz/adventofcode2021/day02"
	"github.com/danielmmetz/adventofcode2021/day03"
	"github.com/danielmmetz/adventofcode2021/day04"
	"github.com/danielmmetz/adventofcode2021/day05"
	"github.com/danielmmetz/adventofcode2021/day06"
	"github.com/danielmmetz/adventofcode2021/day07"
	"github.com/danielmmetz/adventofcode2021/day08"
	"github.com/danielmmetz/adventofcode2021/day09"
	"github.com/danielmmetz/adventofcode2021/day10"
	"github.com/danielmmetz/adventofcode2021/day11"
)

var (
	//go:embed day01/input.txt
	input01 string
	//go:embed day02/input.txt
	input02 string
	//go:embed day03/input.txt
	input03 string
	//go:embed day04/input.txt
	input04 string
	//go:embed day05/input.txt
	input05 string
	//go:embed day06/input.txt
	input06 string
	//go:embed day07/input.txt
	input07 string
	//go:embed day08/input.txt
	input08 string
	//go:embed day09/input.txt
	input09 string
	//go:embed day10/input.txt
	input10 string
	//go:embed day11/input.txt
	input11 string
)

func TestSolutions(t *testing.T) {
	cases := []struct {
		solution             interface{}
		input                string
		expected1, expected2 int
	}{
		{day01.Solution{}, input01, 7, 5},
		{day02.Solution{}, input02, 150, 900},
		{day03.Solution{}, input03, 198, 230},
		{day04.Solution{}, input04, 4512, 1924},
		{day05.Solution{}, input05, 5, 12},
		{day06.Solution{}, input06, 5934, 26984457539},
		{day07.Solution{}, input07, 37, 168},
		{day08.Solution{}, input08, 26, 61229},
		{day09.Solution{}, input09, 15, 1134},
		{day10.Solution{}, input10, 26397, 288957},
		{day11.Solution{}, input11, 1656, 195},
	}
	for i, c := range cases {
		day := i + 1
		lines := strings.Split(strings.TrimSpace(c.input), "\n")
		t.Run(fmt.Sprintf("day %d part 1", day), func(t *testing.T) {
			p1, ok := c.solution.(Part1er)
			if !ok {
				t.Fatal("part 1 not yet implemented")
			}
			got, err := p1.Part1(lines)
			if err != nil {
				t.Fatal(err)
			}
			if c.expected1 != got {
				t.Fatalf("expected %d, got %d", c.expected1, got)
			}
		})
		t.Run(fmt.Sprintf("day %d part 2", day), func(t *testing.T) {
			p2, ok := c.solution.(Part2er)
			if !ok {
				t.Fatal("part 2 not yet implemented")
			}
			got, err := p2.Part2(lines)
			if err != nil {
				t.Fatal(err)
			}
			if c.expected2 != got {
				t.Fatalf("expected %d, got %d", c.expected2, got)
			}
		})
	}
}
