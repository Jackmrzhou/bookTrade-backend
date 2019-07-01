package dao

import (
	"bookTrade-backend/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

var db *gorm.DB

func InitDatabase(config *conf.AppConfig) error {
	dbConf := config.DatabaseConfig
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		strconv.Itoa(dbConf.Port),
		dbConf.DatabaseName,
		dbConf.Charset,
	)
	db, err = gorm.Open("mysql", connStr)
	return err
}
