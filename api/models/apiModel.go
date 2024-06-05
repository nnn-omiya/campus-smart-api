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

type AddressList struct {
	Address string
	Id      int
}

func (m *ApiModel) GetControllerAddress() ([]AddressList, error) {
	DeviceList, err := lib.GetDeviceTypeList(m.DB, "controller")
	if err != nil {
		return nil, err
	}
	// もし最後のログが5分以上前のものがあれば、そのデバイスはダウンしていると判断する
	// 今ログを送るデバイスがないので、コメントアウト

	// limitTime := time.Now().Add(-5 * time.Minute)
	// fmt.Println(limitTime)
	// var aliveList []int
	// for _, v := range DeviceList {
	// 	var id int
	// 	err := m.DB.QueryRow("SELECT id FROM status_records WHERE device_id = ? AND status = 200 AND created_at <= ? ORDER BY id DESC", v, limitTime).Scan(&id)
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			name, _ := lib.GetDeviceName(m.DB, v)
	// 			fmt.Println("1")
	// 			fmt.Printf("%s status log not found\n", name)
	// 			continue
	// 		}
	// 		return nil, err
	// 	} else {
	// 		aliveList = append(aliveList, id)
	// 	}
	// }

	var addressList []AddressList
	// for _, v := range aliveList {
	for _, v := range DeviceList {
		var address string
		err := m.DB.QueryRow("SELECT address FROM device_address_records WHERE device_id = ? ORDER BY id DESC", v).Scan(&address)
		if err != nil {
			if err == sql.ErrNoRows {
				name, _ := lib.GetDeviceName(m.DB, v)
				fmt.Printf("%s IP address not found\n", name)
				continue
			}
			return nil, err
		} else {
			addressList = append(addressList, AddressList{Address: address, Id: 1})
		}
	}
	if len(addressList) == 0 {
		return nil, fmt.Errorf("no addresses found")
	}
	return addressList, nil
}

func (m *ApiModel) PostControlLog(log DeviceControl) ([]int, error) {
	DeviceList, err := lib.GetDeviceTypeList(m.DB, "controller")
	if err != nil {
		return nil, err
	}
	fmt.Println("aaa")
	var rowList []int
	for _, v := range DeviceList {
		fmt.Println(v)
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
