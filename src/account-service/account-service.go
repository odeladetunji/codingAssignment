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

	dbConnection := accountRepo.DBconnection();
	tx := dbConnection.Begin();

	if payload.InitialCredit != 0 {
		account.Balance = payload.InitialCredit;
		errAct := accountRepo.CreateAccount(account, &tx);
		if errAct != nil {
			return errors.New(errAct.Error());
		}

		if errAct == nil {

			var transPayload Dto.TransPayload;
			transPayload.CustomerId = payload.CustomerId;
			transPayload.InitialCredit = payload.InitialCredit;

			errMic := microservice.CreateTransaction(transPayload);
			if errMic != nil {
				tx.Rollback();
				return errors.New("Error Creating Transactions")
			}

			tx.Commit();
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
	
	var customerList []Entity.CustomerList;
	var customer Entity.Customers;

	//FirstCustomer
	customer.Name = "Sam";
	customer.Surname = "Grey";
	customer.CustomerId = int(random.Int31());
	customer.CreatedDate = time.Now().String();
	customer.CreatedBy = "Admin";
	customer.LastActivityBy = "Admin";
	customer.LastActivityDate = time.Now().String();
    customerList = append(customerList, customer);

	//Second Customer
	customer.Name = "Victor";
	customer.Surname = "Bay";
	customer.CustomerId = int(random.Int31());
	customerList = append(customerList, customer);

	//Third Customer
	customer.Name = "Ford";
	customer.Surname = "Henry";
	customer.CustomerId = int(random.Int31());
	customerList = append(customerList, customer);

	customers, err := customerRepo.CreateCustomerList(customerList);
	if err != nil {
		panic(err.Error());
	}

}















