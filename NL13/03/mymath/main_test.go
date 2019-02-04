package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}

	tests := []test{
		test{[]int{1,2,3},2},
		test{[]int{1,5,10},5},
		test{[]int{10,20,30},20},
		test{[]int{100,150,200,250},175},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)

		}
	}


 }

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{4, 6, 10}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}