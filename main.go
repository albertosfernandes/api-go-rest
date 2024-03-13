package main

import (
	"log"

	"github.com/albertosfernandes/api-go-rest/database"
	//"github.com/albertosfernandes/api-go-rest/routes"
)

func main() {

	database.Dbconnect()
	println("Servidor inicializado, ouvindo em :8000")
	//routes.HandleRequest()
	if err := HandleRequest(); err != nil {
		log.Fatalln(err)
	}
}
