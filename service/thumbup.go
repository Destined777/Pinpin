package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"errors"
	"gorm.io/gorm"
)

func CreateThumbUpReplyService(param http_param.ThumbUp) (err error) {
	_, err = dao.GetReplyById(param.ReplyId)
	if err != nil {
		err = errors.New("点赞的评论不存在")
		return
	}
	thumbup := model.ThumbUp {
		Model:				gorm.Model{},
		Email:				param.Email,
		ReplyId:  			param.ReplyId,
		IsThumbUp: 			true,
	}
	err = dao.CreateThumbUpReply(thumbup)
	return
}

func DeleteThumbUpReplyService(param http_param.ThumbUp) (err error) {
	_, err = dao.GetReplyById(param.ReplyId)
	if err != nil {
		err = errors.New("该评论已被删除")
		return
	}
	err = dao.DeleteThumbUpReply(param)
	return
}