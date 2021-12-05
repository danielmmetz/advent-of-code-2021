package day05

import (
	"reflect"
	"testing"
)

func input() []string {
	return []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "5" {
		t.Fatalf("expected 5, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "12" {
		t.Fatalf("expected 12 got %s\n", result)
	}
}

func TestDiagonalPointsBetween(t *testing.T) {
	ps, err := diagonalPointsBetween(point{1, 1}, point{3, 3})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ps, []point{{1, 1}, {2, 2}, {3, 3}}) {
		t.Fatalf("incorrect diagonal points between 1,1 and 3,3; got %v", ps)
	}
	ps, err = diagonalPointsBetween(point{9, 7}, point{7, 9})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ps, []point{{9, 7}, {8, 8}, {7, 9}}) {
		t.Fatalf("incorrect diagonal points between 9,7 and 7,9; got %v", ps)
	}
}
