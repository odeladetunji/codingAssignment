package main

import (
	"github.com/gin-gonic/gin"
	Repository "services.com/repository"
	"time"
	random "math/rand"
	Dto "services.com/dto"
	Entity "services.com/entity"
	"errors"
	// "fmt"
)

type AccountService struct {

}

var accountService AccountService;
var customerRepo Repository.CustomerRepository = &Repository.CustomerRepo{};
var accountRepo Repository.AccountRepository = &Repository.AccountRepo{};

func (acct *AccountService) GetCustomers() ([]Entity.Customers, error) {
    customerList, err := customerRepo.GetAllCustomers();
    if err != nil {
		return []Entity.Customers{}, errors.New(err.Error());
	}

	return customerList, nil;
}

func (acct *AccountService) CreateAccount(c *gin.Context) error {

	type Payload struct {
		CustomerId int `json:"customerId"`
		InitialCredit int `json:"initialCredit"`
	}

	var payload Payload;
	if err := c.BindJSON(&payload); err != nil {
		return errors.New(err.Error());
	}

	var account Entity.Account;
	account.CustomerId = payload.CustomerId;
	account.Type = "CURRENT ACCOUNT";
	account.CreatedDate = time.Now().String();
	account.CreatedBy = "Admin";
	account.LastActivityBy = "Admin";
	account.LastActivityDate = time.Now().String();

	if payload.InitialCredit != 0 {
		account.Balance = payload.InitialCredit;
		errAct := accountRepo.CreateAccount(account);
		if errAct != nil {
			return errors.New(errAct.Error());
		}

		if errAct == nil {

			var transPayload Dto.TransPayload;
			transPayload.CustomerId = payload.CustomerId;
			transPayload.InitialCredit = payload.InitialCredit;

			errMic := microservice.CreateTransaction(transPayload);
			if errMic != nil {
				return errors.New("Error Creating Transactions")
			}
		}
	}

	if payload.InitialCredit == 0 {
		errAct := accountRepo.CreateAccount(account);
		if errAct != nil {
			return errors.New(errAct.Error());
		}
	}

	return nil;
}

func (acct *AccountService) CreateCustomers(){
	
	var customer Entity.Customers;
	customer.Id = 1;
	customer.Name = "Sam";
	customer.Surname = "Grey";
	customer.CustomerId = int(random.Int31());
	customer.CreatedDate = time.Now().String();
	customer.CreatedBy = "Admin";
	customer.LastActivityDate = time.Now().String();
	customer.LastActivityBy = "Admin";

    customerRepo.CreateCustomer(customer);

	customer.Id = 2;
	customer.Name = "Victor";
	customer.Surname = "Bay";
	customer.CustomerId = int(random.Int31());
	customer.CreatedDate = time.Now().String();
	customer.CreatedBy = "Admin";
	customer.LastActivityDate = time.Now().String();
	customer.LastActivityBy = "Admin";

	customerRepo.CreateCustomer(customer);

	customer.Id = 3;
	customer.Name = "Ford";
	customer.Surname = "Henry";
	customer.CustomerId = int(random.Int31());
	customer.CreatedDate = time.Now().String();
	customer.CreatedBy = "Admin";
	customer.LastActivityDate = time.Now().String();
	customer.LastActivityBy = "Admin";

	customerRepo.CreateCustomer(customer);

}















