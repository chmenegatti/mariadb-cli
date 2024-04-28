package database

import (
	"fmt"
	"log"
	"nemesis-cli/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	if dbInstance == nil {

		var (
			dbname = config.GetEnvKeys("DB_NAME")
			dbuser = config.GetEnvKeys("DB_USER")
			dbpass = config.GetEnvKeys("DB_PASS")
			dbhost = config.GetEnvKeys("DB_HOST")
			dbport = config.GetEnvKeys("DB_PORT")
		)

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			dbuser, dbpass, dbhost, dbport, dbname,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Erro ao conectar ao banco de dados:", err)
		}

		dbInstance = db
	}

	return dbInstance
}
