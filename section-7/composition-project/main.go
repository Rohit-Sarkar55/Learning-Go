package main

import (
	"errors"
	"fmt"
)

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited $%.2f to $s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient funds in the account")
	}
	acc.Balance -= amount
	fmt.Printf("Amount of $%.2f from Account Number $s. has been withdrawn successfully, New Amount: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [%s] Owner: %s , Balance: %.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}
func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.InterestRate * sa.Balance
	fmt.Printf("Added interest rate to $%.2f to Savings Account :%s", interest, sa.AccountNumber)
}

func main() {

}
