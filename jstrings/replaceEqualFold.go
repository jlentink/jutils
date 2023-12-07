package jstrings

import (
	"github.com/jlentink/jutils/jslices"
	"strings"
)

// ReplaceEqualFold replaces all occurrences of old with new not case sensitively.
func ReplaceEqualFold(s, old, new string, n int) string {
	if old == new || n == 0 {
		return s // avoid allocation
	}
	sl := strings.ToLower(s)
	ol := strings.ToLower(old)

	// Find first match if any.
	index := strings.Index(sl, ol)
	if index == -1 {
		return s
	}

	// Count the number of matches.
	var matches int
	pos := []int{index}
	for i := index; i != -1; {
		matches++
		if n > 0 && matches == n {
			break
		}
		i = strings.Index(sl[i+len(ol):], ol)
		if i != -1 {
			i += pos[len(pos)-1] + len(ol)
			pos = append(pos, i)
		}
	}
	tl := len(old)
	output := ""
	for i := 0; i < len(s); i++ {
		if jslices.Contains(pos, i) {
			output += new
			i += tl
			if i >= len(s) {
				break
			}
		}
		output += string(s[i])
	}
	return output
}
