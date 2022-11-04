package ocp

import (
	"testing"
)

func setup() {
}

func teardown() {
}

func TestBankQuery(t *testing.T) {
	bankQuery := &BankQuery{}
	bankQuery.Exec()
}

func TestBankDraw(t *testing.T) {
	bankDraw := &BankDraw{}
	bankDraw.Exec()
}

func TestBankSave(t *testing.T) {
	bankSave := &BankSave{}
	bankSave.Exec()
}
