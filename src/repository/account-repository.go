package repository

import (
	"gorm.io/gorm"
	Entity "services.com/entity"
	Migration "services.com/migration"
	"errors"
	// "time"
)

type AccountRepository interface {
	CreateAccount(account Entity.Account, tx *gorm.DB) (error) 
	DBconnection() *gorm.DB
	GetAccountByCustomerId(customerId int) (Entity.Account, error)
}

type AccountRepo struct {

}

func (acctR *AccountRepo) DBconnection() (*gorm.DB) {
	var migration Migration.Migration = &Migration.MigrationService{};
	return migration.ConnectToDb();
}

func (acctR *AccountRepo) CreateAccount(account Entity.Account, tx *gorm.DB) (error) {
	dbError := tx.Create(&account).Error;
	if dbError != nil {
		return errors.New(dbError.Error());
	}

	return nil;
}

func (acctR *AccountRepo) GetAccountByCustomerId(customerId int) (Entity.Account, error) {
	var database *gorm.DB = dbsi.ConnectToDb();
	var account Entity.Account;
	dbError := database.Model(&Entity.Account{}).Where(&Entity.Account{CustomerId: customerId}).Find(&account).Error;
	if dbError != nil {
		return Entity.Account{}, errors.New(dbError.Error());
	}

	return account, nil;
}












