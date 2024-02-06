package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cliente struct {
	gorm.Model
	Limit int
	Saldo int
}

type Transacao struct {
	gorm.Model
	ClienteID uint
	Descricao string
	Tipo      string 
}

func CreateDb() *gorm.DB {
	dsn := "rinha:1234@tcp(localhost:3306)/rinha_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&Cliente{}, &Transacao{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to auto migrate database")
	}


	clientes := []Cliente{
		{Limit: 100000, Saldo: 0},
		{Limit: 80000, Saldo: 0},
		{Limit: 1000000, Saldo: 0},
		{Limit: 10000000, Saldo: 0},
		{Limit: 500000, Saldo: 0},
	}

	for _, cliente := range clientes {
		result := db.Create(&cliente)
		if result.Error != nil {
			panic("Failed to create cliente")
		}
	}

	return db
}
