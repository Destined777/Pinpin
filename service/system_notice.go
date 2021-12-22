package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
)

func CreateSystemNoticeService(param http_param.CreateSystemNotice) (err error) {
	err = dao.CreateSystemNotice(param)
	return
}

func GetSystemNoticeService(email string) (res []model.System_notice, err error) {
	res, err = dao.GetSystemNotice(email)
	if err != nil {
		return
	}
	if len(res) == 0 {
		return make([]model.System_notice, 0), nil
	}
	return
}

func ReadSystemNoticeService(param http_param.ReadNoticeArgument) (err error) {
	err = dao.ReadSystemNotice(param.NoticeId)
	return
}
