package proteins

import (
	"testing"
)

type FindLeastInsertionsTest struct {
	N int
	String string
	Expected int
}

func TestFindLeastInsertions(t *testing.T) {

	tests := []FindLeastInsertionsTest{
		{2, "TGATGC", 1},
		{4, "ATCATGCATGATGTG", 3},
	}

	for _, test := range tests {
		result := FindLeastInsertions(test.N, test.String)
		if result != test.Expected {
			t.Errorf("Expected %d as result but got %d for %d \"%v\"", test.Expected, result, test.N, test.String)
		}
	}

}
