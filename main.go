package main

import (
	"gorm/postgresql/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	app := controllers.StartOp()
	router.HandleFunc("/health", healthStatus).Methods("GET")
	router.HandleFunc("/newuser", app.NewUser).Methods("POST")
	router.HandleFunc("/allusers", app.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", app.GetUser).Methods("GET")
	router.HandleFunc("/delete/user/{name}", app.DeleteUser).Methods("GET")

	http.ListenAndServe(":3000", router)
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("live on 3000"))
}
