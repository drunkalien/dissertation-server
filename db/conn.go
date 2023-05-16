package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connStr := "user=jamshidkhujaburikhujaev password=uhea67YglWxv dbname=neondb host=ep-silent-resonance-295764.eu-central-1.aws.neon.tech sslmode=verify-full"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(&User{}, &Product{}, &Order{})
}
