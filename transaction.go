package main

import "strconv"

// Transaction represents a transaction from one party to another
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    int
	Signature string
}

// NewTransaction creates a new transaction
func NewTransaction(sender string, receiver string, amount int) *Transaction {
	transaction := Transaction{sender, receiver, amount, ""}
	transaction.SignTransaction()
	return &transaction
}

// SignTransaction 'signs' a transaction by creating a simple 'signature'.
// Note that this doesn't provide any real security or validation.
func (t *Transaction) SignTransaction() {
	t.Signature = t.Sender + t.Receiver + strconv.Itoa(t.Amount)
}

// IsTransactionValid checks the validity of a transaction
// In a real blockchain, this would involve much more than just checking the 'signature'
func (t *Transaction) IsTransactionValid() bool {
	return t.Signature == t.Sender+t.Receiver+strconv.Itoa(t.Amount)
}
