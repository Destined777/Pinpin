package dao

import (
	"Pinpin/global"
	"Pinpin/model"
	"errors"
)

func CreateFollowPinpin(follow model.Follow) (err error) {
	err = global.DB.Create(&follow).Error
	return
}

func DeleteFollowPinpin(follow model.Follow) (err error) {
	temp := model.Follow{}
	global.DB.Where(map[string]interface{}{"pinpin_id": follow.PinpinId, "email": follow.Email}).Find(&temp)
	if (temp == model.Follow{}) || !temp.IsFollow {
		err = errors.New("你没有关注该拼拼帖")
		return
	} else {
		temp.IsFollow = false
		err = global.DB.Save(&temp).Error
		return
	}
}