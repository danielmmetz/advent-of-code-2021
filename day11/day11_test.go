package day11

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

const (
	expected1 = `
6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637`

	expected2 = `
8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848`

	expected3 = `
0050900866
8500800575
9900000039
9700000041
9935080063
7712300000
7911250009
2211130000
0421125000
0021119000`
)

func TestSmallStep(t *testing.T) {
	start := `
11111
19991
19191
19991
11111`

	expected1 := `
34543
40004
50005
40004
34543`

	expected2 := `
45654
51115
61116
51115
45654`
	grid := mustParse(t, start)
	grid = reset(step(grid))
	ok, msg := equal(mustParse(t, expected1), grid)
	if !ok {
		t.Fatalf("step 1: %v", msg)
	}
	grid = reset(step(grid))
	ok, msg = equal(mustParse(t, expected2), grid)
	if !ok {
		t.Fatalf("step 2: %v", msg)
	}
}

func TestStep(t *testing.T) {
	grid, err := parse(strings.Split(input, "\n"))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	grid = reset(step(grid))
	ok, msg := equal(mustParse(t, expected1), grid)
	if !ok {
		t.Fatalf("step 1: %v", msg)
	}
	grid = reset(step(grid))
	ok, msg = equal(mustParse(t, expected2), grid)
	if !ok {
		t.Fatalf("step 2: %v", msg)
	}
	grid = reset(step(grid))
	ok, msg = equal(mustParse(t, expected3), grid)
	if !ok {
		t.Fatalf("step 3: %v", msg)
	}
}

func mustParse(t *testing.T, grid string) [][]int {
	result, err := parse(strings.Split(strings.TrimSpace(grid), "\n"))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return result
}

func equal(expected, got [][]int) (bool, string) {
	for x := 0; x < len(expected); x++ {
		for y := 0; y < len(expected[x]); y++ {
			e, g := expected[x][y], got[x][y]
			if e != g {
				return false, fmt.Sprintf("mismatch at (%d, %d): expected %d got %d", x, y, e, g)
			}
		}
	}
	return true, ""
}
