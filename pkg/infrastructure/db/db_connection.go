package db_postNrelSvc

import (
	"database/sql"
	"fmt"
	"time"

	domain_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/domain"
	config_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/pkg/infrastructure/config"
	interface_hash_postNrelSvc "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/hash_password/interface"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config_postNrelSvc.DataBase, hashUtil interface_hash_postNrelSvc.IhashPassword) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBPort)
	fmt.Println("database connection string ------", connectionString)
	sql, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	rows, err := sql.Query("SELECT 1 FROM pg_database WHERE datname = '" + config.DBName + "'")
	if err != nil {
		fmt.Println("Error checking database existence:", err)
		return nil, err
	}
	defer rows.Close()

	if rows == nil {
		fmt.Println("row is nil database not connection condtion true")
	}

	if rows.Next() {
		fmt.Println("Database" + config.DBName + " already exists.")
	} else {
		_, err = sql.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error creating database:", err)
		}
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC() // Set the timezone to UTC
		},
	})
	if dberr != nil {
		return DB, nil
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain_postNrelSvc.Post{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain_postNrelSvc.PostLikes{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain_postNrelSvc.PostMedia{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain_postNrelSvc.Relationship{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain_postNrelSvc.Comment{}); err != nil {
		return DB, err
	}

	return DB, nil
}
