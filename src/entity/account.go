package entity

// import "time"

type Account struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Balance float64 `json:"balance"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	CustomerId int `json:"customerId"`
	CreatedDate string `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
	LastActivityBy string `json:"lastActivityBy"`
	LastActivityDate string `json:"lastActivityDate"`
}









