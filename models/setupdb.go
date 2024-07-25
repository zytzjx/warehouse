package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	mysql_server   = "10.10.5.18"
	mysql_port     = 3306
	mysql_user     = "jeffery"
	mysql_password = "20040413"
	mysql_database = "db_fd_handset"
)

func ConnectDatabase() {
	mysql_connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		mysql_user, mysql_password, mysql_server, mysql_port, mysql_database)

	db, err := gorm.Open(mysql.Open(mysql_connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB = db
}
