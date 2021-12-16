package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"errors"
	"gorm.io/gorm"
)
//关注拼拼帖
func CreateFollowPinpinService(param http_param.FollowPinpinArgument, email string) (err error) {
	_ , err = dao.GetPinpinById(param.PinpinId)
	if err != nil {
		return errors.New("未找到该拼拼帖")
	}
	follow := model.Follow{
		Email:				email,
		PinpinId:			param.PinpinId,
		IsFollow:			true,
	}
	err = dao.CreateFollowPinpin(follow)
	if err != nil {
		return errors.New("服务器发生错误")
	}
	return nil
}
//取消关注拼拼帖
func DeleteFollowPinpinService(param http_param.FollowPinpinArgument, email string) (err error) {
	_ , err = dao.GetPinpinById(param.PinpinId)
	if err != nil {
		return errors.New("未找到该拼拼帖")
	}
	follow := model.Follow{
		Model:				gorm.Model{},
		Email:				email,
		PinpinId:			param.PinpinId,
		IsFollow:			false,
	}
	err = dao.DeleteFollowPinpin(follow)
	if err != nil {
		return err
	} else {
		return nil
	}
}