package model

import (
	"time"
)

type Competition_pinpin struct {
	PinpinId                 uint `gorm:"primarykey"`
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
	Contact_notnull          int64 //联系方式至少选一项触发器
	IsDeleted                bool
	Owner_email              string
	ReplyNum                 int64
	CreatedAt                int64
	UpdatedAt                int64
}
