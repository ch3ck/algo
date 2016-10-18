// Package knuthMorrisPratt provides a simple Match function to find occurrences
// of a pattern in a block of text
package knuthMorrisPratt

// Match uses the Knuth-Morris-Pratt algorithm to find occurrences of `pattern` in `text`
// and returns the list of indexes where the matches are in `text`
func Match(text, pattern string) (result []int) {
	n := len(text)
	m := len(pattern)

	if n == 0 || m == 0 {
		return
	}

	pi := prefix(pattern)
	q := 0

	for i := 0; i < n; i++ {
		for q > 0 && pattern[q] != text[i] {
			q = pi[q-1]
		}
		if pattern[q] == text[i] {
			q++
		}
		if q == m {
			result = append(result, i-m+1)
			q = pi[q-1]
		}
	}

	return
}

// prefix function, pi, encapsulates knowledege about how a pattern matches againsts
// shifts of itself.
func prefix(pattern string) []int {
	m := len(pattern)
	if m == 0 {
		return []int{}
	}

	pi := make([]int, m)
	pi[0] = 0
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = pi[k-1]
		}
		if pattern[k] == pattern[q] {
			k++
		}
		pi[q] = k
	}

	return pi
}
