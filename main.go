package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/danielmmetz/adventofcode2021/day01"
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
	var s Solution
	switch day {
	case 1:
		s = day01.Solution{}
	default:
		return fmt.Errorf("not yet implemented: day %d", day)
	}

	var result string
	var err error
	switch part {
	case 1:
		result, err = s.Part1(lines)
	case 2:
		result, err = s.Part2(lines)
	default:
		return fmt.Errorf("unexpected part: %d", part)
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

type Solution interface {
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}
