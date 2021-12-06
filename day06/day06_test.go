package day06

import (
	"testing"
)

func input() []string {
	return []string{"3,4,3,1,2"}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "5934" {
		t.Fatalf("expected 5934, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "26984457539" {
		t.Fatalf("expected 26984457539 got %s\n", result)
	}
}
