package dao

import (
	"Pinpin/global"
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
	"errors"
	"fmt"
	"time"
)

func CompetitionCreate(competition model.Competition_pinpin) (err error) {
	err = global.DB.Create(&competition).Error
	fmt.Println(competition)
	fmt.Println(competition.PinpinId)
	return
}

func UpdateCompetition(param http_param.UpdateCompetition) (err error) {
	var pinpin model.Competition_pinpin
	err = global.DB.Where("pinpin_id=?", param.PinpinId).Find(&pinpin).Error
	if err != nil {
		return
	}
	if param.Owner_email != pinpin.Owner_email {
		err = errors.New("不能更新他人的拼拼帖")
	}
	pinpin.Teammate_introduction = param.Teammate_introduction
	pinpin.Master_introduction = param.Master_introduction
	pinpin.Master_gradeandmajor = param.Master_gradeandmajor
	pinpin.Master_sex = param.Master_sex
	pinpin.Master_name = param.Master_name
	pinpin.Demanding_introduction = param.Demanding_introduction
	pinpin.Demanding_num = param.Demanding_num
	pinpin.Now_num = param.Now_num
	pinpin.Competition_introduction = param.Competition_introduction
	pinpin.Deadline = param.Deadline
	pinpin.Contact_tel = param.Contact_tel
	pinpin.Contact_wechat = param.Contact_wechat
	pinpin.Contact_qq = param.Contact_qq
	pinpin.Title = param.Title
	pinpin.UpdatedAt = util.GetTimeStamp()
	err = global.DB.Save(&pinpin).Error
	return
}

func UpdateCompetitionNum(param http_param.UpdateCompetitionNum) (err error) {
	var pinpin model.Competition_pinpin
	global.DB.Where("pinpin_id = ?", param.PinpinId).Find(&pinpin)
	pinpin.Demanding_num = param.Demanding_num
	pinpin.Now_num = param.Now_num
	pinpin.UpdatedAt = util.GetTimeStamp()
	err = global.DB.Save(&pinpin).Error
	return
}

func DeleteCompetition(param http_param.DeleteCompetitionRequestArgument) (err error) {
	pinpin, err := GetPinpinById(param.PinpinId)
	if err != nil {
		return
	}
	if param.Email != pinpin.Owner_email {
		err = errors.New("不能删除他人的拼拼帖")
		return
	}
	if pinpin.IsDeleted == true {
		err = errors.New("该拼拼帖已经被删除了")
	}
	pinpin.IsDeleted = true
	err = global.DB.Save(&pinpin).Error
	return
}

func GetAllCompetitionsDetails() (res []http_param.ComperitionEntails, err error) {
	err = global.DB.Model(&model.Competition_pinpin{}).Select("pinpin_id",
		"Title",
		"Created_at",
		"Updated_at",
		"Contact_qq",
		"Contact_wechat",
		"Contact_tel",
		"Deadline",
		"Competition_introduction",
		"Demanding_num",
		"Now_num",
		"Demanding_introduction",
		"Master_name",
		"Master_sex",
		"Master_gradeandmajor",
		"Master_introduction",
		"Teammate_introduction",
		"Owner_email",
		"reply_num", ).Where("Is_deleted = ? and Deadline >= ?", 0, time.Now()).Order("Updated_at DESC").Find(&res).Error
	return
}

func SearchCompetitions(title string) (res []http_param.ComperitionEntails, err error) {
	err = global.DB.Model(&model.Competition_pinpin{}).Select("pinpin_id",
		"Title",
		"Created_at",
		"Updated_at",
		"Contact_qq",
		"Contact_wechat",
		"Contact_tel",
		"Deadline",
		"Competition_introduction",
		"Demanding_num",
		"Now_num",
		"Demanding_introduction",
		"Master_name",
		"Master_sex",
		"Master_gradeandmajor",
		"Master_introduction",
		"Teammate_introduction",
		"Owner_email",
		"reply_num").Where("Is_deleted = ? and Deadline >= ? and Title LIKE ?", 0, time.Now(), "%"+title+"%").Order("Updated_at DESC").Find(&res).Error
	return
}
func GetPinpinById(pinpinId uint) (Pinpin model.Competition_pinpin, err error) {
	err = global.DB.Model(&model.Competition_pinpin{}).Where("pinpin_id=?", pinpinId).Find(&Pinpin).Error
	return
}
