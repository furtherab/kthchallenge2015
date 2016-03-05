package blackfriday

import (
	"testing"
	"fmt"
)

type TestFindLargestUniqueNumberTest struct {
	Numbers []int
	ExpectNone bool
	ExpectNumber int
}

func TestFindLargestUniqueNumber(t *testing.T) {

	tests := []TestFindLargestUniqueNumberTest{
		{[]int{1,1,1,5,3,4,6,6}, false, 4},
		{[]int{4,4,4}, true, 0},
	}

	for _, test := range tests {
		none, number := FindLargestUniqueNumber(test.Numbers)
		fmt.Printf("%v %v", none, number)
		if test.ExpectNone != none {
			t.Errorf("Expected none %v but got %v\n", test.ExpectNone, none)
		} else if(test.ExpectNumber != number) {
			t.Errorf("Expected number %d but got %d\n", test.ExpectNumber, number)
		}
	}

}
