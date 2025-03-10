package db

import (
	"fmt"
	"log"
	
	"os"
	"shorty/models"

	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Db *gorm.DB

func ConnectDb(){
	
var err error
	dsn:=os.Getenv("DB_URL")
	if dsn==""{
		log.Fatal("Emty Db url")
	}
	Db,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("Problem COnnecting to db")
	}

	fmt.Println("succefully connected to db")

	Db.AutoMigrate(&models.LongShort{})
	fmt.Println("Table intialized")




}