package models

import (
	"database/sql"

	"github.com/nnn-omiya/campus-smart-api/lib"
)

type DeviceModel struct {
	DB *sql.DB
}

type CreateController struct {
	Addr string `json:"address"`
	Type int    `json:"type"`
}

type DeviceAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DeviceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DeviceError struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewDeviceModel(DB *sql.DB) *DeviceModel {
	return &DeviceModel{DB: DB}
}

func (m *DeviceModel) CreateController(device CreateController) (int, error) {
	result, err := m.DB.Exec("INSERT INTO devices (mac_address, device_type) VALUES (?, ?)", device.Addr, "controller")
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	result, err = m.DB.Exec("INSERT INTO devices_controller (device_id, device_type) VALUES (?, ?)", id, device.Type)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *DeviceModel) InsertAddress(address DeviceAddress) (int, error) {
	device_id, err := lib.GetDeviceId(m.DB, address.Name)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec("INSERT INTO device_address_records (device_id, address) VALUES (?, ?)", device_id, address.Address)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *DeviceModel) InsertStatus(status DeviceStatus) (int, error) {
	device_id, err := lib.GetDeviceId(m.DB, status.Name)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec("INSERT INTO status_records (device_id, status) VALUES (?, ?)", device_id, status.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (m *DeviceModel) InsertErrorLog(deviceError DeviceError) (int, error) {
	device_id, err := lib.GetDeviceId(m.DB, deviceError.Name)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec("INSERT INTO error_message_records (device_id, error_type, message) VALUES (?, ?, ?)", device_id, deviceError.Type, deviceError.Message)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}
