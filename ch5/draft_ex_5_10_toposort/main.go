package main

import (
	"fmt"
	"sort"
)

// var prereqs = map[string][]string{
// 	"algorithms":            {"data structures"},
// 	"calculus":              {"linear algebra"},
// 	"compilers":             {"data structures", "formal languages", "computer organization"},
// 	"data structures":       {"discrete math"},
// 	"databases":             {"data structures"},
// 	"discrete math":         {"intro to programming"},
// 	"formal languages":      {"discrete math"},
// 	"networks":              {"operating systems"},
// 	"operating systems":     {"data structures", "computer organization"},
// 	"programming languages": {"data structures", "computer organization"},
// }

var prereqs = map[string]map[string]bool{
	"algorithms":            {"data structures": true},
	"calculus":              {"linear algebra": true},
	"compilers":             {"data structures": true, "formal languages": true, "computer organization": true},
	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) (order []string) {
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
