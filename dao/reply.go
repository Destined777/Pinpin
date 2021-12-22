package dao

import (
	"Pinpin/global"
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
	"gorm.io/gorm"
)

func GetReplyById(id uint) (reply model.Reply, err error) {
	err = global.DB.Where("ID=?", id).Find(&reply).Error
	return
}

func UpdateReplyNum(id uint) error { //拼拼帖的评论数量加一
	pinpin, _ := GetPinpinById(id)
	pinpin.ReplyNum++
	err := global.DB.Save(&pinpin).Error
	return err
}

func CreateReply(reply model.Reply) (err error) {
	err = global.DB.Create(&reply).Error
	if err != nil {
		return
	}
	pinpin, _ := GetPinpinById(reply.PinpinId)
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		e := UpdateReplyNum(reply.PinpinId)
		if e != nil {
			return e
		}
		//如果ReplyTo不为0，即用户回复了另一个用户，此时通知该被回复的用户
		if reply.ReplyTo != 0 {
			re, _ := GetReplyById(reply.ReplyTo)
			if re.Email != reply.Email {
				err = CreateNotice(re.Email, reply.PinpinId, reply.ID, 2, util.GetTimeStamp())
				if err != nil {
					return err
				}
			}
		} else if pinpin.Owner_email != reply.Email {
			//如果帖主与评论者不同，会给帖主发送通知
			err = CreateNotice(pinpin.Owner_email, reply.PinpinId, reply.ID, 1, util.GetTimeStamp())
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func DeleteReply(pinpinId uint, replyId uint) (err error) {
	var reply model.Reply
	err = global.DB.Where("Pinpin_Id=? and ID=?", pinpinId, replyId).Find(&reply).Error
	if err != nil {
		return nil
	} else { //更改IsDeleted字段
		pinpin, _ := GetPinpinById(pinpinId)
		pinpin.ReplyNum--
		global.DB.Save(&pinpin)
		reply.IsDeleted = true
		err = global.DB.Save(&reply).Error
		return err
	}
}

//获取全部评论，包括已经删除了的
func GetReplies(param http_param.ReplyListRequestArgument) (results []model.Reply, err error) {
	err = global.DB.Where("Pinpin_Id=?", param.PinpinId).Order("Created_at DESC").Find(&results).Error
	return
}

func IsThumbUp(id uint, email string) bool { //确定该用户是否点赞过该条评论
	var temp model.ThumbUp
	global.DB.Select("is_thumb_up").Where(map[string]interface{}{"Reply_Id": id, "Email": email}).Find(&temp)
	return temp.IsThumbUp
}
