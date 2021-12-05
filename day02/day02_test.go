package day02

import "testing"

func input() []string {
	return []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "150" {
		t.Fatalf("expected 150, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "900" {
		t.Fatalf("expected 900, got %s\n", result)
	}
}
