package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
)

func GetNoticeService(email string) (res []model.Notice, err error) {
	res, err = dao.GetNotice(email)
	if err != nil {
		return
	}
	if len(res) == 0 {
		return make([]model.Notice, 0), nil
	}
	return
}

func ReadNoticeService(param http_param.ReadNoticeArgument) (err error) {
	err = dao.ReadNotice(param.NoticeId)
	return
}
