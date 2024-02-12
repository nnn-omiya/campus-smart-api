package models

import (
	"database/sql"
	"fmt"

	"github.com/nnn-omiya/campus-smart-api/lib"
)

type ApiModel struct {
	DB *sql.DB
}
type Api struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

func NewApiModel(DB *sql.DB) *ApiModel {
	return &ApiModel{DB: DB}
}

type detail struct {
	Mode        int
	Temperature int
	Direction   int
	Volume      int
}

type DeviceControl struct {
	Power  bool
	Detail detail
}

func (m *ApiModel) GetControllerAddress() ([]string, error) {
	DeviceList, err := lib.GetDeviceTypeList(m.DB, "controller")
	if err != nil {
		return nil, err
	}
	var addressList []string
	for _, v := range DeviceList {
		var address string
		err := m.DB.QueryRow("SELECT address FROM device_address_records WHERE device_id = ?", v).Scan(&address)
		if err != nil {
			if err == sql.ErrNoRows {
				name, _ := lib.GetDeviceName(m.DB, v)
				return nil, fmt.Errorf("%s IP address not found", name)
			}
			return nil, err
		} else {
			addressList = append(addressList, address)
		}
	}
	return addressList, nil
}

func (m *ApiModel) PostControlLog(log DeviceControl) ([]int, error) {
	DeviceList, err := lib.GetDeviceTypeList(m.DB, "controller")
	if err != nil {
		return nil, err
	}
	var rowList []int
	for _, v := range DeviceList {
		result, err := m.DB.Exec("INSERT INTO status_records (device_id, control_record) VALUES (?, ?)", v, log)
		if err != nil {
			return nil, nil
		} else {
			id, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}
			rowList = append(rowList, int(id))
		}
	}
	return rowList, nil
}
