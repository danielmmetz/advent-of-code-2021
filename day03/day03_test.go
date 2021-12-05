package day03

import "testing"

func input() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "198" {
		t.Fatalf("expected 198, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "230" {
		t.Fatalf("expected 230, got %s\n", result)
	}
}
