package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mohammadinasab-dev/employee-salary/configuration"
)

//SQLHandler is a type
type SQLHandler struct {
	DB *gorm.DB
}

func CreateDBConnection(config configuration.Config) (*SQLHandler, error) {
	connstring := fmt.Sprintf(config.DBUsername + ":" + config.DBPassword + "@" + config.DBAddress + "/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(config.DBDriver, connstring)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	db.AutoMigrate(&Employee{}, &Salary{})
	return &SQLHandler{
		DB: db,
	}, nil

}
