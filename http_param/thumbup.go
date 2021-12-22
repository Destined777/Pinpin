package http_param

import (
	"Pinpin/util"
)

type ThumbUp struct {
	ReplyId uint `form:"ReplyId" binding:"required"`
	Email   string
}

func (r *ThumbUp) GetError(err error) string {
	m := map[string]string{
		"ReplyId": "评论序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
