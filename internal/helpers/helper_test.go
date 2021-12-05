package helpers

import (
	"fmt"
	"testing"
)

func TestGetInts(t *testing.T) {
	a, err := ReadInts("ints.txt")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if len(a) == 0 {
		t.Errorf("expected values, got len(a)==0")
	}
	fmt.Printf("len = %d\n", len(a))
}
