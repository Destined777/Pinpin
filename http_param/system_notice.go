package http_param

import "Pinpin/util"

type CreateSystemNotice struct {
	Title   string `form:"Title" binding:"required"`
	Content string `form:"Content" binding:"required"`
	Email   string `form:"Email" binding:"required"`
}

func (r *CreateSystemNotice) GetError(err error) string {
	m := map[string]string{
		"Title":   "系统通知标题",
		"Content": "系统通知内容",
		"Email":   "被通知者",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
