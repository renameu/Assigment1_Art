package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("Deposited %.2f Tenge. Current balance: %.2f Tenge\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.Balance {
		fmt.Println("Insufficient balance. Transaction failed.")
		return
	}
	b.Balance -= amount
	fmt.Printf("Withdrawn %.2f Tenge. Current balance: %.2f Tenge\n", amount, b.Balance)
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f Tenge\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Bank Account System")

	fmt.Print("Enter Account Number: ")
	scanner.Scan()
	accountNumber := scanner.Text()

	fmt.Print("Enter Holder Name: ")
	scanner.Scan()
	holderName := scanner.Text()

	account := &BankAccount{
		AccountNumber: accountNumber,
		HolderName:    holderName,
		Balance:       0,
	}

	for {
		fmt.Println("Choose an option: 1. Deposit 2. Withdraw 3. Get Balance 4. Exit")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid option. Please enter a number.")
			continue
		}

		switch option {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			if !scanner.Scan() {
				break
			}
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Deposit(amount)

		case 2:
			fmt.Print("Enter amount to withdraw: ")
			if !scanner.Scan() {
				break
			}
			amount, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid amount. Please enter a positive number.")
				continue
			}
			account.Withdraw(amount)

		case 3:
			account.GetBalance()

		case 4:
			fmt.Println("Exiting program. Thank you for using the Bank Account System!")
			return

		default:
			fmt.Println("Invalid option. Please choose 1, 2, 3, or 4.")
		}
	}
}
