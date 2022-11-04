package ocp

import "fmt"

type BankQuery struct {
}

func (self *BankQuery) Exec() interface{} {
	fmt.Println("Exec BanckQuery")
	return nil
}
