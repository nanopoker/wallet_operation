package main

import (
	"errors"
	"testing"
)

func TestDeposit(t *testing.T) {
	testDetails := PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
	}
	confirmation, err := Deposit(testDetails)
	if err != nil {
		t.Errorf("got error %v", err)
	}
}
