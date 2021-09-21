package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLServer struct {
	URLBD      string
	NameBD     string
	UserBD     string
	PasswordBD string
	PortBD     string
}

// GetConnectionSQLServer function
func GetConnectionPostgres(dbSQL *SQLServer) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbSQL.URLBD, dbSQL.UserBD, dbSQL.PasswordBD, dbSQL.NameBD, dbSQL.PortBD)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("sql open:", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Printf("Conexion exitosa con la base de datos!\n")
	return db, nil
}
