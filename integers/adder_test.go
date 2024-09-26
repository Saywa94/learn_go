package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	assertCorrectResult(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectResult(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("got %d want %d", sum, expected)
	}

}
