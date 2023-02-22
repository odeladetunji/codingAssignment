package dto;

import Entity "services.com/entity"
type TransPayload struct {
	CustomerId int `json:"customerId"`
	InitialCredit float64 `json:"initialCredit"`
}

type CustomerTransactionDetails struct {
	Account Entity.Account `json:"customer"`
	Transactions []Entity.CustomerTransactions `json:"transactions"`
}





