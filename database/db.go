package database

import (
	"github.com/albertosfernandes/api-go-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Conectar ao banco de dados MySQL

var (
	DB  *gorm.DB
	err error
)

func Dbconnect() {
	string := "*:*@123@tcp(192.168.15.111:3306)/asf?charset=utf8mb4&parseTime=True"

	DB, err = gorm.Open(mysql.Open(string), &gorm.Config{})
	if err != nil {
		panic("Falha na conexão com o banco de dados: " + err.Error())
	}

	// Migrar as tabelas
	DB.AutoMigrate(&models.Network{}, &models.Storage{}, &models.VM{})
}
