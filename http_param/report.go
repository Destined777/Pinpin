package http_param

import "Pinpin/util"

type ReportParam struct {
	PinpinId		uint   `form:"PinpinId" binding:"required,gte=0"`
	Email   		string
	Content         string `form:"Content" binding:"required"`
}

func (r *ReportParam) GetError(err error) string {
	m := map[string]string{
		"Content":		"评论内容",
		"PinpinId":		"拼拼帖序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}