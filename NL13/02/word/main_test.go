package word

import (
	"fmt"
	"github.com/rheeger/go-programming-class/NL13/quote"
	"testing"
)

func TestCount(t *testing.T) {
	x := `Hello there friends`
	y := Count(x)
	if y != 3 {
		t.Error("Expected", 3, "Got", y)
	}
}

//func TestUseCount(t *testing.T) {
//}


func ExampleCount() {
	const x = `Hello there friends`
	xi := Count(x)
	fmt.Println(xi)
	// Output:
	// 3
}


func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

