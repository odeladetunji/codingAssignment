package main

import (
	"github.com/gin-gonic/gin"
	Repository "services/repository"
	"time"
	random "math/rand"
)

type TransactionService struct {

}

var transactionService TransactionService;
var transactionRepo CustomerTransactionsRepository.CustomerTransactionsRepo = 
&CustomerTransactionsRepository.CustomerTransactionsRepo{};

func (transa *TransactionService) CreateTransaction(c *gin.Context) error {

	type Payload struct {
		CustomerId int `json:"customerId"`
		InitialCredit int `json:"initialCredit"`
	}

	var payload Payload;
	if err := c.BindJSON(&payload); err != nil {
		return errors.New(err.Error());
	}

	var trans Entity.CustomerTransactions;

	trans.Amount = payload.InitialCredit;
	trans.Type = "CREDIT";
    trans.CustomerId = payload.CustomerId;
	trans.CreatedDate = fmt.Sprint(time.Now());
	trans.CreatedBy = "Admin";
	trans.LastActivityBy = "Admin";
	trans.LastActivityDate = fmt.Sprint(time.Now());

    errTrn := transactionRepo.CreateCustomerTransaction(trans);
	if errTrn != nil {
		return errors.New(errTrn.Error());
	}

	return nil;
    
}

func (transa *TransactionService) GetAllCustomerTransactions(c *gin.Context) ([]Entity.CustomerTransactions, error) {
	
	customerId, err := strconv.Atoi(c.Query("customerId"));
	if err != nil {
		return 0, errors.New(err.Error());
	}

	transactionList, errT := transactionRepo.GetAllCustomerTransactions(customerId);
	if errT != nil {
		return []Entity.CustomerTransactions, errors.New(errT.Error());
	}

	return transactionList;

}








