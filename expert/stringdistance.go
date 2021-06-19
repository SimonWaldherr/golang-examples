package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func invBool2int(b bool) int {
	if !b {
		return 1
	}
	return 0
}

func LevenshteinDistance(a, b string) int {
	len1, len2 := len(a), len(b)
	if len1 < len2 {
		return LevenshteinDistance(b, a)
	}
	row1, row2 := make([]int, len2+1), make([]int, len2+1)

	for i := 0; i < len2+1; i++ {
		row2[i] = i
	}

	for i := 0; i < len1; i++ {
		row1[0] = i + 1

		for j := 0; j < len2; j++ {
			x := min(row2[j+1]+1, row1[j]+1)
			y := row2[j] + invBool2int(a[i] == b[j])
			row1[j+1] = min(x, y)
		}

		row1, row2 = row2, row1
	}
	return row2[len2]
}

func DamerauLevenshteinDistance(a, b string) int {
	len1, len2 := len(a), len(b)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	if len1 < len2 {
		return DamerauLevenshteinDistance(b, a)
	}
	curr, next := 0, 0
	row := make([]int, len2+1)

	for i := 0; i < len2+1; i++ {
		row[i] = i
	}

	for i := 0; i < len1; i++ {
		curr = i + 1

		for j := 0; j < len2; j++ {
			cost := invBool2int(a[i] == b[j] || (i > 0 && j > 0 && a[i-1] == b[j] && a[i] == b[j-1]))

			next = min(min(
				row[j+1]+1,
				row[j]+cost),
				curr+1)

			row[j], curr = curr, next
		}
		row[len2] = next
	}
	return next
}

func main() {
	fmt.Printf("LevenshteinDistance(\"foobar\", \"fubar\"): %v\n", LevenshteinDistance("foobar", "fubar"))
	fmt.Printf("DamerauLevenshteinDistance(\"foobar\", \"fubar\"): %v\n", DamerauLevenshteinDistance("foobar", "fubar"))
}
