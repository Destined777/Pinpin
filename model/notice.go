package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Email       		string
	PinpinId			uint//拼拼帖的序号，如果是给评论点赞，此项为0
	ReplyId				uint//评论的序号，若为点赞，则表示点赞的评论号
	Type            	int//0为点赞，1为评论，2为评论的回复
	CreatedTimestamp 	int64
	IsRead          	bool
}
