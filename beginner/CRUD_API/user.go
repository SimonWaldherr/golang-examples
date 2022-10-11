package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


var DB *gorm.DB
var err error


//Gorm is used to convert our struct into a model struct so that we can store date in db using ORM
type User struct{

	gorm.Model
	FirstName string `json:"first_name"`
	LasName string `json:"last_name"`
	Email string `json:"email"`


}

func IntialMigration(){
	DB, err = gorm.Open("sqlite3", "user.db")
	if err != nil {
        panic("Failed to open the SQLite database.")
    }
	//defer DB.Close()

    // Create the table from our struct
    DB.AutoMigrate(&User{})
       
    
}

//arguments includes the response writer and the a pointer pointing to the actual request that came in
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//slice 
	var users []User
    DB.Find(&users)
    json.NewEncoder(w).Encode(users)
	
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params :=mux.Vars(r)
	DB.First(&user,params["id"])
    json.NewEncoder(w).Encode(user)
	
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params :=mux.Vars(r)
	DB.Delete(&user,params["id"])
    json.NewEncoder(w).Encode("The user is deleted")
	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
        return
	}
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)


	
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	params :=mux.Vars(r)
	DB.First(&user,params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
    json.NewEncoder(w).Encode(user)

	
}