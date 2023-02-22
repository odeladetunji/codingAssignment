package repository

import (
	"gorm.io/gorm"
	Entity "services.com/entity"
	Migration "services.com/migration"
	"errors"
	// "time"
)

var dbsi Migration.Migration = &Migration.MigrationService{}

type CustomerTransactionsRepository interface {
	GetAllCustomerTransactions(customerId int) ([]Entity.CustomerTransactions, error)
	CreateCustomerTransaction(customerTransaction Entity.CustomerTransactions) (error)
	DBconnection() (*gorm.DB)
}

type CustomerTransactionsRepo struct {

}

func (trans *CustomerTransactionsRepo) DBconnection() (*gorm.DB) {
	var migration Migration.Migration = &Migration.MigrationService{};
	return migration.ConnectToDb();
}

func (trans *CustomerTransactionsRepo) CreateCustomerTransaction(customerTransaction Entity.CustomerTransactions) (error) {
	var database *gorm.DB = dbsi.ConnectToDb();
	dbError := database.Create(&customerTransaction).Error;
	if dbError != nil {
		return errors.New(dbError.Error());
	}

	return nil;
}

func (trans *CustomerTransactionsRepo) GetAllCustomerTransactions(customerId int) ([]Entity.CustomerTransactions, error) {
	var database *gorm.DB = dbsi.ConnectToDb();
	var customerTransactionList []Entity.CustomerTransactions;
	dbError := database.Model(&Entity.CustomerTransactions{}).Where(&Entity.CustomerTransactions{CustomerId: customerId}).Find(&customerTransactionList).Error;
	if dbError != nil {
		return []Entity.CustomerTransactions{}, errors.New(dbError.Error());
	}

	return customerTransactionList, nil;
}


















