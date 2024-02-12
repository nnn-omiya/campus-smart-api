package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nnn-omiya/campus-smart-api/lib"
	"github.com/nnn-omiya/campus-smart-api/models"
)

type ApiController struct {
	Model *models.ApiModel
}

func NewApiController(db *sql.DB) *ApiController {
	m := models.NewApiModel(db)
	return &ApiController{Model: m}
}

func (h *ApiController) PostControlDevice(w http.ResponseWriter, r *http.Request) {
	controlMode, err := lib.AccessLogHandler[models.DeviceControl](r)
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}
	fmt.Println(controlMode)
	addressList, err := h.Model.GetControllerAddress()
	if err != nil {
		lib.ErrorHandler(w, err, http.StatusInternalServerError)
		return
	}

	intList, err := h.Model.PostControlLog(*controlMode)
	if err != nil {
		return
	}
	for _, address := range addressList {
		url := fmt.Sprintf("http:%s/ac?power=%v&mode=%d&temperature=%d&direction=%d&volume=%d",
			address, controlMode.Power,
			controlMode.Detail.Mode,
			controlMode.Detail.Temperature,
			controlMode.Detail.Direction,
			controlMode.Detail.Volume)
		fmt.Println(url)
	}

	json.NewEncoder(w).Encode(intList)
}

func (h *ApiController) GetControlDevice(w http.ResponseWriter, r *http.Request) {}
