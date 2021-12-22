package dao

import (
	"Pinpin/global"
	"Pinpin/model"
)

//获取用户的所有拼拼帖，包括过期的和已经删除的
func GetMyPinpin(email string) (pinpins []model.Competition_pinpin, err error) {
	err = global.DB.Model(&model.Competition_pinpin{}).Where("Owner_email=?", email).Find(&pinpins).Error
	return
}

//判断用户是否收藏过这个拼拼帖子
func IsFollowed(pinpinId uint, email string) bool {
	var temp model.Follow
	global.DB.Where(map[string]interface{}{"Pinpin_Id": pinpinId, "Email": email}).Find(&temp)
	return temp.IsFollow
}

//获取用户的收藏帖子
func GetMyFollow(email string) (follows []model.Follow, err error) {
	err = global.DB.Model(&model.Follow{}).Where("Email=? and Is_Follow=?", email, true).Find(&follows).Error
	return
}
