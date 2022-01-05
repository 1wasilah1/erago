package database

import (
  "log"
	"erago/entity"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection success")
	return nil
}
func Migrate(table * entity.Product) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
