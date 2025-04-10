package main

import (
	"fmt"
)
type BankAcc struct {
	AccNumber string
	Balance float64
}

//define a method 
//when a function is defined on a struct its called a method

func (b BankAcc) DisplayBalance() { //b is a value reciever
	fmt.Printf("Account number: %s, Balance : $%.2f", b.AccNumber, b.Balance)
}

//we can use pointer recievers also change the original location value instead of taking a copy and modify it

func (c *BankAcc) withdraw(amount float64) {
	if c.Balance > amount {
		c.Balance -= amount
	}
	
}
func main() {
	account := BankAcc{AccNumber: "102350", Balance: 850.222}
	account.withdraw(10.00)
	account.DisplayBalance() //when this called bAcc recieve the values of account
	//because of that bAcc argument identify as a reciever argument
	//methods can only have one reciever

}