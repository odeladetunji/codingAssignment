package repository

import (
	"gorm.io/gorm"
	Entity "services.com/entity"
	Migration "services.com/migration"
	"errors"
	"time"
)

type AccountRepository interface {
	CreateAccount(Entity.Account) (error) 
	DBconnection() (*gorm.DB)
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
		return Entity.Account{}, errors.New(dbError.Error());
	}

	return account, nil;
}














