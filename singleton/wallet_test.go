package singleton

import (
	"testing"
)

func TestGetWallet(t *testing.T) {
	wallet := GetWallet()
	if wallet == nil {
		t.Errorf("Wallet is null.\n")
	}
}

func TestDepositToWallet(t *testing.T) {
	var amount float32 = 1.0
	wallet := GetWallet()
	wallet.Deposit(amount)
	actual := wallet.Check()

	if actual != amount {
		t.Errorf("Wallet has %f and the test added %f.\n", actual, amount)
	}
}

func TestWithdrawFromWallet(t *testing.T) {
	wallet := GetWallet()
	wallet.Deposit(2.0)
	var money float32 = 1.0
	amount := wallet.Withdraw(money)

	if money != amount {
		t.Errorf("Wallet returns %f when the amount should be %f.\n", amount, money)
	}
}

func TestTwoWallet(t *testing.T) {
	var amountW1 float32 = 2.0
	wallet1 := GetWallet()
	wallet1.Deposit(amountW1)

	var amountW2 float32 = 4.0
	wallet2 := GetWallet()
	wallet2.Deposit(amountW2)

	if wallet1.Check() != wallet2.Check() {
		t.Errorf("The amounts are different between the two wallets.\n")
	}
}
