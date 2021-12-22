package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	Email      string
	Content    string
	PinpinId   uint
	ThumbUpNum int64
	IsDeleted  bool
	ReplyTo    uint //reply to another reply，if not,it will be 0
}
