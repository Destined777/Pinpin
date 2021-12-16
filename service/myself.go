package service

import (
	"Pinpin/dao"
	"errors"
)

func GetMyPinpinService(email string) (results []map[string]interface{}, err error) {
	pinpins, err := dao.GetMyPinpin(email)
	if err != nil {
		err = errors.New("服务器内部发生错误")
		return
	}
	var isfollow bool
	for _, value := range pinpins {
		isfollow = dao.IsFollowed(value.PinpinId, email)
		results = append(results, map[string]interface{}{
			"PinpinId":					value.PinpinId,
			"Title":					value.Title,
			"Contact_qq":				value.Contact_qq,
			"Contact_wechat":			value.Contact_wechat,
			"Contact_tel": 				value.Contact_tel,
			"Deadline":					value.Deadline,
			"Competition_introduction":	value.Competition_introduction,
			"Demanding_num":			value.Demanding_num,
			"Now_num":					value.Now_num,
			"Demanding_introduction":	value.Demanding_introduction,
			"Master_name":				value.Master_name,
			"Master_sex":				value.Master_sex,
			"Master_gradeandmajor":		value.Master_gradeandmajor,
			"Master_introduction":		value.Master_introduction,
			"Teammate_introduction":	value.Teammate_introduction,
			"IsDeleted":          		value.IsDeleted,
			"Owner_email":				value.Owner_email,
			"ReplyNum":           		value.ReplyNum,
			//用户是否收藏该拼拼帖
			"IsFollow":					isfollow,
		})
	}
	return
}

func GetMyFollowService(email string) (results []map[string]interface{}, err error) {
	follows, err := dao.GetMyFollow(email)
	if err != nil {
		err = errors.New("服务器内部发生错误")
		return
	}
	for _, val := range follows {
		value, err := dao.GetPinpinById(val.PinpinId)
		if err == nil {
			isfollow := dao.IsFollowed(value.PinpinId, email)
			results = append(results, map[string]interface{} {
				"PinpinId":					value.PinpinId,
				"Title":					value.Title,
				"Contact_qq":				value.Contact_qq,
				"Contact_wechat":			value.Contact_wechat,
				"Contact_tel": 				value.Contact_tel,
				"Deadline":					value.Deadline,
				"Competition_introduction":	value.Competition_introduction,
				"Demanding_num":			value.Demanding_num,
				"Now_num":					value.Now_num,
				"Demanding_introduction":	value.Demanding_introduction,
				"Master_name":				value.Master_name,
				"Master_sex":				value.Master_sex,
				"Master_gradeandmajor":		value.Master_gradeandmajor,
				"Master_introduction":		value.Master_introduction,
				"Teammate_introduction":	value.Teammate_introduction,
				"IsDeleted":          		value.IsDeleted,
				"Owner_email":				value.Owner_email,
				"ReplyNum":           		value.ReplyNum,
				//用户是否收藏该拼拼帖
				"IsFollow":					isfollow,
			})
		}
	}
	return
}