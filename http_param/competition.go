package http_param

import (
	"Pinpin/util"
	"time"
)

type Competition struct {
	PinpinId                 uint      `gorm:"column:pinpin_id;-;AUTO_INCREMENT;comment:'主键'"`
	Title                    string    `form:"Title" binding:"required"`
	Contact_qq               string    `form:"Contact_qq" binding:"omitempty"`
	Contact_wechat           string    `form:"Contact_wechat" binding:"omitempty"`
	Contact_tel              string    `form:"Contact_tel" binding:"omitempty"`
	Deadline                 time.Time `form:"Deadline" binding:"required"`
	Competition_introduction string    `form:"Competition_introduction" binding:"omitempty"`
	Demanding_num            int64     `form:"Demanding_num" binding:"required"`
	Now_num                  int64     `form:"Now_num" binding:"required"`
	Demanding_introduction   string    `form:"Demanding_introduction" binding:"required,max=200"`
	Master_name              string    `form:"Master_name" binding:"omitempty"`
	Master_sex               string    `form:"Master_sex" binding:"omitempty"`
	Master_gradeandmajor     string    `form:"Master_gradeandmajor" binding:"omitempty"`
	Master_introduction      string    `form:"Master_introduction" binding:"omitempty,max=100"`
	Teammate_introduction    string    `form:"Teammate_introduction" binding:"omitempty,max=100"`
	Contact_notnull          int64
	Owner_email              string
	CreatedAt                int64
	UpdatedAt                int64
}

func (r *Competition) GetError(err error) string {
	m := map[string]string{
		"Title":                  "比赛名称",
		"Deadline":               "截止日期",
		"Demanding_num":          "总需求人数",
		"Now_num":                "已有人数",
		"Demanding_introduction": "需求介绍",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type UpdateCompetition struct {
	PinpinId                 uint      `form:"PinpinId" binding:"required"`
	Title                    string    `form:"Title" binding:"required"`
	Contact_qq               string    `form:"Contact_qq" binding:"omitempty"`
	Contact_wechat           string    `form:"Contact_wechat" binding:"omitempty"`
	Contact_tel              string    `form:"Contact_tel" binding:"omitempty"`
	Deadline                 time.Time `form:"Deadline" binding:"required"`
	Competition_introduction string    `form:"Competition_introduction" binding:"omitempty"`
	Demanding_num            int64     `form:"Demanding_num" binding:"required"`
	Now_num                  int64     `form:"Now_num" binding:"required"`
	Demanding_introduction   string    `form:"Demanding_introduction" binding:"required,max=200"`
	Master_name              string    `form:"Master_name" binding:"omitempty"`
	Master_sex               string    `form:"Master_sex" binding:"omitempty"`
	Master_gradeandmajor     string    `form:"Master_gradeandmajor" binding:"omitempty"`
	Master_introduction      string    `form:"Master_introduction" binding:"omitempty,max=100"`
	Teammate_introduction    string    `form:"Teammate_introduction" binding:"omitempty,max=100"`
	Contact_notnull          int64
	Owner_email              string
	CreatedAt                int64
	UpdatedAt                int64
}

func (r *UpdateCompetition) GetError(err error) string {
	m := map[string]string{
		"PinpinId":               "拼拼帖序号",
		"Title":                  "比赛名称",
		"Deadline":               "截止日期",
		"Demanding_num":          "总需求人数",
		"Now_num":                "已有人数",
		"Demanding_introduction": "需求介绍",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type UpdateCompetitionNum struct { //一键更新拼拼帖人数
	PinpinId      uint  `form:"PinpinId" binding:"required"`
	Demanding_num int64 `form:"Demanding_num" binding:"required"`
	Now_num       int64 `form:"Now_num" binding:"required"`
	Owner_email   string
}

func (r *UpdateCompetitionNum) GetError(err error) string {
	m := map[string]string{
		"PinpinId":      "拼拼帖序号",
		"Demanding_num": "总需求人数",
		"Now_num":       "已有人数",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type DeleteCompetitionRequestArgument struct {
	Email    string
	PinpinId uint `form:"PinpinId" binding:"required,gte=0"`
}

func (r *DeleteCompetitionRequestArgument) GetError(err error) string {
	m := map[string]string{
		"PinpinId": "拼拼帖序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ComperitionEntails struct { //帖子详情页
	PinpinId                 uint
	Title                    string
	Contact_qq               string
	Contact_wechat           string
	Contact_tel              string
	Deadline                 time.Time
	Competition_introduction string
	Demanding_num            int64
	Now_num                  int64
	Demanding_introduction   string
	Master_name              string
	Master_sex               string
	Master_gradeandmajor     string
	Master_introduction      string
	Teammate_introduction    string
	Owner_email              string
	ReplyNum                 int64
	CreatedAt                int64
	UpdatedAt                int64
}

type SearchCompetitions struct {
	Title string `form:"Title" binding:"required"`
}

func (r *SearchCompetitions) GetError(err error) string {
	m := map[string]string{
		"Title": "搜索标题",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
