package slices

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

var stringTestSets = map[*[]string]struct {
	contains   map[string]bool
	duplicates map[string]uint
	unique     []string
}{
	// Empty slice
	&[]string{}: {
		contains: map[string]bool{
			"":       false,
			"string": false,
			"0":      false,
		},
		duplicates: map[string]uint{},
		unique:     []string{},
	},
	// Duplicate element
	&[]string{
		"test",
		"test",
		"test",
	}: {
		contains: map[string]bool{
			"":       false,
			"string": false,
			"0":      false,
			"test":   true,
		},
		duplicates: map[string]uint{
			"test": 3,
		},
		unique: []string{
			"test",
		},
	},
	// Case sensitivity with multiple duplicates
	&[]string{
		"Test",
		"Test",
		"Test",
		"Test",
		"test",
		"test",
		"1",
		"2",
		"3",
	}: {
		contains: map[string]bool{
			"":       false,
			"string": false,
			"test":   true,
			"Test":   true,
			"TEST":   false,
			"0":      false,
			"1":      true,
			"2":      true,
		},
		duplicates: map[string]uint{
			"Test": 4,
			"test": 2,
		},
		unique: []string{
			"Test",
			"test",
			"1",
			"2",
			"3",
		},
	},
}

func TestContainString(t *testing.T) {
	for s, m := range stringTestSets {
		for c, v := range m.contains {
			if ContainsString(*s, c) != v {
				if v {
					t.Error("Could not find string \"" + c + "\" but should have")
				} else {
					t.Error("Found string \"" + c + "\" but should not have")
				}
			}
		}
	}
}

func TestFindDuplicateStrings(t *testing.T) {
	for s, m := range stringTestSets {
		var got = FindDuplicateStrings(*s)
		if !reflect.DeepEqual(got, m.duplicates) {
			t.Error("Find Duplicate Strings failed:")
			println("Got:")
			for v, c := range got {
				println("  " + v + " (" + strconv.Itoa(int(c)) + ")")
			}
			println("Want:")
			for v, c := range m.duplicates {
				println("  " + v + " (" + strconv.Itoa(int(c)) + ")")
			}
			println()
		}
	}
}

func TestListDistinctStrings(t *testing.T) {
	for s, m := range stringTestSets {
		var got = ListDistinctStrings(*s)
		sort.Strings(got)
		sort.Strings(m.unique)

		if !reflect.DeepEqual(got, m.unique) {
			t.Error("Find Distinct Strings failed:")
			println("Got:")
			for _, v := range got {
				println("  " + v)
			}
			println("Want:")
			for _, v := range m.unique {
				println("  " + v)
			}
			println()
		}
	}
}
