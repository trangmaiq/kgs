package permutation

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermutation(t *testing.T) {
	result := Permutation([]string{}, "", []string{"A", "B", "C"}, 2)
	expected := []string{"AA", "AB", "AC", "BA", "BB", "BC", "CA", "CB", "CC"}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

func TestPermutationChan(t *testing.T) {
	var (
		result   []string
		expected = []string{"AA", "AB", "AC", "BA", "BB", "BC", "CA", "CB", "CC"}
	)

	out := PermutationChan([]string{"A", "B", "C"}, 2)
	for o := range out {
		result = append(result, o)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[j] > result[i]
	})

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
