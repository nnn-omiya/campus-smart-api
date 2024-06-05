package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/nnn-omiya/campus-smart-api/controllers"
)

func Router(db *sql.DB, router *mux.Router) {
	deviceRouter(db, router)
	apiRouter(db, router)
}

func deviceRouter(db *sql.DB, router *mux.Router) {
	deviceHandler := controllers.NewDeviceController(db)

	router.HandleFunc("/device/create_controller", deviceHandler.PostCreateController).Methods("POST")
	router.HandleFunc("/device/login", deviceHandler.PostLoginDevice).Methods("POST")
	router.HandleFunc("/device/status", deviceHandler.PostDeviceStatus).Methods("POST")
	router.HandleFunc("/device/error", deviceHandler.PostDeviceError).Methods("POST")
}

func apiRouter(db *sql.DB, router *mux.Router) {
	apiHandler := controllers.NewApiController(db)

	router.HandleFunc("/api/device_control", apiHandler.PostControlDevice).Methods("POST")
	router.HandleFunc("/api/device_status", apiHandler.GetControlDevice).Methods("GET")
}
