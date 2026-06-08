package methods

import "testing"

func TestDeposit(t *testing.T) {
	amount := Amount{}

	amount.Deposit(10)

	got := amount.Balance
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
