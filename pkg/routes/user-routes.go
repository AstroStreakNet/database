package routes

import (
	"github.com/AstroStreakNet/database/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{userID}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userID}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userID}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{astronomerID}", controllers.GetUserByAstronomerID).Methods("GET")
}
