package controllers

import (
	"encoding/json"
	"fmt"
	"gorm/postgresql/configs"
	"gorm/postgresql/model"
	"gorm/postgresql/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	NewUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type control struct {
	service services.Service
}

func StartOp() Controller {
	conn := configs.DatabaseConnection()
	conn.Table("accounts").AutoMigrate(&model.Person{})
	mod := model.CreateConntection(conn)
	service := services.NewService(mod)
	return &control{
		service: service,
	}
}

func (c *control) NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type person struct {
		Name   string `json:"name,omitempty"`
		Age    int    `json:"age,omitempty"`
		Number string `json:"number,omitempty"`
	}
	var account person
	json.NewDecoder(r.Body).Decode(&account)

	err := c.service.AddNewAccount(account.Name, account.Number, account.Age)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	log.Println("Created New User")
	result := fmt.Sprintf("Created user %s, %d, %s", account.Name, account.Age, account.Number)
	json.NewEncoder(w).Encode(result)
}

func (c *control) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := c.service.FindAllAccounts()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	log.Println("Found all the users")
	json.NewEncoder(w).Encode(users)
}

func (c *control) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := mux.Vars(r)["name"]
	account, err := c.service.FindAccount(name)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	log.Println("Found the User")
	json.NewEncoder(w).Encode(account)
}

func (c *control) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := mux.Vars(r)["name"]
	account, err := c.service.DeleteAccount(name)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	log.Println("Deleted the user")
	json.NewEncoder(w).Encode(account)
}
