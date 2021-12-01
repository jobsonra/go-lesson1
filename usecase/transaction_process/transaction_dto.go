package transaction_process

//DTO data transfer object

type TransactionDtoInput struct {
	ID        string
	AccountID string
	Amount    float64
}

type TransactionDtoOutput struct {
	ID           string
	Status       string
	ErrorMessage string
}
