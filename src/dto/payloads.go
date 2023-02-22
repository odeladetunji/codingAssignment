package dto;

import Entity "services.com/entity"
type TransPayload struct {
	CustomerId int `json:"customerId"`
	InitialCredit int `json:"initialCredit"`
}


type CustomerTransactionDetails struct {
	Customer Entity.Customers `json:"customer"`
	Transactions []Entity.CustomerTransactions `json:"transactions"`
}





