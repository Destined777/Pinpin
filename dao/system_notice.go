package dao

import (
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
)

func CreateSystemNotice(param http_param.CreateSystemNotice) (err error) {
	var temp model.System_notice
	err = temp.CreateNotice(param.Title, param.Content, param.Email, util.GetTimeStamp())
	return
}

func GetSystemNotice(email string) (res []model.System_notice, err error) {
	var temp model.System_notice
	res, err = temp.GetSystemNotices(email)
	return
}

func ReadSystemNotice(id uint) (err error) {
	var temp model.System_notice
	err = temp.ReadSystemNotice(id)
	return
}