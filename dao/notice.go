package dao

import (
	"Pinpin/global"
	"Pinpin/model"
	"errors"
)

func CreateNotice(email string, pinpinId uint, replyTo uint, Type int, time int64) (err error) {
	notice := model.Notice{
		Email:            email,
		PinpinId:         pinpinId,
		ReplyId:          replyTo,
		Type:             Type,
		CreatedTimestamp: time,
		IsRead:           false,
	}
	err = global.DB.Create(&notice).Error
	return
}

func GetNotice(email string) (res []model.Notice, err error) {
	err = global.DB.Where("Email=?", email).Order("Created_at DESC").Find(&res).Error
	return
}

func ReadNotice(id uint) (err error) { //阅读通知
	var notice model.Notice
	global.DB.Model(&model.Notice{}).Where("ID=?", id).Find(&notice)
	if (notice == model.Notice{}) {
		err = errors.New("没有该条通知")
		return err
	}
	notice.IsRead = true
	err = global.DB.Save(&notice).Error
	return
}
