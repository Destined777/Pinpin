package http_param

import "Pinpin/util"

type ReadNoticeArgument struct {
	NoticeId uint `form:"NoticeId" binding:"required"`
}

func (r *ReadNoticeArgument) GetError(err error) string {
	m := map[string]string{
		"NoticeId": "通知序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
