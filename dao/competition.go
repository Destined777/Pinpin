package dao

import (
	"Pinpin/global"
	"Pinpin/model"
	"time"
)

func CompetitionCreate(competition model.Competition_recruit) (err error) {
	err = global.DB.Create(&competition).Error
	return
}

func GetAllCompetitionsDetails() (res []model.ComperitionEntails, err error) {
	err = global.DB.Model(&model.Competition_recruit{}).Select("title",
						"created_at",
								"contact_qq",
								"contact_wechat",
								"contact_tel",
								"deadline",
								"competition_introduction",
								"demanding_num",
								"now_num",
								"demanding_introduction",
								"master_name",
								"master_sex",
								"master_gradeandmajor",
								"master_introduction",
								"teammate_introduction").Where("deadline >= ?",time.Now()).Order("created_at DESC").Find(&res).Error
	return
}