package main

import (
	"github.com/gin-gonic/gin"
	Entity "services.com/entity"
	Repository "services.com/repository"
	"time"
	"errors"
	"strconv"
)

type TransactionService struct {

}

var transactionService TransactionService;
var transactionRepo Repository.CustomerTransactionsRepository = &Repository.CustomerTransactionsRepo{};

func (transa *TransactionService) CreateTransaction(c *gin.Context) error {

	type Payload struct {
		CustomerId int `json:"customerId"`
		InitialCredit float64 `json:"initialCredit"`
	}

	var payload Payload;
	if err := c.BindJSON(&payload); err != nil {
		return errors.New(err.Error());
	}

	var trans Entity.CustomerTransactions;

	trans.Amount = payload.InitialCredit;
	trans.Type = "CREDIT";
    trans.CustomerId = payload.CustomerId;
	trans.CreatedDate = time.Now();
	trans.CreatedBy = "Admin";
	trans.LastActivityBy = "Admin";
	trans.LastActivityDate = time.Now();

    errTrn := transactionRepo.CreateCustomerTransaction(trans);
	if errTrn != nil {
		return errors.New(errTrn.Error());
	}

	return nil;
    
}

func (transa *TransactionService) GetAllCustomerTransactions(c *gin.Context) ([]Entity.CustomerTransactions, error) {

	customerId, err := strconv.Atoi(c.Query("customerId"));
	if err != nil {
		return []Entity.CustomerTransactions{}, errors.New(err.Error());
	}

	transactionList, errT := transactionRepo.GetAllCustomerTransactions(customerId);
	if errT != nil {
		return []Entity.CustomerTransactions{}, errors.New(errT.Error());
	}

	return transactionList, nil;

}








