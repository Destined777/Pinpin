package http_param

import "Pinpin/util"

type CreateReplyRequestArgument struct {
	Email      string
	Content    string `form:"Content" binding:"required"`
	PinpinId   uint   `form:"PinpinId" binding:"required,gte=0"`
	ThumbUpNum int64
	IsDeleted  bool
	ReplyTo    uint `form:"ReplyTo,default=0"`
}

func (r *CreateReplyRequestArgument) GetError(err error) string {
	m := map[string]string{
		"Content":  "评论内容",
		"PinpinId": "拼拼帖序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type DeleteReplyRequestArgument struct {
	Email    string
	PinpinId uint `form:"PinpinId" binding:"required,gte=0"`
	ReplyId  uint `form:"ReplyId" binding:"required,gte=0"`
}

func (r *DeleteReplyRequestArgument) GetError(err error) string {
	m := map[string]string{
		"PinpinId": "拼拼帖序号",
		"ReplyId":  "评论序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ReplyListRequestArgument struct {
	PinpinId uint `form:"PinpinId" binding:"required"`
}

func (r *ReplyListRequestArgument) GetError(err error) string {
	m := map[string]string{
		"PinpinId": "拼拼帖序号",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
