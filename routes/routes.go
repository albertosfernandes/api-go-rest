package routes

import (
	"log"
	"net/http"

	"github.com/albertosfernandes/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/vms", controllers.GetVms).Methods("Get")
	r.HandleFunc("/api/vm/{id}", controllers.GetVm).Methods("Get")
	r.HandleFunc("/api/vm", controllers.Postvm).Methods("Post")
	r.HandleFunc("/api/vm/{id}", controllers.Deletevm).Methods("Delete")
	r.HandleFunc("/api/vm/{id}", controllers.Putvm).Methods("Put")
	r.HandleFunc("/api/networks", controllers.GetNetworks).Methods("Get")
	r.HandleFunc("/api/network/{id}", controllers.GetNetwork).Methods("Get")
	r.HandleFunc("/api/network", controllers.PostNetwork).Methods("Post")
	r.HandleFunc("/api/vm/{id}", controllers.DeleteNetwork).Methods("Delete")
	r.HandleFunc("/api/vm/{id}", controllers.PutNetwork).Methods("Put")
	r.HandleFunc("/api/storages", controllers.GetStorages).Methods("Get")
	r.HandleFunc("/api/storage/{id}", controllers.GetStorage).Methods("Get")
	r.HandleFunc("/api/storage", controllers.PostStorage).Methods("Post")
	r.HandleFunc("/api/storage/{id}", controllers.DeleteStorage).Methods("Delete")
	r.HandleFunc("/api/storage/{id}", controllers.PutStorage).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
