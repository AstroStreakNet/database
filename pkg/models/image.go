package models

import (
	"time"

	"github.com/AstroStreakNet/database/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Image struct {
	ImageID           int       `gorm:"primaryKey;AUTO_INCREMENT" json:"imageID"`
	ImagePath         string    `json:"imagePath"`
	ObservatoryCode   string    `json:"observatoryCode"`
	TimeOfObservation time.Time `json:"timeOfObservation"`
	CreationDate      string    `json:"creationDate"`
	ExposureDuration  time.Time `json:"exposureDuration"`
	Coordinates       string    `json:"coordinates"`
	StreakType        string    `json:"streakType"`
	UserID            int       `json:"userID"`
	AllowPublic       bool      `json:"allowPublic"`
	AllowML           bool      `json:"allowML"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Image{})
}

func CreateImage(image *Image) *Image {
	db.NewRecord(image)
	db.Create(&image)
	return image
}

func GetAllImages() []Image {
	var images []Image
	db.Find(&images)
	return images
}

func GetImageByID(ID int64) (*Image, error) {
	var image Image
	if err := db.First(&image, ID).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func DeleteImage(ID int64) error {
	if err := db.Delete(&Image{}, ID).Error; err != nil {
		return err
	}
	return nil
}
