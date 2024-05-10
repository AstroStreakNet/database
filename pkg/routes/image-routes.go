package routes

import (
	"github.com/AstroStreakNet/database/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterImageRoutes = func(router *mux.Router) {
	router.HandleFunc("/image/", controllers.CreateImage).Methods("POST")
	router.HandleFunc("/image/", controllers.GetAllImages).Methods("GET")
	router.HandleFunc("/image/{imageID}", controllers.GetImageByID).Methods("GET")
	router.HandleFunc("/image/{imageID}", controllers.UpdateImage).Methods("PUT")
	router.HandleFunc("/image/{imageID}", controllers.DeleteImage).Methods("DELETE")
}
