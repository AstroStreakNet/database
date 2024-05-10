package routes

import (
	"github.com/AstroStreakNet/database/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthenticationRoutes = func(router *mux.Router) {
	router.HandleFunc("/authentication/", controllers.CreateAuthentication).Methods("POST")
	router.HandleFunc("/authentication/", controllers.GetAllAuthentications).Methods("GET")
	router.HandleFunc("/authentication/{authID}", controllers.GetAuthenticationByID).Methods("GET")
	router.HandleFunc("/authentication/{authID}", controllers.DeleteAuthentication).Methods("DELETE")
}
