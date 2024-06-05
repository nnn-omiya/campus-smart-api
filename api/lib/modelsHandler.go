package lib

import (
	"database/sql"
	"fmt"
)

func GetDeviceId(m *sql.DB, name string) (int, error) {
	var device_id int
	err := m.QueryRow("SELECT id FROM devices WHERE mac_address = ?", name).Scan(&device_id)
	fmt.Println(device_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("device %s not found", name)
		}
		return 0, err
	}
	return device_id, nil
}

func GetDeviceName(m *sql.DB, id int) (string, error) {
	var device_name string
	err := m.QueryRow("SELECT mac_address FROM devices WHERE id = ?", id).Scan(&device_name)
	fmt.Println(device_name)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("device id %d not found", id)
		}
		return "", err
	}
	return device_name, nil
}

func GetDeviceTypeList(m *sql.DB, deviceType string) ([]int, error) {
	var deviceIDs []int

	rows, err := m.Query("SELECT id FROM devices WHERE device_type = ? AND control_flg = ?", deviceType, true)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		deviceIDs = append(deviceIDs, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deviceIDs, nil
}
