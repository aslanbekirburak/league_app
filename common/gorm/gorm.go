package gorm

import (
	"league_app/entities"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBGORM *gorm.DB

func GetInstance() *gorm.DB {
	return DBGORM
}

func InitGorm(host, port, dbName, user, pass string) {
	var err error
	fmt.Println("host", host,user)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
	DBGORM, err = gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println("err", err)
		// LogError(err)
	}
	//INIT DATABASES
	isExist := DBGORM.HasTable(&entities.Player{})
	if !isExist {
		DBGORM.Set("gorm:DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", "ENGINE=InnoDB").CreateTable(&entities.Player{})
	}

	isExist = DBGORM.HasTable(&entities.Team{})
	fmt.Println("isExist", !isExist)
	if !isExist {
		DBGORM.Set("gorm:DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci", "ENGINE=InnoDB").CreateTable(&entities.Team{})
	}
}
