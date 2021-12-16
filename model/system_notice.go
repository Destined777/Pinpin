package model

import (
	"Pinpin/global"
	"gorm.io/gorm"
)

type System_notice struct {
	gorm.Model
	Title     string
	Content   string
	Email     string
	Timestamp int64
	IsRead  bool
}

func (sn System_notice) GetSystemNotices(email string) (notices []System_notice, err error) {
	err = global.DB.Where("Email=? or Email='toall'", email).Find(&notices).Error
	return
}

func (sn System_notice) ReadSystemNotice(id uint) (err error){
	var system_notice System_notice
	err = global.DB.Where("ID=?", id).Find(&system_notice).Error
	if err != nil {
		return
	}
	system_notice.IsRead = true
	err = global.DB.Save(&system_notice).Error
	return
}

func (sn System_notice) CreateNotice(title string, content string, userEmail string, timestamp int64) (err error) {
	systemNotice := System_notice{
		Title:     title,
		Content:   content,
		Email:     userEmail,
		Timestamp: timestamp,
		IsRead:  false,
	}
	err = global.DB.Create(&systemNotice).Error
	return
}

