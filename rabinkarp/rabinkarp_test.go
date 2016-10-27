package rabinkarp

import (
	"reflect"
	"sort"
	"testing"
)

// Struct for test data
var SearchTests = []struct {
	txt      string
	patterns []string
	expected []string
}{
	{"walter h white", []string{"white", "pinkman", "walter"}, []string{"walter", "white"}},
	{"damon salvatore", []string{"damon", "stephen", "salvatore"}, []string{"damon", "salvatore"}},
}

// run testcase
func TestSearch(t *testing.T) {
	for _, ct := range SearchTests {
		found := Search(ct.txt, ct.patterns)
		if !eq(found, ct.expected) {
			t.Errorf("Broken test .....",
				ct.txt, ct.patterns, found, ct.expected)
		}
	}
}

// sort string for comparison
func eq(f, s []string) bool {
	fx := make([]string, len(f))
	sx := make([]string, len(s))
	copy(fx, f)
	copy(sx, s)

	sort.Strings(fx)
	sort.Strings(sx)
	return reflect.DeepEqual(fx, sx)
}
