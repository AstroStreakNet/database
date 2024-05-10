package models

import (
	"github.com/AstroStreakNet/database/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	UserID       int    `gorm:"primaryKey" json:"userID"`
	AstronomerID int    `json:"astronomerID"`
	Password     string `json:"password"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	DOB          int    `json:"DOB"`
	Gender       string `json:"gender"`
	Permissions  bool   `json:"permissions"`
	ImagePublish bool   `json:"imagePublish"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func GetUserByAstronomerID(astronomerID int) (*User, error) {
	var user User
	if err := db.Where("astronomerID = ?", astronomerID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
