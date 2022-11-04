package ocp

type BankOp interface {
	Exec() interface{}
}
