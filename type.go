package main

type PaymentDetails struct {
	SourceAccount string
	TargetAccount string
	Amount        int
}

type Account struct {
	AccountNumber string
	Balance       int64
}

type Bank struct {
	Accounts []account
}

// balance is insufficient
type InsufficientFundsError struct{}

// account number is invalid
type InvalidAccountError struct{}

type BankingService struct {
	BankAPIAddress string
}
