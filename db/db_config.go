package db

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/golang-twitter-clone-api/domain"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {

	getENV()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	migrateDDL(db)

	return db
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(domain.User{})
	db.AutoMigrate(domain.Tweet{})
	db.AutoMigrate(domain.TweetConversation{})
}

func getENV() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading .env file")
	}
	logrus.Info("Twitter API succees db connection")
}
