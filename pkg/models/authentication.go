package models

import (
	"time"

	"github.com/AstroStreakNet/database/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Authentication struct {
	AuthID       int        `gorm:"primaryKey;AUTO_INCREMENT" json:"authID"`
	UserID       int        `json:"userID"`
	SessionToken string     `json:"sessionToken"`
	LoginTime    time.Time  `json:"loginTime"`
	LogoutTime   *time.Time `json:"logoutTime"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Authentication{})
}

func CreateAuthentication(auth *Authentication) *Authentication {
	db.NewRecord(auth)
	db.Create(&auth)
	return auth
}

func GetAllAuthentications() []Authentication {
	var authentications []Authentication
	db.Find(&authentications)
	return authentications
}

func GetAuthenticationByID(ID int64) (*Authentication, error) {
	var auth Authentication
	if err := db.First(&auth, ID).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func DeleteAuthentication(ID int64) error {
	if err := db.Delete(&Authentication{}, ID).Error; err != nil {
		return err
	}
	return nil
}
