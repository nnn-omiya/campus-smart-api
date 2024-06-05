package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/nnn-omiya/campus-smart-api/lib"
	"github.com/nnn-omiya/campus-smart-api/models"
)

type DeviceController struct {
	Model *models.DeviceModel
}

func NewDeviceController(db *sql.DB) *DeviceController {
	m := models.NewDeviceModel(db)
	return &DeviceController{Model: m}
}

func (h *DeviceController) PostCreateController(w http.ResponseWriter, r *http.Request) {
	createDevice, err := lib.AccessLogHandler[models.CreateController](r)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	id, err := h.Model.CreateController(*createDevice)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *DeviceController) PostLoginDevice(w http.ResponseWriter, r *http.Request) {
	deviceAddress, err := lib.AccessLogHandler[models.DeviceAddress](r)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	id, err := h.Model.InsertAddress(*deviceAddress)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *DeviceController) PostDeviceStatus(w http.ResponseWriter, r *http.Request) {
	deviceStatus, err := lib.AccessLogHandler[models.DeviceStatus](r)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}

	id, err := h.Model.InsertStatus(*deviceStatus)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(id)
}

func (h *DeviceController) PostDeviceError(w http.ResponseWriter, r *http.Request) {
	deviceError, err := lib.AccessLogHandler[models.DeviceError](r)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusBadRequest)
	}

	id, err := h.Model.InsertErrorLog(*deviceError)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(id)
}
