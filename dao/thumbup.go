package dao

import (
	"Pinpin/global"
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
	"errors"
)

func CreateThumbUpReply(thumbup model.ThumbUp) (err error) {
	var temp model.ThumbUp
	global.DB.Where(map[string]interface{}{"Email": thumbup.Email, "Reply_Id": thumbup.ReplyId}).Find(&temp)
	if (temp == model.ThumbUp{}) {//没有找到点赞记录
		err = global.DB.Create(&thumbup).Error
		if err == nil {
			var reply model.Reply
			global.DB.Where("id=?", thumbup.ReplyId).Find(&reply)
			reply.ThumbUpNum++
			err = global.DB.Save(&reply).Error
			CreateNotice(thumbup.Email, 0,thumbup.ReplyId, 0, util.GetTimeStamp())
		}
	} else if temp.IsThumbUp {//已经点赞过了
		err = errors.New("您已经为这条评论点过赞了")
	} else {//点赞被取消过
		var reply model.Reply
		global.DB.Where("ID=?", thumbup.ReplyId).Find(&reply)
		reply.ThumbUpNum++
		global.DB.Save(&reply)
		temp.IsThumbUp = true
		err = global.DB.Save(&temp).Error
	}
	return
}

func DeleteThumbUpReply(thumbup http_param.ThumbUp) (err error) {
	var temp model.ThumbUp
	global.DB.Where(map[string]interface{}{"Email": thumbup.Email, "Reply_Id": thumbup.ReplyId}).Find(&temp)
	if (temp == model.ThumbUp{}) {//没有找到点赞记录
		err = errors.New("您没有为该评论点过赞")
	} else if temp.IsThumbUp == true {//已经点赞过了
		var reply model.Reply
		global.DB.Where("ID=?", thumbup.ReplyId).Find(&reply)
		reply.ThumbUpNum--
		global.DB.Save(&reply)
		temp.IsThumbUp = false
		err = global.DB.Save(&temp).Error
	} else {//点赞被取消过
		err = errors.New("您的点赞已经被取消过了")
	}
	return
}