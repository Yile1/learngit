package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"techtrainingcamp-AppUpgrade/model"
	"techtrainingcamp-AppUpgrade/utils"
)

func Hit(c *gin.Context) {
	var downloadUrl string
	var md5 string
	var title string
	var updateTips string
	// var updateVersionCode string

	platform := c.PostForm("platform")
	channel := c.PostForm("channel")
	deviceId := c.PostForm("deviceId")
	osApi := c.PostForm("osApi")
	updateVersionCode := c.PostForm("updateVersionCode")
	cpuArch := c.PostForm("cpuArch")

	rules := model.GetAllRules()

	for index := 0; index < len(*rules); index++ {
		if !(*rules)[index].IsAvailable {
			continue
		}
		if platform != (*rules)[index].Platform || channel != (*rules)[index].Channel{
			continue
		}
		if !utils.IsStringInArray(deviceId, utils.SplitStringToList((*rules)[index].NewDeviceIdList)) {
			continue
		}
		if cpuArch != (*rules)[index].CpuArch {
			continue
		}
		if cast.ToInt(osApi) < (*rules)[index].MinOsApi || cast.ToInt(osApi) > (*rules)[index].MaxOsApi {
			continue
		}
		if !utils.IsUpdateVersionAvailable(updateVersionCode, (*rules)[index].MinUpdateVersionCode, (*rules)[index].MaxUpdateVersionCode) {
			continue
		}
		downloadUrl = (*rules)[index].DownlaodUrl
		md5  = (*rules)[index].Md5
		title = (*rules)[index].Title
		updateTips = (*rules)[index].UpdateTips

		break
	}
	c.JSON(200, gin.H{"downloadUrl": downloadUrl, "md5": md5, "title": title, "updateTips": updateTips, "updateVersionCode": updateVersionCode})
}

func AddRule(c *gin.Context)  {
	// 将Rule添加到数据库
	c.JSON(200, gin.H{"message": "Success"})
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