package helpers

// kang from: https://github.com/kmaher9/levenshtein/blob/master/levenshtein.go
import (
	"unicode/utf8"
)

// StrCompare - Compare similarity of 2 strings.
func StrCompare(a, b string) float64 {
	distance := computeLevenshteinValue(a, b)
	length := calculateLongestWord(a, b)

	return (1.00 - float64(distance)/float64(length)) * 100.00
}
func computeLevenshteinValue(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)

	for j := range f {
		f[j] = j
	}

	for _, ca := range a {
		j := 1
		fj1 := f[0]
		f[0]++
		for _, cb := range b {
			mn := min(f[j]+1, f[j-1]+1)
			if cb != ca {
				mn = min(mn, fj1+1)
			} else {
				mn = min(mn, fj1)
			}

			fj1, f[j] = f[j], mn
			j++
		}
	}

	return f[len(f)-1]
}

func calculateLongestWord(a, b string) int {
	if len(a) >= len(b) {
		return len(a)
	}

	return len(b)
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
