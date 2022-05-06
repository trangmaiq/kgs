package permutation

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	result := Permutation([]string{}, "", []string{"A", "B", "C"}, 2)
	expected := []string{"AA", "AB", "AC", "BA", "BB", "BC", "CA", "CB", "CC"}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
