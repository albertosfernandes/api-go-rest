package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/albertosfernandes/api-go-rest/database"
	"github.com/albertosfernandes/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetVms(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(models.Vms) # retornar moc
	var allvms []models.VM
	database.DB.Find(&allvms)
	json.NewEncoder(w).Encode(allvms)

}

func GetVm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var vm models.VM
	database.DB.First(&vm, id)
	json.NewEncoder(w).Encode(vm)
}

func Postvm(w http.ResponseWriter, r *http.Request) {
	var vm models.VM
	json.NewDecoder(r.Body).Decode(&vm)
	database.DB.Create(&vm)
	json.NewEncoder(w).Encode(vm)
}

func Deletevm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var vm models.VM
	database.DB.Delete(&vm, id)
	json.NewEncoder(w).Encode(vm)
}

func Putvm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var vm models.VM
	database.DB.First(&vm, id)
	json.NewDecoder(r.Body).Decode(&vm)
	database.DB.Save(&vm)
	json.NewEncoder(w).Encode(vm)
}

func GetNetworks(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(models.Vms) # retornar moc
	var allnetwork []models.Network
	database.DB.Find(&allnetwork)
	json.NewEncoder(w).Encode(allnetwork)

}

func GetNetwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var net models.Network
	database.DB.First(&net, id)
	json.NewEncoder(w).Encode(net)
}

func PostNetwork(w http.ResponseWriter, r *http.Request) {
	var net models.Network
	json.NewDecoder(r.Body).Decode(&net)
	database.DB.Create(&net)
	json.NewEncoder(w).Encode(net)
}

func DeleteNetwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var net models.Network
	database.DB.Delete(&net, id)
	json.NewEncoder(w).Encode(net)
}

func PutNetwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var net models.Network
	database.DB.First(&net, id)
	json.NewDecoder(r.Body).Decode(&net)
	database.DB.Save(&net)
	json.NewEncoder(w).Encode(net)
}

func GetStorages(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(models.Vms) # retornar moc
	var allstorage []models.Storage
	database.DB.Find(&allstorage)
	json.NewEncoder(w).Encode(allstorage)

}

func GetStorage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var storage models.Storage
	database.DB.First(&storage, id)
	json.NewEncoder(w).Encode(storage)
}

func PostStorage(w http.ResponseWriter, r *http.Request) {
	var storage models.Storage
	json.NewDecoder(r.Body).Decode(&storage)
	database.DB.Create(&storage)
	json.NewEncoder(w).Encode(storage)
}

func DeleteStorage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var storage models.Storage
	database.DB.Delete(&storage, id)
	json.NewEncoder(w).Encode(storage)
}

func PutStorage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var storage models.Storage
	database.DB.First(&storage, id)
	json.NewDecoder(r.Body).Decode(&storage)
	database.DB.Save(&storage)
	json.NewEncoder(w).Encode(storage)
}
