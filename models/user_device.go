package models

type UserDevice struct {
	ID        string `gorm:"primary_key" json:"id"`
	UserId    string `json:"user_id"`
	DeviceOs  string `json:"device_os"`
	DeviceMac string `json:"device_mac"`
	Model
}
