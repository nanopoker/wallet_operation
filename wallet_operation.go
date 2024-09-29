package main

import (
	"fmt"
	"log"
)

func Deposit(data PaymentDetails) (string, error) {
	log.Printf("Depositing $%d into account %s.\n\n",
		data.Amount,
		data.TargetAccount,
	)

	bank := BankingService{"https://BankAPIAddress.com"}
	confirmation, err := bank.Deposit(data.TargetAccount, data.Amount)
	return confirmation, err
}

func Withdraw(data PaymentDetails) (string, error) {
	log.Printf("Withdrawing $%d from account %s.\n\n",
		data.Amount,
		data.SourceAccount,
	)

	bank := BankingService{"https://BankAPIAddress.com"}
	confirmation, err := bank.Withdraw(data.SourceAccount, data.Amount)
	return confirmation, err
}

func Send(data PaymentDetails) (string, error) {
	log.Printf("Sending $%d from account %s to account %s.\n\n",
		data.Amount,
		data.SourceAccount,
		data.TargetAccount,
	)

	bank := BankingService{"https://BankAPIAddress.com"}
	confirmation, err := bank.Send(data.SourceAccount, data.TargetAccount, data.Amount)
	return confirmation, err
}

func Check(data PaymentDetails) (int, string, error) {
	log.Printf("Checking the balance for account %s.\n\n",
		data.SourceAccount,
	)

	bank := BankingService{"https://BankAPIAddress.com"}
	confirmation, err := bank.Check(data.SourceAccount)
	return confirmation, err
}
