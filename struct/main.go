package main

import (
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt time.Time
	LastModified time.Time
}

type BankAcc struct {
	AccName string
	Balance float64
	AuditInfo //embeddingaudiinfo struct 
	
}
type Employee struct {

	// taking variables
	name string
	empid int
}
func main() {
	account := BankAcc{AccName: "minuda", Balance: 5000.00}
	fmt.Println(account)

	account2 := &BankAcc{AccName: "saswitha", Balance: 7500.00}
	fmt.Println(account2)

	account3 := BankAcc{
		AccName: "H.D.Minuda",
		Balance: 1000.20,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	fmt.Println(account3)

	emp := Employee{"ABC", 19078}
	pts := &emp
	fmt.Println(pts)

	// accessing the struct fields(liem employee's name)
	// using a pointer but here we are not using
	// dereferencing explicitly
	fmt.Println(pts.name)

	// same as above by explicitly using
	// dereferencing concept
	// means the result will be the same
	fmt.Println((*pts).name)

}