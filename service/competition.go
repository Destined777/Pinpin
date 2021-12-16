package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
	"errors"
	"gorm.io/gorm"
)

func CompetitionCreateService(param http_param.Competition) (err error) {
	if param.Contact_qq == "" && param.Contact_wechat == "" && param.Contact_tel == "" { //联系方式三选一触发器
		param.Contact_notnull = 0
		return errors.New("联系方式不可均为空")
	} else {
		param.Contact_notnull = 1
	}
	competition := model.Competition_pinpin{
		Title:						param.Title,
		Contact_qq:					param.Contact_qq,
		Contact_wechat:				param.Contact_wechat,
		Contact_tel: 				param.Contact_tel,
		Deadline:					param.Deadline,
		Competition_introduction:	param.Competition_introduction,
		Demanding_num:				param.Demanding_num,
		Now_num:					param.Now_num,
		Demanding_introduction:		param.Demanding_introduction,
		Master_name:				param.Master_name,
		Master_sex:					param.Master_sex,
		Master_gradeandmajor:		param.Master_gradeandmajor,
		Master_introduction:		param.Master_introduction,
		Teammate_introduction:		param.Teammate_introduction,
		Contact_notnull:			param.Contact_notnull,
		IsDeleted:					false,
		Owner_email:				param.Owner_email,
		ReplyNum: 					0,
		CreatedAt: 					util.GetTimeStamp(),
		UpdatedAt:   				util.GetTimeStamp(),
	}
	err = dao.CompetitionCreate(competition)
	if err != nil {
		return errors.New("服务器发生错误")
	}
	return nil
}

func UpdateCompetitionService(param http_param.UpdateCompetition) (err error) {
	if param.Contact_qq == "" && param.Contact_wechat == "" && param.Contact_tel == "" { //联系方式三选一触发器
		param.Contact_notnull = 0
		return errors.New("联系方式不可均为空")
	} else {
		param.Contact_notnull = 1
	}
	err = dao.UpdateCompetition(param)
	if err == gorm.ErrRecordNotFound {
		return errors.New("未找到该拼拼帖")
	}
	if err != nil {
		err = errors.New("服务器内部发生错误")
	}
	return err
}

func UpdateCompetitionNumService(param http_param.UpdateCompetitionNum) (err error) {
	err = dao.UpdateCompetitionNum(param)
	if err == gorm.ErrRecordNotFound {
		return errors.New("未找到该拼拼帖")
	}
	if err != nil {
		err = errors.New("服务器内部发生错误")
	}
	return err
}

func DeleteCompetitionService(param http_param.DeleteCompetitionRequestArgument) (err error) {
	err = dao.DeleteCompetition(param)
	if err != nil {
		err = errors.New("服务器内部发生错误")
	}
	return err
}

func GetCompetitionDetailsService() (res []http_param.ComperitionEntails, err error) {
	res, err = dao.GetAllCompetitionsDetails()
	if err != nil {
		return
	}
	if len(res) == 0 {
		return make([]http_param.ComperitionEntails, 0), nil
	}
	return
}

func SearchCompetitionsService(title string) (res []http_param.ComperitionEntails, err error) {
	res, err = dao.SearchCompetitions(title)
	if err != nil {
		return
	}
	if len(res) == 0 {
		return make([]http_param.ComperitionEntails, 0), nil
	}
	return
}