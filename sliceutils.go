package sliceutils

import (
	"sort"
)

func Contains(haystack []string, needle string) bool {
	var found = false
	for _, s := range haystack {
		if s == needle {
			found = true
			break
		}
	}
	return found
}

func FindDuplicateValues(s []string) []string {
	var previous = ""
	var count = uint(0)
	var duplicateValues = []string{}
	sort.Strings(s)
	for _, value := range s {
		if value == "" {
			continue
		} else if value == previous {
			if count == 0 {
				duplicateValues = append(duplicateValues, value)
			}
			count++
		} else {
			count = 0
		}
		previous = value
	}
	return duplicateValues
}

func HasEmptyValues(s []string) bool {
	for _, value := range s {
		if value == "" {
			return true
		}
	}
	return false
}

func ListUniqueValues(s []string) []string {
	var returnSlice = []string{}
	sort.Strings(s)
	var previous = ""
	for _, value := range s {
		if value != previous {
			returnSlice = append(returnSlice, value)
			previous = value
		}
	}
	return returnSlice
}
