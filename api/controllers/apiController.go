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

type URLResponse struct {
	URL      string
	Response *http.Response
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

	var urlResponses []URLResponse

	for _, address := range addressList {
		url := fmt.Sprintf("http://%s/ac?id=%d&power=%d&mode=%d&temperature=%d&direction=%d&volume=%d",
			address.Address,
			address.Id,
			func() int {
				if controlMode.Power {
					return 1
				}
				return 0
			}(),
			controlMode.Detail.Mode,
			controlMode.Detail.Temperature,
			controlMode.Detail.Direction,
			controlMode.Detail.Volume)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making GET request: %v\n", err)
			continue
		}

		urlResponse := URLResponse{URL: url, Response: resp}
		urlResponses = append(urlResponses, urlResponse)

		resp.Body.Close()
	}
	fmt.Println(urlResponses)

	json.NewEncoder(w).Encode(intList)
}

func (h *ApiController) GetControlDevice(w http.ResponseWriter, r *http.Request) {}
