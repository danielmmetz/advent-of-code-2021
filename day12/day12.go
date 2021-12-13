package day12

import (
	"fmt"
	"strings"
)

type Solution struct {
}

func (s Solution) Part1(lines []string) (int, error) {
	paths, err := solution(lines, canVisit1)
	return len(paths), err
}

func (s Solution) Part2(lines []string) (int, error) {
	paths, err := solution(lines, canVisit2)
	return len(paths), err
}

func solution(lines []string, canVisit func(path, string) bool) ([]path, error) {
	adjacency := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("expected 2 parts for line, got %d: %s", len(parts), line)
		}
		a, b := parts[0], parts[1]
		adjacency[a] = append(adjacency[a], b)
		adjacency[b] = append(adjacency[b], a)
	}
	seen := make(map[string]bool)
	incompletePaths := []path{{"end"}}
	for len(incompletePaths) > 0 {
		candidate := incompletePaths[len(incompletePaths)-1]
		incompletePaths = incompletePaths[:len(incompletePaths)-1]
		for _, a := range adjacency[candidate[len(candidate)-1]] {
			if !canVisit(candidate, a) {
				continue
			}
			extended := candidate.Extend(a)
			if seen[extended.String()] {
				continue
			}
			seen[extended.String()] = true
			if a == "start" {
				continue
			}
			incompletePaths = append(incompletePaths, extended)
		}
	}
	var complete []path
	for p := range seen {
		if isComplete(newPath(p)) {
			complete = append(complete, reverse(newPath(p)))
		}
	}
	return complete, nil

}

func canVisit1(soFar path, candidate string) bool {
	history := make(map[string]bool)
	for _, component := range soFar {
		history[component] = true
	}
	switch candidate {
	case "end":
		return false
	case "start":
		return true
	case strings.ToUpper(candidate):
		return true
	}
	return !history[candidate]
}

func canVisit2(soFar path, candidate string) bool {
	history := make(map[string]int)
	var sameLowerSeenCount int
	for _, component := range soFar {
		history[component]++
		switch component {
		case "start", "end":
			continue
		case strings.ToLower(component):
			if sameLowerSeenCount < history[component] {
				sameLowerSeenCount = history[component]
			}
		}
	}
	switch candidate {
	case "end":
		return false
	case "start":
		return true
	case strings.ToUpper(candidate):
		return true
	}
	if history[candidate] == 0 {
		return true
	}
	return sameLowerSeenCount < 2
}

func reverse(elems path) path {
	result := make(path, 0, len(elems))
	for i := 0; i < len(elems); i++ {
		result = append(result, elems[len(elems)-i-1])
	}
	return result
}

type path []string

func (p path) String() string {
	return strings.Join(p, ",")
}

func (p path) Extend(n string) path {
	clone := make([]string, 0, len(p)+1)
	for _, component := range p {
		clone = append(clone, component)
	}
	clone = append(clone, n)
	return clone
}

func newPath(s string) path {
	return strings.Split(s, ",")
}

func isComplete(p path) bool {
	if len(p) < 2 {
		return false
	}
	return p[0] == "end" && p[len(p)-1] == "start"
}

func hasPrefix(p path, prefix path) bool {
	if len(prefix) > len(p) {
		return false
	}
	for i := range prefix {
		if prefix[i] != p[i] {
			return false
		}
	}
	return true
}
