package main

import (
	"fmt"
)

type BankAcc struct {
	AccName string
	Balance float64
	
}
func main() {
	account := BankAcc{AccName: "minuda", Balance: 5000.00}
	fmt.Println(account)

	account2 := &BankAcc{AccName: "saswitha", Balance: 7500.00}
	fmt.Println(account2)
	fmt.Println(&account2)
	account3 := &account2
	fmt.Println(&account3)
	fmt.Println(*account3)
	*account3 = &BankAcc{AccName: "saswitha", Balance: 800.00}
	fmt.Println(account2)

}