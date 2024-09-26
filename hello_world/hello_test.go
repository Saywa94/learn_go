package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Saywa", "")
		want := "Hello, Saywa!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, world!' when empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Helodie", "spanish")
		want := "Hola, Helodie!"

		assertCorrectMessage(t, got, want)

	})
	t.Run("In french", func(t *testing.T) {
		got := Hello("Jacques", "french")
		want := "Bonjour, Jacques!"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
