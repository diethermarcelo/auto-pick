package structs

import (
	structs "auto-pick.com/structs/bank"
)

type Account struct {
	FirstName   string
	LastName    string
	BankAccount structs.Bank
}
