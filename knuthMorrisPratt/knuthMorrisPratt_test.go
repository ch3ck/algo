package knuthMorrisPratt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchBlankPattern(t *testing.T) {
	result := Match("text", "")
	assert.Empty(t, result)
}

func TestMatchBlankText(t *testing.T) {
	result := Match("", "pattern")
	assert.Empty(t, result)
}

func TestMatchBlank(t *testing.T) {
	result := Match("", "")
	assert.Empty(t, result)
}

func TestMatchSimple(t *testing.T) {
	result := Match("abc", "b")
	assert.Equal(t, 1, len(result))
	assert.Contains(t, result, 1)
}

func TestMatchMultiple(t *testing.T) {
	result := Match("abcb", "b")
	assert.Equal(t, 2, len(result))
	assert.Contains(t, result, 1, 3)
}

func TestMatchLongPattern(t *testing.T) {
	result := Match("abcb", "cb")
	assert.Equal(t, 1, len(result))
	assert.Contains(t, result, 2)
}

func TestMatchLongPatternMultiple(t *testing.T) {
	result := Match("acbcacb", "cb")
	assert.Equal(t, 2, len(result))
	assert.Contains(t, result, 1, 5)
}

func TestMatchComplex(t *testing.T) {
	result := Match("acbcabcababcbabcaacb", "abcababcbabca")
	assert.Equal(t, 1, len(result))
	assert.Contains(t, result, 4)
}

func TestMatchNoMatch(t *testing.T) {
	result := Match("acbcacb", "xyz")
	assert.Empty(t, result)
}

func TestPrefixBlank(t *testing.T) {
	result := prefix("")
	assert.Empty(t, result)
}

func TestPrefixSimple(t *testing.T) {
	result := prefix("ab")
	assert.Equal(t, []int{0, 0}, result)
}

func TestPrefixLong(t *testing.T) {
	result := prefix("abcdefghij")
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, result)
}

func TestPrefixSimpleRepeat(t *testing.T) {
	result := prefix("aba")
	assert.Equal(t, []int{0, 0, 1}, result)
}
func TestPrefixSimpleRepeatSame(t *testing.T) {
	result := prefix("aaa")
	assert.Equal(t, []int{0, 1, 2}, result)
}

func TestPrefixComplexRepeat(t *testing.T) {
	result := prefix("abcababcbabca")
	assert.Equal(t, []int{0, 0, 0, 1, 2, 1, 2, 3, 0, 1, 2, 3, 4}, result)
}
