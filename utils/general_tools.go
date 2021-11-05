package utils

import (
	"strconv"
	"strings"
	"techtrainingcamp-AppUpgrade/common"
)

func SplitDeviceStringToList(str string)  []string{
	return strings.Split(str,",")
}

func IsStringInArray(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element{
			return true
		}
	}
	return false
}

func IsDeviceIdInWhiteList(deviceId string, newDeviceIdListStr string, deletedDeviceIdListStr string) bool {
	newDeviceIdList := SplitDeviceStringToList(newDeviceIdListStr)
	deletedDeviceIdList := SplitDeviceStringToList(deletedDeviceIdListStr)

	if IsStringInArray(deviceId, newDeviceIdList) {
		return true
	}

	if IsStringInArray(deviceId, deletedDeviceIdList) {
		return false
	}

	RDC := common.GetRdc()
	if RDC.SIsMember("WhiteList", deviceId).Val() {
		return true
	}

	return false
}

func CompareVersion(Version1 []int, Version2 []int) bool {
	// return true if Version1 < Version2
	i := 0
	for i < len(Version1) && i < len(Version2) {
		if Version1[i] < Version2[i] {
			return true
		}else if Version1[i] > Version2[i]{
			return false
		}
		i++
	}
	for i < len(Version1) {
		if Version1[i] != 0{
			return false
		}
	}
	for i < len(Version2) {
		if Version2[i] != 0{
			return true
		}
	}
	return false // Version1 = Version2
}

func IsUpdateVersionAvailable(updateVersionCode string, minUpdateVersionCode string, maxUpdateVersionCode string) bool {
	// compare rules
	// return true if minUpdateVersionCode < updateVersionCode < maxUpdateVersionCode
	currVersionStr := strings.Split(updateVersionCode, ".")
	minVersionStr := strings.Split(minUpdateVersionCode, ".")
	maxVersionStr := strings.Split(maxUpdateVersionCode, ".")
	currVersion := make([]int, len(currVersionStr), cap(currVersionStr))
	minVersion := make([]int, len(minVersionStr), cap(minVersionStr))
	maxVersion := make([]int ,len(maxVersionStr), cap(maxVersionStr))
	for i := 0; i < len(currVersionStr); i++ {
		currVersion[i], _ = strconv.Atoi(currVersionStr[i])
	}
	for i := 0; i < len(minVersionStr); i++ {
		minVersion[i], _ = strconv.Atoi(minVersionStr[i])
	}
	for i := 0; i < len(maxVersionStr); i++ {
		maxVersion[i], _ = strconv.Atoi(maxVersionStr[i])
	}
	flag1 := CompareVersion(minVersion, currVersion)
	flag2 := CompareVersion(currVersion, maxVersion)
	return flag1 && flag2
}