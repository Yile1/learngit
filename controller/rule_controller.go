package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"techtrainingcamp-AppUpgrade/common"
	"techtrainingcamp-AppUpgrade/model"
	"techtrainingcamp-AppUpgrade/utils"
)

func Hit(c *gin.Context) {
	DB := common.GetDB()

	var downloadUrl string
	var md5 string
	var title string
	var updateTips string
	var rules []model.Rule
	// var updateVersionCode string

	platform := c.PostForm("platform")
	channel := c.PostForm("channel")
	deviceId := c.PostForm("device_id")
	osApi := c.PostForm("os_api")
	updateVersionCode := c.PostForm("update_version_code")
	cpuArch := c.PostForm("cpu_arch")

	DB.Find(&rules)
	//rules = model.GetAllRules()

	for index := 0; index < len(rules); index++ {
		if !rules[index].IsAvailable {
			continue
		}
		if platform != rules[index].Platform || channel != rules[index].Channel{
			continue
		}
		if !utils.IsStringInArray(deviceId, utils.SplitDeviceStringToList(rules[index].NewDeviceIdList)) {
			continue
		}
		if cpuArch != rules[index].CpuArch {
			continue
		}
		if cast.ToInt(osApi) < rules[index].MinOsApi || cast.ToInt(osApi) > rules[index].MaxOsApi {
			continue
		}
		if !utils.IsUpdateVersionAvailable(updateVersionCode, rules[index].MinUpdateVersionCode, rules[index].MaxUpdateVersionCode) {
			continue
		}
		downloadUrl = rules[index].DownlaodUrl
		md5  = rules[index].Md5
		title = rules[index].Title
		updateTips = rules[index].UpdateTips

		break
	}
	c.JSON(200, gin.H{"download_url": downloadUrl, "md5": md5, "title": title, "update_tips": updateTips, "update_version_code": updateVersionCode})
}

func AddRule(c *gin.Context)  {
	// 将Rule添加到数据库
	DB := common.GetDB()

	platform := c.PostForm("platform")
	updateVersionCode := c.PostForm("update_version_code")
	md5 := c.PostForm("md_5")
	newDeviceIdList := c.PostForm("new_device_id_list")
	deletedDeviceIdList := c.PostForm("deleted_device_id_list")
	maxUpdateVersionCode := c.PostForm("max_update_version_code")
	minUpdateVersionCode := c.PostForm("min_update_version_code")
	maxOsApi := cast.ToInt(c.PostForm("max_os_api"))
	minOsApi := cast.ToInt(c.PostForm("min_os_api"))
	cpuArch := c.PostForm("cpu_arch")
	channel := c.PostForm("channel")
	title := c.PostForm("title")
	updateTips := c.PostForm("update_tips")
	downloadUrl := c.PostForm("download_url")

	newRule := model.Rule{
		Platform:             platform,
		UpdateVersionCode:    updateVersionCode,
		Md5:                  md5,
		NewDeviceIdList:      newDeviceIdList,
		DeletedDeviceIdList:  deletedDeviceIdList,
		MaxUpdateVersionCode: maxUpdateVersionCode,
		MinUpdateVersionCode: minUpdateVersionCode,
		MaxOsApi:             maxOsApi,
		MinOsApi:             minOsApi,
		CpuArch:              cpuArch,
		Channel:              channel,
		Title:                title,
		UpdateTips:           updateTips,
		DownlaodUrl:          downloadUrl,
		IsAvailable:          true,
	}

	DB.Create(&newRule)

	c.JSON(200, gin.H{"message": "Success", "newRule": newRule})
}

func DeleteRule(c *gin.Context)  {
	// 将Rule从数据库删除
	c.JSON(200, gin.H{"message": "Success"})
}

func DisableRule(c *gin.Context)  {
	// 将Rule禁用
	c.JSON(200, gin.H{"message": "Success"})
}

func EnableRule(c *gin.Context)  {
	// 将Rule启用
	c.JSON(200, gin.H{"message": "Success"})
}