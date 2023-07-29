package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewWallet creates and returns Wallet
func NewWallet() *Wallet {
	privateKey, _ := getPrivateKey()
	publicKey := &privateKey.PublicKey

	return &Wallet{privateKey, publicKey}
}

// createTransaction creates a new transaction from this wallet
func (w *Wallet) createTransaction(receiver string, amount int) *Transaction {
	t := NewTransaction(w.getPublicKeyAsString(), receiver, amount)
	t.SignTransaction()
	return t
}

// getPublicKeyAsString returns public key as string
func (w *Wallet) getPublicKeyAsString() string {
	pubASN1, _ := x509.MarshalPKIXPublicKey(w.PublicKey)
	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return string(pubBytes)
}

// getPrivateKey creates a new private key
func getPrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		os.Exit(1)
	}

	return privateKey, nil
}
