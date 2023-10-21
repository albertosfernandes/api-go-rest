package main

import (
	"github.com/albertosfernandes/api-go-rest/database"
	"github.com/albertosfernandes/api-go-rest/models"
	"github.com/albertosfernandes/api-go-rest/routes"
)

func main() {

	models.Vms = []models.VM{
		{Id: 1, Name: "vm-teste01", Vcpu: 8, Mem: 12000,
			Network: models.Network{Hostname: "net-01", IPAddress: "192.168.0.101", Gateway: "192.168.0.1"},
			Storage: models.Storage{Name: "storage-01", Size: 100000, Type: "ssd"}},
		{Id: 2, Name: "vm-teste02", Vcpu: 8, Mem: 12000,
			Network: models.Network{Hostname: "net-02", IPAddress: "192.168.0.102", Gateway: "192.168.0.1"},
			Storage: models.Storage{Name: "storage-02", Size: 100000, Type: "ssd"}},
	}

	database.Dbconnect()

	println("Iniciando servidor...")
	routes.HandleRequest()
}
