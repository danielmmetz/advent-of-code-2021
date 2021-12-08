package day08

import (
	"sort"
	"testing"
)

func input() []string {
	return []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
}

func TestPart1(t *testing.T) {
	var s Solution
	result, err := s.Part1(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "26" {
		t.Fatalf("expected 26, got %s\n", result)
	}
}

func TestPart2(t *testing.T) {
	var s Solution
	result, err := s.Part2(input())
	if err != nil {
		t.Fatal(err)
	}
	if result != "61229" {
		t.Fatalf("expected 61229 got %s\n", result)
	}
}

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
