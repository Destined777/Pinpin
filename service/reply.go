package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"errors"
	"gorm.io/gorm"
)

func CreateReplyService(param http_param.CreateReplyRequestArgument) (err error) {
	_ , err = dao.GetPinpinById(param.PinpinId)
	if err != nil {
		return errors.New("未找到该拼拼帖")
	}
	reply := model.Reply{
		Model:				gorm.Model{},
		Email:				param.Email,
		Content: 			param.Content,
		PinpinId:			param.PinpinId,
		ThumbUpNum:			0,
		IsDeleted: 			false,
		ReplyTo: 			param.ReplyTo,
	}
	err = dao.CreateReply(reply)
	return err
}

func DeleteReplyService(param http_param.DeleteReplyRequestArgument) (err error) {
	reply, err := dao.GetReplyById(param.ReplyId)
	//没有找到对应的评论
	if err != nil {
		return errors.New("未找到该条评论")
	}
	//不能删除别人的评论
	if param.Email != reply.Email {
		return errors.New("不能删除他人的评论")
	}
	//该条评论已被删除
	if reply.IsDeleted == true {
		return errors.New("该条评论已被删除")
	}
	//删除评论
	err = dao.DeleteReply(param.PinpinId, param.ReplyId)
	return err
}

func GetRepliesService(param http_param.ReplyListRequestArgument, email string) (results []map[string]interface{}, err error) {
	pinpin, e := dao.GetPinpinById(param.PinpinId)
	if e != nil || pinpin.IsDeleted {
		return nil, errors.New("该拼拼帖不存在")
	}
	replies, err := dao.GetReplies(param)
	if err != nil {
		return
	}
	var isthumbup bool
	for _, temp := range replies {
		isthumbup = dao.IsThumbUp(temp.ID, email)
		results = append(results, map[string]interface{}{
			"ID":			temp.ID,
			"Created_at":	temp.CreatedAt,
			"Email":		temp.Email,
			"Content":		temp.Content,
			"PinpinId":		temp.PinpinId,
			"ThumbUpNum":	temp.ThumbUpNum,
			"IsDeleted":	temp.IsDeleted,
			"ReplyTo":		temp.ReplyTo,
			//用户是否点赞过这条评论
			"IsThumbUp":	isthumbup,
		})
	}
	return results, nil
}
