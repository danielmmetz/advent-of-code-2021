package day08

import (
	"fmt"
	"sort"
	"strings"
)

type Solution struct{}

func (s Solution) Part1(lines []string) (int, error) {
	entries, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	uniqueLengths := map[int]bool{
		2: true,
		3: true,
		4: true,
		7: true,
	}
	var count int
	for _, entry := range entries {
		for _, digit := range entry.outputDigits {
			if uniqueLengths[len(digit)] {
				count++
			}
		}
	}
	return count, nil
}

func (s Solution) Part2(lines []string) (int, error) {
	entries, err := parse(lines)
	if err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}
	var normalizedEntries []entry
	for i, entry := range entries {
		normalized, err := normalize(entry)
		if err != nil {
			return 0, fmt.Errorf("error normalizing entry %d: %w", i, err)
		}
		normalizedEntries = append(normalizedEntries, normalized)
	}

	var total int
	for _, entry := range normalizedEntries {
		var number int
		for _, pattern := range entry.outputDigits {
			number = number*10 + pattern.Digit()
		}
		total += number
	}

	return total, nil
}

func parse(lines []string) ([]entry, error) {
	var entries []entry
	for i, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 15 {
			return nil, fmt.Errorf("line %d: expected 15 fields, got %d", i, len(line))
		}
		var patterns []pattern
		for _, pattern := range append(parts[:10], parts[11:]...) {
			p := make(map[rune]bool)
			for _, letter := range pattern {
				p[letter] = true
			}
			patterns = append(patterns, p)
		}
		entries = append(entries, entry{
			patterns:     patterns[:10],
			outputDigits: patterns[10:],
		})
	}
	return entries, nil
}

type entry struct {
	patterns     []pattern
	outputDigits []pattern
}

// normalize converts a given entry to an equivalent entry that uses the reference mapping.
// Internally, it uses the following strategy:
// 1. 'cf' can be ambiguously determined by use of the pattern of length 2.
// 2. 'a' can be determined by use of the pattern of length 3, minus 'cf'.
// 3. 'c' can be determined by looking at the runes that appear 8 times, minus 'a'.
// 4. 'b' can be determined by looking at the rune that appears in 6 patterns.
// 5. 'e' can be determined by looking at the rune that appears in 4 patterns.
// 6. 'f' can be determined by looking at the rune that appears in 9 patterns.
// 7. 'd' is the the pattern of length 4 (bcdf) minus the resovlved b, c, and f.
// 8. 'g' is now determined by elimination.
func normalize(input entry) (normalized entry, err error) {
	defer func() {
		// The error checking ordinarily required for the resolution below
		// is so involved that in this rare instance we prefer to panic and recover
		// rather than add noise and further obfuscate the complexity entailed.
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	runeFrequency := make(map[rune]int)
	lengthToPatterns := make(map[int][]pattern)
	for _, pattern := range input.patterns {
		l := len(pattern)
		lengthToPatterns[l] = append(lengthToPatterns[l], pattern)
		for r := range pattern {
			runeFrequency[r]++
		}
	}
	frequencyToRunes := make(map[int][]rune)
	for k, v := range runeFrequency {
		frequencyToRunes[v] = append(frequencyToRunes[v], k)
	}

	// step 1
	cf := lengthToPatterns[2][0]
	acf := lengthToPatterns[3][0]

	// step 2
	a := keys(mapMinus(acf, cf))[0]

	// step 3
	eightFrequencyRunes := frequencyToRunes[8]
	c := remove(eightFrequencyRunes, a)[0]

	// step 4, 5, 6
	b := frequencyToRunes[6][0]
	e := frequencyToRunes[4][0]
	f := frequencyToRunes[9][0]

	// step 7
	bcdf := lengthToPatterns[4][0]
	d := remove(keys(bcdf), b, c, f)[0]

	// step 8
	g := remove([]rune("abcdefg"), a, b, c, d, e, f)[0]

	inputToReference := map[rune]rune{
		a: 'a',
		b: 'b',
		c: 'c',
		d: 'd',
		e: 'e',
		f: 'f',
		g: 'g',
	}

	normalized = entry{}
	for _, p := range input.patterns {
		translated := make(pattern)
		for k, v := range p {
			translated[inputToReference[k]] = v
		}
		normalized.patterns = append(normalized.patterns, translated)
	}
	for _, p := range input.outputDigits {
		translated := make(pattern)
		for k, v := range p {
			translated[inputToReference[k]] = v
		}
		normalized.outputDigits = append(normalized.outputDigits, translated)
	}
	return
}

type pattern map[rune]bool

func (p pattern) String() string {
	rs := keys(p)
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return string(rs)
}

/*
Reference Display
    aaaa
   b    c
   b    c
    dddd
   e    f
   e    f
    gggg
*/
func (p pattern) Digit() int {
	return map[string]int{
		// length 2
		"cf": 1,
		// length 3
		"acf": 7,
		// length 4
		"bcdf": 4,
		// length 6
		"acdeg": 2,
		"acdfg": 3,
		"abdfg": 5,
		// length 7
		"abcdfg": 9,
		"abcefg": 0,
		"abdefg": 6,
		// length 7
		"abcdefg": 8,
	}[p.String()]
}

func keys(m map[rune]bool) []rune {
	result := make([]rune, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func mapMinus(a, b map[rune]bool) map[rune]bool {
	result := make(map[rune]bool)
	for k := range a {
		if _, ok := b[k]; !ok {
			result[k] = true
		}
	}
	return result
}

func remove(original []rune, omissions ...rune) []rune {
	omissionMap := make(map[rune]bool)
	for _, o := range omissions {
		omissionMap[o] = true
	}
	var result []rune
	for _, r := range original {
		if !omissionMap[r] {
			result = append(result, r)
		}
	}
	return result
}
