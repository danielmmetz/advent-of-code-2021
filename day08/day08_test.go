package day08

import (
	"sort"
	"testing"
)

func TestNormalize(t *testing.T) {
	entries, err := parse([]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"})
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}

	entry, err := normalize(entries[0])
	if err != nil {
		t.Fatal(err)
	}
	firstPattern := keys(entry.patterns[0])
	sort.Slice(firstPattern, func(i, j int) bool {
		return firstPattern[i] < firstPattern[j]
	})
	if string(firstPattern) != "abcdefg" {
		t.Fatalf("expected first normalized pattern to be 'abcdefg', got %s", string(firstPattern))
	}
}
