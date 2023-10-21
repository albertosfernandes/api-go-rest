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
	var connectionString = fmt.Spintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Falha na conexão com o banco de dados: " + err.Error())
	}

	// Migrar as tabelas
	DB.AutoMigrate(&models.Network{}, &models.Storage{}, &models.VM{})
}
