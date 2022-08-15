package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *DataBase

type DataBase struct {
	*gorm.DB
}

func init() {
	DB = &DataBase{initDB()}
	fmt.Println("init!")
}

func initDB() *gorm.DB {

	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	address := viper.GetString("db.address")

	return openPostgres(address, username, password)
}

func openPostgres(address, username, password string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(address))

	if err != nil {
		log.Fatalf("database - failed to open %v", err)
		panic(err)
	}

	return db
}
