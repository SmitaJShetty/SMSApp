package router

import (
	"SMSApp/internal"

	"github.com/gorilla/mux"
)

// addRoutes adds routes
func addRoutes(router *mux.Router) {
	router.HandleFunc("/send", internal.SendMessageHandler).Methods("POST").Schemes("http")
	router.HandleFunc("/ping", internal.PingHandler).Methods("GET").Schemes("http")
}

// GetRouter gets router
func GetRouter() *mux.Router {
	router := mux.NewRouter()
	addRoutes(router)
	return router
}
