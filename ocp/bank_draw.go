package ocp

import "fmt"

type BankDraw struct {
}

func (self *BankDraw) Exec() interface{} {
	fmt.Println("Exec BankDraw")
	return nil
}
