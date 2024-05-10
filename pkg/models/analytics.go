package models

import (
	"time"

	"github.com/AstroStreakNet/database/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Analytics struct {
	AnalyticsID       int       `gorm:"primaryKey" json:"analyticsID"`
	ImageID           int       `json:"imageID"`
	PageViews         int       `json:"pageViews"`
	ImageViews        int       `json:"imageViews"`
	Clicks            int       `json:"clicks"`
	SessionData       string    `json:"sessionData"`
	ReferralSources   string    `json:"referralSources"`
	NavigationPaths   string    `json:"navigationPaths"`
	AbandonmentRate   float64   `json:"abandonmentRate"`
	RepeatVisits      bool      `json:"repeatVisits"`
	TimeBetweenVisits time.Time `json:"timeBetweenVisits"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Analytics{})
}

func CreateAnalytics(analytics *Analytics) *Analytics {
	db.NewRecord(analytics)
	db.Create(&analytics)
	return analytics
}

func GetAllAnalytics() []Analytics {
	var analytics []Analytics
	db.Find(&analytics)
	return analytics
}

func GetAnalyticsByID(ID int64) (*Analytics, error) {
	var analytics Analytics
	if err := db.First(&analytics, ID).Error; err != nil {
		return nil, err
	}
	return &analytics, nil
}

func DeleteAnalytics(ID int64) error {
	if err := db.Delete(&Analytics{}, ID).Error; err != nil {
		return err
	}
	return nil
}
