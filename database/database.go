package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pingkuan/go-fiber-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var (
	User string = os.Getenv("MYSQL_USER")
	Password string = os.Getenv("MYSQL_PASSWORD")
	Addr string = os.Getenv("MYSQL_ADDR")
	Port string = os.Getenv("MYSQL_PORT")
	Database string = os.Getenv("MYSQL_DATABASE")
)

func ConnectDb(){

	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",User,Password,Addr,Port,Database)
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	
	db.AutoMigrate(&models.User{})

	log.Println("Database migrated")

	Db=db

}