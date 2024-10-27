package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25)}
		want := Bitcoin(10)
		err := wallet.Withdraw(Bitcoin(15))

		assertNoError(t, err)
		assertBalance(t, &wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(25))

		assertError(t, err, ErrInsuficientFunds)
		assertBalance(t, &wallet, startingBalance)

	})

}

func assertBalance(t testing.TB, wallet *Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
