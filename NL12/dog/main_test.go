package Dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(2)
	if x != 14 {
		t.Error("Expected", 14, "Got", x)
	}

}

func ExampleYears() {
	fmt.Println(Years(2))
	//Output:
	//14
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}