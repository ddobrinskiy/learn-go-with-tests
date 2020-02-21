package iteration

import (
    "fmt"
    "testing"
)

func TestRepeat(t *testing.T) {

    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("default case", func(t *testing.T) {

        got := Repeat("a", 0)
        want := ""
        assertCorrectMessage(t, got, want)

    })
    t.Run("specify num of repetitions", func(t *testing.T) {

        got := Repeat("a", 3)
        want := "aaa"
        assertCorrectMessage(t, got, want)

    })
}
func BenchmarkRepeat(b *testing.B) {
    for i:=0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    res := Repeat("a", 3)
    fmt.Println(res)
    // output: aaa
}

func ExampleRepeatZero() {
    res := Repeat("a", 0)
    fmt.Println(res)
    // output:
}


func ExampleBanana() {
    res := "ba" + Repeat("na", 2)
    fmt.Println(res)
    // output: banana
}
