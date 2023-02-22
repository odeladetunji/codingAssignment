package main

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	Repository "services.com/repository"
	"time"
	random "math/rand"
	Dto "services.com/dto"
	Entity "services.com/entity"
	"errors"
	"fmt"
	"strings"
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
		InitialCredit float64 `json:"initialCredit"`
	}

	var payload Payload;
	if err := c.BindJSON(&payload); err != nil {
		return errors.New(err.Error());
	}

	if payload.CustomerId == 0 {
		return errors.New("CustomerId is required")
	}

	customer, errCusI := customerRepo.GetCustomerByCustomerId(payload.CustomerId);
	if errCusI != nil {
		return errors.New(errCusI.Error());
	}

	if customer.Id == 0 {
		return errors.New(strings.Join([]string{"This customerId ", fmt.Sprint(payload.CustomerId), " does not exits"}, ""));
	}

	acctn, errActn := accountRepo.GetAccountByCustomerId(payload.CustomerId);
	if errActn != nil {
		return errors.New(errActn.Error())
	}

	if acctn.Id != 0 {
		return errors.New("Customer already created a Current Account")
	}

	var account Entity.Account;
	account.CustomerId = payload.CustomerId;
	account.Name = customer.Name;
	account.Surname = customer.Surname;
	account.Type = "CURRENT ACCOUNT";
	account.CreatedDate = time.Now().String();
	account.CreatedBy = "Admin";
	account.LastActivityBy = "Admin";
	account.LastActivityDate = time.Now().String();

	dbConnection := accountRepo.DBconnection();
	var tx *gorm.DB = dbConnection.Begin();

	if payload.InitialCredit != 0 {
		account.Balance = payload.InitialCredit;
		errAct := accountRepo.CreateAccount(account, tx);
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
		errAct := accountRepo.CreateAccount(account, tx);
		if errAct != nil {
			tx.Rollback();
			return errors.New(errAct.Error());
		}

		tx.Commit();
	}

	return nil;
}

func (acct *AccountService) CreateCustomers(){
	
	var customerList []Entity.Customers;
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

	err := customerRepo.CreateCustomerList(customerList);
	if err != nil {
		panic(err.Error());
	}

}















