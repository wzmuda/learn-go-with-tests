package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

// Examples are compiled and executed as part of test suite
// % go test -v
// === RUN   TestAdder
// --- PASS: TestAdder (0.00s)
// === RUN   ExampleAdd
// --- PASS: ExampleAdd (0.00s)
// PASS
// ok      integers        0.187s
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// THE LAST LINE IS CRUCIAL - it's a special syntax comment
	// without it the example will be silently skipped from execution
	// Output: 6
}
