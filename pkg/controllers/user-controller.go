package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AstroStreakNet/database/pkg/models"
	"github.com/AstroStreakNet/database/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &models.User{}
	utils.ParseBody(r, newUser)
	u := newUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userID := vars["userID"]
	ID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.DOB != 0 {
		userDetails.DOB = updateUser.DOB
	}
	if updateUser.Gender != "" {
		userDetails.Gender = updateUser.Gender
	}
	if updateUser.Permissions != userDetails.Permissions {
		userDetails.Permissions = updateUser.Permissions
	}
	if updateUser.ImagePublish != userDetails.ImagePublish {
		userDetails.ImagePublish = updateUser.ImagePublish
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByAstronomerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	astronomerID := vars["astronomerID"]
	ID, err := strconv.Atoi(astronomerID)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, err := models.GetUserByAstronomerID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
