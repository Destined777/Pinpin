package model

import (
	"Pinpin/global"
	"errors"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Email           string
	PinpinId		uint
	Content         string
	CreateTimestamp int64
}

func (temp Report) CreateReport(email string, pinpinId uint, content string, timestamp int64) (err error) {
	report := Report{Email: email, PinpinId: pinpinId, Content: content, CreateTimestamp: timestamp}
	global.DB.Where(map[string]interface{}{"Email": email, "PinpinId": pinpinId}).Order("CreateTimestamp").Find(&temp)
	if temp == (Report{}) || timestamp-temp.CreateTimestamp > 10371037 {//举报间隔时间设置
		global.DB.Create(&report)
		return nil
	} else {
		return errors.New("您已经举报过该拼拼帖,我们会尽快处理，请不要过于频繁的举报")
	}
}

