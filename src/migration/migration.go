package migration

import (
	Entity "services.com/entity"
	"time"
    "fmt"
	"log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Migration interface {
	MigrateTables() *gorm.DB
	ConnectToDb() *gorm.DB
}

type MigrationService struct {
   
}

func (migration *MigrationService) MigrateTables() *gorm.DB {
	
	var connectionString string = "postgresql://postgres:root@localhost:5432/postgres";
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{})

    if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

    db.AutoMigrate(&Entity.Customers{});
	db.AutoMigrate(&Entity.Account{});
	db.AutoMigrate(&Entity.CustomerTransactions{});

    postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Minute * 10);
        fmt.Println("Database connection timeout has been set to 10mins")
    }
    
	return db;

}

func (migration *MigrationService) ConnectToDb() *gorm.DB {
	
	var connectionString string = "postgresql://postgres:root@localhost:5432/postgres";
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{});

	if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

    postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Minute * 10)
        fmt.Println("Database connection timeout has been set to 10mins")
    }

    return db;
}








