package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

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
	"github.com/danielmmetz/adventofcode2021/day12"
	"github.com/danielmmetz/adventofcode2021/day13"
	"github.com/danielmmetz/adventofcode2021/day14"
)

func main() {
	day := flag.Int("day", 1, "the day's challenge to execute")
	part := flag.Int("part", 1, "the part of the day's challenge to execute")
	flag.Parse()

	if err := mainE(*day, *part); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mainE(day, part int) error {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading lines from stdin: %w", err)
	}
	var s interface{}
	switch day {
	case 1:
		s = day01.Solution{}
	case 2:
		s = day02.Solution{}
	case 3:
		s = day03.Solution{}
	case 4:
		s = day04.Solution{}
	case 5:
		s = day05.Solution{}
	case 6:
		s = day06.Solution{}
	case 7:
		s = day07.Solution{}
	case 8:
		s = day08.Solution{}
	case 9:
		s = day09.Solution{}
	case 10:
		s = day10.Solution{}
	case 11:
		s = day11.Solution{}
	case 12:
		s = day12.Solution{}
	case 13:
		s = day13.Solution{}
	case 14:
		s = day14.Solution{}
	default:
		return fmt.Errorf("not yet implemented: day %d", day)
	}

	var result int
	var err error
	switch part {
	case 1:
		p, ok := s.(Part1er)
		if !ok {
			return fmt.Errorf("not yet implemented: day %d part 1", day)
		}
		result, err = p.Part1(lines)
	case 2:
		p, ok := s.(Part2er)
		if !ok {
			return fmt.Errorf("not yet implemented: day %d part 2", day)
		}
		result, err = p.Part2(lines)
	default:
		return fmt.Errorf("unexpected part: %d", part)
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

type Part1er interface {
	Part1([]string) (int, error)
}

type Part2er interface {
	Part2([]string) (int, error)
}
