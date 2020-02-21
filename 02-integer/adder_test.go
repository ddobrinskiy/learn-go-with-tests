package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	}

	t.Run("two plus two", func(t *testing.T) {

		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)

	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output: 6
}
