package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000
	error := transaction.IsValid()
	assert.Error(t, error)
	assert.Equal(t, "you dont have limit for this transaction", error.Error())
}

func TestTransactionWithAmountLesserThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	error := transaction.IsValid()
	assert.Error(t, error)
	assert.Equal(t, "the amount must be greater than 1", error.Error())
}
