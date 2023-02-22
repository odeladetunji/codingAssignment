package entity

import "time"

type CustomerTransactions struct {
	Id int32 `json:"id"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
	CustomerId int `json:"customerId"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
	LastActivityBy string `json:"lastActivityBy"`
	LastActivityDate time.Time `json:"lastActivityDate"`
}











