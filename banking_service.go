package main

import (
	"errors"
	"math/rand"
)

func (b bank) findAccount(accountNumber string) (account, error) {

	for _, v := range b.Accounts {
		if v.AccountNumber == accountNumber {
			return v, nil
		}
	}

	return account{}, errors.New("account not found")
}

func (m *InsufficientFundsError) Error() string {
	return "Insufficient Funds"
}

func (m *InvalidAccountError) Error() string {
	return "Account number supplied is invalid"
}

func (bank BankingService) Deposit(accountNumber string, amount int) (string, error) {
	var b Bank
	account, err := b.findAccount(accountNumber)
	if err != nil {
		return "", &InvalidAccountError{}
	}
	account.Balance += amount
	return generateTransactionID("D", 10), nil
}

func (Bank BankingService) Withdraw(accountNumber string, amount int) (string, error) {
	var b Bank
	account, err := b.findAccount(accountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	if amount > int(account.Balance) {
		return "", &InsufficientFundsError{}
	}
	account.Balance -= amount
	return generateTransactionID("W", 10), nil
}

func (Bank BankingService) Send(sourceAccountNumber string, targetAccountNumber string, amount int) (string, error) {
	var b Bank
	sourceAccount, err := b.findAccount(sourceAccountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	targetAccount, err := b.findAccount(targetAccountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	if amount > int(sourceAccount.Balance) {
		return "", &InsufficientFundsError{}
	}

	sourceAccount.Balance -= amount
	targetAccount.Balance += amount

	return generateTransactionID("S", 10), nil
}

func (Bank BankingService) Check(accountNumber string, amount int) (int, string, error) {
	var b Bank
	account, err := b.findAccount(accountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	return account.Balance, generateTransactionID("C", 10), nil
}

// simulating the process of generating transactionID
func generateTransactionID(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}
