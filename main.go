package main

import (
	"flag"
	"fmt"
)

func main() {
	// Initialize a new blockchain
	bc := NewBlockchain()

	// Initialize a new wallet
	w := NewWallet()

	// Use command line arguments to determine behavior
	addTransaction := flag.Bool("add", false, "Add a transaction")
	amount := flag.Int("amount", 0, "Amount of the transaction")
	receiver := flag.String("receiver", "", "Receiver of the transaction")
	flag.Parse()

	if *addTransaction {
		if *amount == 0 || *receiver == "" {
			fmt.Println("You must include an amount and receiver with the -amount and -receiver flags.")
			return
		}

		// Create and sign a transaction
		t := w.createTransaction(*receiver, *amount)

		// Add the transaction to a block
		bc.AddBlock(t)
	} else {
		// Print out the blockchain
		for _, block := range bc.blocks {
			fmt.Printf("Previous hash: %s\n", block.PrevBlockHash)
			fmt.Printf("Transaction: %s\n", block.Transaction)
			fmt.Printf("Hash: %s\n", block.Hash)
			fmt.Println()
		}
	}
}
