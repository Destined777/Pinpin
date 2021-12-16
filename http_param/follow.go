package http_param

import "Pinpin/util"

type FollowPinpinArgument struct {
	 PinpinId		uint `form:"PinpinId" binding:"required"`
}

func (r *FollowPinpinArgument) GetError(err error) string {
	m := map[string]string{
		"PinpinId":		"拼拼帖序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}