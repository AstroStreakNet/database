package routes

import (
	"github.com/AstroStreakNet/database/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAnalyticsRoutes = func(router *mux.Router) {
	router.HandleFunc("/analytics/", controllers.CreateAnalytics).Methods("POST")
	router.HandleFunc("/analytics/", controllers.GetAllAnalytics).Methods("GET")
	router.HandleFunc("/analytics/{analyticsID}", controllers.GetAnalyticsById).Methods("GET")
	router.HandleFunc("/analytics/{analyticsID}", controllers.DeleteAnalytics).Methods("DELETE")
}
