package main

import (
	"errors"
	"fmt"
)

/*
Let's build an atm machine! Woohoo
It has multiple error types --> branch logic based on error

Scenario: I have an account with balance = 100
I try to withdraw money.
Possible errors:
	Negative amount
	Zero amount
	Insufficient funds
Each error should be handled differently.
Method receiver: The method receives the address of the Account struct. modify the original struct & avoids copying the whole struct
The Withdraw function belongs to the Account type
--> account.Withdraw(100)
*/

//define the error type, root cause, repository layer talking to db
var (
	ErrInvalidAmount = errors.New("Invalid amount") //*errors.errorString
	ErrInsufficientFund = errors.New("Insufficient fund")
)

type Account struct {
	balance int
	id		int
}

//Domain Layer (Business Rules)
func (a *Account) Withdraw(amount int) error {
	
	//1.BASIC HANDLING (OLD)
	// if amount < 0 {
	// 	return fmt.Errorf("You're trying to with draw negative amount")
	// } 
	// if amount == 0 {
	// 	return fmt.Errorf("Why are you withdrawing 0 money?")
	// }
	// if amount > a.balance {
	// 	return fmt.Errorf("Dude you're broke")
	// }
	
	//2.WITH ERROR SENTINEL
	// if amount < 0 {
	// 	return ErrInvalidAmount
	// }
	// if (amount == 0) {
	// 	return ErrInvalidAmount
	// }
	// if (amount > a.balance) {
	// 	return ErrInsufficientFund
	// }

	//3.WITH ERROR WRAPPING
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.balance {
		return ErrInsufficientFund
	}

	a.balance -= amount
	return nil
}

// service/business layer: wraps with context (traceability)
//Coordinates between: Repository & Domain logic
func (a *Account)WithdrawService(amount int) error {
	err:= a.Withdraw(amount)
	if (err != nil) {
		return fmt.Errorf("withdrawal failed for account %d: %w ", a.id, err)
	}
	return nil
}

//presentation layer
func (a *Account)WithdrawServiceCustomer(amount int) {
	err:= a.WithdrawService(777)
	if (err == nil) {
		fmt.Println("Withdrawal successful. New balance:", a.balance)
		return
	}
	
	//maps error to customer-facing message
	switch {
	case errors.Is(err, ErrInvalidAmount):
		fmt.Println("For customer: Please enter a positive amount.")

	case errors.Is(err, ErrInsufficientFund):
		fmt.Println("For customer: Not enough balance.")
	
	default:	
		fmt.Println("For customer: Something went wrong.")

	}
	fmt.Println("DEBUG:", err)
}

func main() {

	fmt.Println("WELCOME TO THE ATM! HOW MUCH YOU WANT TO WITHDRAW?")

	account := Account {balance: 100, id: 999}
	amount := 9999

	account.WithdrawServiceCustomer(amount)

	fmt.Println("End of program. Bai!!!")
}