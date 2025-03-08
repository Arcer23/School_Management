package database

import(
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn:="host=localhost user=postgres password=pranish dbname=school_mang sslmode=disable port=5432 TimeZone=Asia/Kathmandu"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database Server", err)
	}
	fmt.Println("database connected successfully")

	DB=db
	
}