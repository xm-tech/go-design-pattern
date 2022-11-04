package main

import "github.com/xm-tech/go-design-pattern/ocp"

func main() {
	bq := &ocp.BankQuery{}
	bq.Exec()

	bs := &ocp.BankSave{}
	bs.Exec()

	bd := &ocp.BankDraw{}
	bd.Exec()
}
