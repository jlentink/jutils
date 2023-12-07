package jslices

import "strings"

// Contains checks if a type is in a slice of type
func Contains[T comparable](list []T, v T) bool {
	for _, i := range list {
		if i == v {
			return true
		}
	}
	return false
}

func ContainsEqualFold(list []string, v string) bool {
	for _, i := range list {
		if strings.EqualFold(i, v) {
			return true
		}
	}
	return false
}
