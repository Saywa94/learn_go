package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("Repeat 5 times c", func(t *testing.T) {
		repeated := Repeat("c", 5)
		expected := "ccccc"

		assertCorrectRepeat(t, repeated, expected)

	})
	t.Run("Repeat 12 times U", func(t *testing.T) {
		repeated := Repeat("U", 12)
		expected := "UUUUUUUUUUUU"

		assertCorrectRepeat(t, repeated, expected)

	})
}

func assertCorrectRepeat(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
