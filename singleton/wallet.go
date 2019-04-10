package singleton

type wallet struct {
	amount float32
}

var instance *wallet

// GetWallet returns the wallet
func GetWallet() *wallet {
	if instance == nil {
		instance = new(wallet)
	}
	return instance
}

// Deposit adds the given amount to the wallet 
func (w *wallet) Deposit(amount float32) {
	w.amount += amount
}

// Withdraw takes "amount" from the wallet and returns it, if it is available.
func (w *wallet) Withdraw(amount float32) float32 {
	if amount < w.amount {
		w.amount -= amount
		return amount
	}
	w.amount = 0
	return w.amount
}

// Check returns the current amount of the wallet.
func (w *wallet) Check() float32 {
	return w.amount
}
