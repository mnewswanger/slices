package slices

import (
	"reflect"
	"testing"
)

var testSlices = map[*[]interface{}]struct {
	name       string
	contains   map[interface{}]bool
	duplicates map[interface{}]uint
	unique     []interface{}
}{
	&[]interface{}{}: {
		name: "Empty Slice",
		contains: map[interface{}]bool{
			"test": false,
			0:      false,
			"":     false,
		},
		duplicates: map[interface{}]uint{},
		unique:     []interface{}{},
	},
	&[]interface{}{
		"test",
		"Test",
		"TEST",
	}: {
		name: "Case Sensivity",
		contains: map[interface{}]bool{
			"test": true,
			"Test": true,
			"TEST": true,
			"TEst": false,
		},
		duplicates: map[interface{}]uint{},
		unique: []interface{}{
			"test",
			"Test",
			"TEST",
		},
	},
	&[]interface{}{
		0,
		"test",
		"Test",
		"TEST",
	}: {
		name: "Mismatched Types",
		contains: map[interface{}]bool{
			0:     true,
			"0":   false,
			-1:    false,
			false: false,
			nil:   false,
		},
		duplicates: map[interface{}]uint{},
		unique: []interface{}{
			0,
			"test",
			"Test",
			"TEST",
		},
	},
	&[]interface{}{
		"test",
		"test",
		"Test",
		"TEST",
	}: {
		name: "Duplicate Strings",
		contains: map[interface{}]bool{
			"test": true,
		},
		duplicates: map[interface{}]uint{
			"test": 2,
		},
		unique: []interface{}{
			"test",
			"Test",
			"TEST",
		},
	},
}

func TestContains(t *testing.T) {
	for sliceToTest, testDesiredResults := range testSlices {
		for item, itemExistsInSlice := range testDesiredResults.contains {
			if Contains(*sliceToTest, item) != itemExistsInSlice {
				if itemExistsInSlice {
					t.Error(item, "not found but exists in slice")
				} else {
					t.Error(item, "found but does not exist in slice")
				}
			}
		}
	}
}

func TestFindDuplicateValues(t *testing.T) {
	for s, m := range testSlices {
		var got = FindDuplicateValues(*s)
		if !reflect.DeepEqual(got, m.duplicates) {
			t.Error("Find Duplicate Value failed: " + m.name)
			println()
		}
	}
}

func TestListDistinctValues(t *testing.T) {
	for s, m := range testSlices {
		var got = ListDistinctValues(*s)

		if !reflect.DeepEqual(got, m.unique) {
			t.Error("Find Distinct Values failed: " + m.name)
		}
	}
}
