package structs

import (
	"bufio"

	account "auto-pick.com/structs/account"
)

type Data struct {
	Account   account.Account
	Scanner   *bufio.Scanner
	OtherData any
}
