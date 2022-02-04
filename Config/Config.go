package Config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbDriver := os.Getenv("dbDriver")
	dbName := os.Getenv("dbName")
	dbUsername := os.Getenv("dbUsername")
	dbPassword := os.Getenv("dbPassword")
	db, err := gorm.Open(dbDriver, dbUsername+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		return nil
	}
	return db
}
