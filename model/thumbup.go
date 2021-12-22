package model

import "gorm.io/gorm"

type ThumbUp struct { //给评论点赞
	gorm.Model
	Email     string
	ReplyId   uint
	IsThumbUp bool
}
