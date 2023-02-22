module services.com/repository

go 1.18

require (
	gorm.io/gorm v1.24.5
	services.com/entity v0.0.0-00010101000000-000000000000
	services.com/migration v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gorm.io/driver/postgres v1.4.8 // indirect
)

replace services.com/dto => ../dto

replace services.com/entity => ../entity

replace services.com/migration => ../migration
