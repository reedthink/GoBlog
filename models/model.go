package models

import (
	"fmt"
	"github.com/spf13/viper"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err         error
		dbType      string
		dbName      string
		user        string
		host        string
		password    string
		tablePrefix string
	)

	dbType = viper.GetString("database.TYPE")
	dbName = viper.GetString("database.NAME")
	user = viper.GetString("database.USER")
	password = viper.GetString("database.PASSWORD")
	host = viper.GetString("database.HOST")
	tablePrefix = viper.GetString("database.TABLE_PREFIX")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
