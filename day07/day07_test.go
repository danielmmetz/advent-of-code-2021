package day07

import (
	"testing"
)

func input() []string {
	return []string{"16,1,2,0,4,2,7,1,2,14"}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "37" {
		t.Fatalf("expected 37, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "168" {
		t.Fatalf("expected 168 got %s\n", result)
	}
}
