package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zuams/simple-reminder/helpers"
)

func New() (*gorm.DB, error) {
	readConfig := helpers.ReadFileJson("config.json")

	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 readConfig["username"].(string),
		Passwd:               readConfig["password"].(string),
		Net:                  "tcp",
		Addr:                 readConfig["address"].(string),
		DBName:               readConfig["database"].(string),
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	return db, err
}
