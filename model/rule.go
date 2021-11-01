package model

import "github.com/jinzhu/gorm"

type Rule struct {
	gorm.Model
	// Rule Condition
	Platform string `json:"platform"`
	NewDeviceIdList string `json:"new_device_id_list"`
	DeletedDeviceIdList string `json:"deleted_device_id_list"`
	MaxUpdateVersionCode string `json:"max_update_version_code"`
	MinUpdateVersionCode string `json:"min_update_version_code"`
	MaxOsApi int `json:"max_os_api"`
	MinOsApi int `json:"min_os_api"`
	CpuArch string `json:"cpu_arch"`
	Channel string `json:"channel"`

	// Return Value
	DownlaodUrl string `json:"downlaod_url"`
	Md5 string `json:"md5"`
	Title string `json:"title"`
	UpdateTips string `json:"update_tips"`
	UpdateVersionCode string `json:"update_version_code"`

	// Available
	IsAvailable bool `json:"is_available"`
}
