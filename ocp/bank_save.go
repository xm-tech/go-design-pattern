package ocp

import "fmt"

type BankSave struct {
}

func (self *BankSave) Exec() interface{} {
	fmt.Println("Exec BankSave")
	return nil
}
