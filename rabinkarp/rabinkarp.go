
package rabinkarp

const base = 2048

// Search searches given patterns in txt and returns the matched ones. Returns
// empty string slice if there is no match.
func Search(txt string, patterns []string) []string {
	in := Indices(txt, patterns)
	matches := make([]string, len(in))
	i := 0
	for j, p := range patterns {
		if _, ok := in[j]; ok {
			matches[i] = p
			i++
		}
	}

	return matches
}

// Indices returns the Indices of the first occurence of each pattern in txt.
func Indices(txt string, patterns []string) map[int]int {
	n, m := len(txt), MinLen(patterns)
	matches := make(map[int]int)

	if n < m || len(patterns) == 0 {
		return matches
	}

	var mult uint32 = 1 // mult = base^(m-1)
	for i := 0; i < m-1; i++ {
		mult = (mult * base)
	}

	hp := HashPatterns(patterns, m)
	h := Hash(txt[:m])
	for i := 0; i < n-m+1 && len(hp) > 0; i++ {
		if i > 0 {
			h = h - mult*uint32(txt[i-1])
			h = h*base + uint32(txt[i+m-1])
		}

		if mps, ok := hp[h]; ok {
			for _, pi := range mps {
				pat := patterns[pi]
				e := i + len(pat)
				if _, ok := matches[pi]; !ok && e <= n && pat == txt[i:e] {
					matches[pi] = i
				}
			}
		}
	}
	return matches
}

func Hash(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = (h*base + uint32(s[i]))
	}
	return h
}

func HashPatterns(patterns []string, l int) map[uint32][]int {
	m := make(map[uint32][]int)
	for i, t := range patterns {
		h := Hash(t[:l])
		if _, ok := m[h]; ok {
			m[h] = append(m[h], i)
		} else {
			m[h] = []int{i}
		}
	}

	return m
}

func MinLen(patterns []string) int {
	if len(patterns) == 0 {
		return 0
	}

	m := len(patterns[0])
	for i := range patterns {
		if m > len(patterns[i]) {
			m = len(patterns[i])
		}
	}

	return m
}