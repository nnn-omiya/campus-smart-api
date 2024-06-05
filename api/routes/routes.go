package routes

import (
	"database/sql"
	"net/http"

	"github.com/nnn-omiya/campus-smart-api/controllers"
)

func Router(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/device/create_controller":
			deviceHandler := controllers.NewDeviceController(db)
			deviceHandler.PostCreateController(w, r)
		case "/device/login":
			deviceHandler := controllers.NewDeviceController(db)
			deviceHandler.PostLoginDevice(w, r)
		case "/device/status":
			deviceHandler := controllers.NewDeviceController(db)
			deviceHandler.PostDeviceStatus(w, r)
		case "/device/error":
			deviceHandler := controllers.NewDeviceController(db)
			deviceHandler.PostDeviceError(w, r)
		case "/api/device_control":
			apiHandler := controllers.NewApiController(db)
			apiHandler.PostControlDevice(w, r)
		case "/api/device_status":
			apiHandler := controllers.NewApiController(db)
			apiHandler.GetControlDevice(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}
