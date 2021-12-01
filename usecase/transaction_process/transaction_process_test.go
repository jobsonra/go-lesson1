package transaction_process

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_entity "github.com/lesson/entity/mock"
	"github.com/stretchr/testify/assert"
)

func TestTransactionProcessWhenItsValid(t *testing.T) {
	input := TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_entity.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	usecase := NewTransactionProcess(repositoryMock)
	output, error := usecase.Execute(input)
	assert.Nil(t, error)
	assert.Equal(t, expectedOutput, output)
}
