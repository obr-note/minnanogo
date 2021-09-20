package clac

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println("わあああいっj")
	if sum(1, 2) != 3 {
		t.Fatal("sum(1,2) should be 3, but doesn't match")
	}
}

func ExampleHello() {
	fmt.Println("Hello")
	// Output: Hellol
}
