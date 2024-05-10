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

var NewImage models.Image

func GetAllImages(w http.ResponseWriter, r *http.Request) {
	images := models.GetAllImages()
	res, _ := json.Marshal(images)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetImageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["imageID"]
	ID, err := strconv.ParseInt(imageID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	imageDetails, _ := models.GetImageByID(ID)
	res, _ := json.Marshal(imageDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	newImage := &models.Image{}
	utils.ParseBody(r, newImage)
	image := models.CreateImage(newImage)
	res, _ := json.Marshal(image)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["imageID"]
	ID, err := strconv.ParseInt(imageID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	image := models.DeleteImage(ID)
	res, _ := json.Marshal(image)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	var updateImage = &models.Image{}
	utils.ParseBody(r, updateImage)
	vars := mux.Vars(r)
	imageID := vars["imageID"]
	ID, err := strconv.ParseInt(imageID, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}
	imageDetails, _ := models.GetImageByID(ID)
	if updateImage.ImagePath != "" {
		imageDetails.ImagePath = updateImage.ImagePath
	}
	if updateImage.ObservatoryCode != "" {
		imageDetails.ObservatoryCode = updateImage.ObservatoryCode
	}
	if !updateImage.TimeOfObservation.IsZero() {
		imageDetails.TimeOfObservation = updateImage.TimeOfObservation
	}
	if updateImage.CreationDate != "" {
		imageDetails.CreationDate = updateImage.CreationDate
	}
	if !updateImage.ExposureDuration.IsZero() {
		imageDetails.ExposureDuration = updateImage.ExposureDuration
	}
	if updateImage.Coordinates != "" {
		imageDetails.Coordinates = updateImage.Coordinates
	}
	if updateImage.StreakType != "" {
		imageDetails.StreakType = updateImage.StreakType
	}
	if updateImage.UserID != 0 {
		imageDetails.UserID = updateImage.UserID
	}
	if updateImage.AllowPublic != imageDetails.AllowPublic {
		imageDetails.AllowPublic = updateImage.AllowPublic
	}
	if updateImage.AllowML != imageDetails.AllowML {
		imageDetails.AllowML = updateImage.AllowML
	}
	db.Save(&imageDetails)
	res, _ := json.Marshal(imageDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
