package db

import (
	"fmt"
	"log"
	
	//"os"
	"shorty/models"

	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var Db *gorm.DB

func ConnectDb(){
	
var err error
	// dsn:=os.Getenv("DB_URL")
	// if dsn==""{
	// 	log.Fatal("Emty Db url")
	// }

dsn:="postgresql://neondb_owner:npg_3yAVSZ7CLefq@ep-long-bonus-a8t1sjn8-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"
	Db,err=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("Problem COnnecting to db")
	}

	fmt.Println("succefully connected to db")

	Db.AutoMigrate(&models.LongShort{})
	fmt.Println("Table intialized")




}