package repository

import (
	"gorm.io/gorm"
	Entity "services.com/entity"
	Migration "services.com/migration"
	"errors"
	"time"
)

type CustomerRepository interface {
	CreateCustomer(customer Entity.Customers)
	GetAllCustomers() ([]Entity.Customers, error)
	CreateCustomerList(customerList []Entity.Customers)
	DBconnection() (*gorm.DB)
}

type CustomerRepo struct {

}

var dbs Migration.Migration = &Migration.MigrationService{}

func (acct *CustomerRepo) DBconnection() (*gorm.DB) {
	var migration Migration.Migration = &Migration.MigrationService{};
	return migration.ConnectToDb();
}

func (acct *CustomerRepo) CreateCustomer(customer Entity.Customers) {
	var database *gorm.DB = dbs.ConnectToDb();
	dbError := database.Create(&customer).Error;
	if dbError != nil {
		return Entity.Customers{}, errors.New(dbError.Error());
	}

	return customer, nil;
}

func (acct *CustomerRepo) CreateCustomerList(customerList []Entity.Customers) {
	var database *gorm.DB = dbs.ConnectToDb();
	dbError := database.Create(&customerList).Error;
	if dbError != nil {
		return []Entity.Customers{}, errors.New(dbError.Error());
	}

	return customerList, nil;
}

func (acct *CustomerRepo) GetAllCustomers() ([]Entity.Customers, error) {
	var database *gorm.DB = dbs.ConnectToDb();
	var customerList []Entity.Customers;
	dbError := database.Find(&customerList).Error;
	if dbError != nil {
		return []Entity.Customers{}, errors.New(dbError.Error());
	}

	return customerList, nil;

}





