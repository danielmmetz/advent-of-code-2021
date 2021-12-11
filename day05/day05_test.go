package day05

import (
	"reflect"
	"testing"
)

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
