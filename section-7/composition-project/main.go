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
	fmt.Printf("Deposited $%.2f to $%s. New Balance: $%.2f\n", amount, acc.AccountNumber, acc.Balance)
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
	err := sa.Deposit(interest)
	if err != nil {
		fmt.Printf("Error depositing $%.2f to Savings Account :%s . %v\n", interest, sa.AccountNumber, err)
	}

}

type OverdraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oa *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if (oa.Balance + oa.OverdraftLimit) < amount {
		return fmt.Errorf("withdrawal of $%.2f exceeds overdraft limit of for %s. Available including overdraft %.2f", amount, oa.AccountNumber, oa.OverdraftLimit+oa.Balance)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew $%.2f from overdraft account %s, New Balance: $%.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Println("...Bank Account System...")

	savAcc := SavingsAccount{
		AccountNumber: "1234",
		Balance:       12000,
		OwnerName:     "John Doe",
		InterestRate:  0.02,
	}
	fmt.Println("\n--- Savings Account Operations ---")
	fmt.Println(savAcc.Account.String())

	err := savAcc.Deposit(1000)
	if err != nil {
		fmt.Printf("error depositing $%.2f to savings account ")
	}

	fmt.Println(savAcc.Account.String())

	savAcc.AddInterest()
	err = savAcc.Withdraw(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Savings Details:", savAcc.Account.String())

	ovdAcc := OverdraftAccount{
		AccountNumber:  "OVD002",
		Balance:        100.00,
		OwnerName:      "Bob Spender",
		OverdraftLimit: 200.00,
	}

	fmt.Println("\n--- Overdraft Account Operations ---")
	fmt.Println(ovdAcc.Account.String())

	err = ovdAcc.Deposit(50.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(200.00)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = ovdAcc.Withdraw(100.00)
	if err != nil {
		fmt.Println("Error (expected for overdraft limit):", err)
	}

	fmt.Println("Final Overdraft Details:", ovdAcc.Account.String())

}
