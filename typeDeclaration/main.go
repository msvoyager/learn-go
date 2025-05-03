package main 

import (
	"fmt" 
)

// if Accnum should have charactoristics like certain lenth and etc it is cant be done by struct so we use a defferent thing

type AccountNumber string //create a new type this type name is become AccountNumber not string

//but we can make type that use the default data types
type balance = float64 //this is not going to be type as balance it is float64 

type newBankAcc struct {
	AccNum	AccountNumber 
	
}

func (a AccountNumber) IsValid() bool {
	return len(string(a)) == 10
}
func main() {
	account := newBankAcc{"1023856496"}

	var balance1 balance

	fmt.Printf("%T \n", balance1)

	fmt.Println(account.AccNum.IsValid()) //this pass the accnum value to the isValid function(method)

}