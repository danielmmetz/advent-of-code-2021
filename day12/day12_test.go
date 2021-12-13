package day12

import "testing"

func TestCanVisit1(t *testing.T) {
	cases := []struct {
		soFar     string
		candidate string
		expected  bool
	}{
		{"end,A,b", "A", true},
		{"end,A,b", "b", false},
		{"end,A,b", "start", true},
		{"end,A,b", "end", false},
		{"end,A,b", "c", true},
	}
	for _, c := range cases {
		t.Run(c.soFar, func(t *testing.T) {
			result := canVisit1(newPath(c.soFar), c.candidate)
			if result != c.expected {
				t.Errorf("%s to visit %s: expected %v, got %v\n", c.soFar, c.candidate, c.expected, result)
			}
		})
	}
}
