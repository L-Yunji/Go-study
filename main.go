package main

import (
	"fmt"

	accounts "github.com/yunji/Go-study/banking"
)

func main() {
	account := accounts.NewAccount("yunji")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
