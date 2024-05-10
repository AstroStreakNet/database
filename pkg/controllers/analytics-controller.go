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

var NewAnalytics models.Analytics

func GetAllAnalytics(w http.ResponseWriter, r *http.Request) {
	analytics := models.GetAllAnalytics()
	res, _ := json.Marshal(analytics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAnalyticsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	analyticsID := vars["analyticsID"]
	ID, err := strconv.ParseInt(analyticsID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	analyticsDetails, _ := models.GetAnalyticsByID(ID)
	res, _ := json.Marshal(analyticsDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAnalytics(w http.ResponseWriter, r *http.Request) {
	createAnalytics := &models.Analytics{}
	utils.ParseBody(r, createAnalytics)
	createAnalytics.TimeBetweenVisits = time.Now()
	a := models.CreateAnalytics(createAnalytics)
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAnalytics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	analyticsID := vars["analyticsID"]
	ID, err := strconv.ParseInt(analyticsID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	err = models.DeleteAnalytics(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
