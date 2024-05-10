package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AstroStreakNet/database/pkg/models"
	"github.com/AstroStreakNet/database/pkg/utils"
	"github.com/gorilla/mux"
)

var NewAuthentication models.Authentication

func GetAllAuthentications(w http.ResponseWriter, r *http.Request) {
	authentications := models.GetAllAuthentications()
	res, _ := json.Marshal(authentications)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAuthenticationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authID := vars["authID"]
	ID, err := strconv.ParseInt(authID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authDetails, err := models.GetAuthenticationByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(authDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAuthentication(w http.ResponseWriter, r *http.Request) {
	newAuth := &models.Authentication{}
	utils.ParseBody(r, newAuth)
	newAuth.LoginTime = time.Now()
	auth := models.CreateAuthentication(newAuth)
	res, _ := json.Marshal(auth)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAuthentication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authID := vars["authID"]
	ID, err := strconv.ParseInt(authID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	err = models.DeleteAuthentication(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
