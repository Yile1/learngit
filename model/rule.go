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

func GetAllRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		Platform:             "Android",
		UpdateVersionCode:    "10.0.1",
		Md5:                  "1234567",
		NewDeviceIdList:      "123,456,789",
		DeletedDeviceIdList:  "",
		MaxUpdateVersionCode: "10.0.3",
		MinUpdateVersionCode: "10.0.0",
		MaxOsApi:             10,
		MinOsApi:             8,
		CpuArch:              "64",
		Channel:              "222222",
		Title:                "欢迎升级",
		UpdateTips:           "在升级之后，请及时给予我们反馈哦",
		DownlaodUrl:          "www.baidu.com",
		IsAvailable:          true,
	})

	rules = append(rules, Rule{
		Platform:             "ios",
		UpdateVersionCode:    "10.0.1",
		Md5:                  "1234567",
		NewDeviceIdList:      "123,456",
		DeletedDeviceIdList:  "",
		MaxUpdateVersionCode: "10.0.3",
		MinUpdateVersionCode: "10.0.0",
		MaxOsApi:             10,
		MinOsApi:             8,
		CpuArch:              "64",
		Channel:              "111111",
		Title:                "欢迎升级",
		UpdateTips:           "在升级之后，请及时给予我们反馈哦",
		DownlaodUrl:          "www.baidu.com",
		IsAvailable:          true,
	})

	return &rules
}
