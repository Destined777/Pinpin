package api

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/redisclt"
	"Pinpin/service"
	"Pinpin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func SendVerificationCodeHandler(ctx *gin.Context) {
	params := http_param.SendVerificationCodeArguments{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if params.Email == "" || len(params.Email) != 10 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "request right email",
		})
		return
	}
	isResetPassword, _ := strconv.ParseBool(params.IsResetPassword)
	if !isResetPassword && dao.IsEmailActivated(params.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "该邮箱已验证，请勿重复验证。",
		})
		return
	}
	err := service.SendVerificationCodeService(params.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功发送验证码，请注意查收。如果没有收到，请稍后重试",
	})
}

//用户登录,使用jwt进行跨域身份认证
func SigninHandler(ctx *gin.Context) {
	params := http_param.SigninArguments{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err, token, res := service.SigninService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":   		"ok",
			"token": 		token,
			"information":  res,
		})
		return
	}
}

func ForgetPasswordHandler(ctx *gin.Context) {
	params := http_param.ForgetPasswordArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err := service.ForgetPasswordService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

//更换用户名
func ChangeUsernameHandler(ctx *gin.Context) {
	params := http_param.ChangeUsernameArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err := service.ChangeUsernameService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

//更换头像
func ChangeImageHandler(ctx *gin.Context) {
	params := http_param.ChangeImageArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err := service.ChangeImageService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func ChangeContactHandler(ctx *gin.Context) {
	params := http_param.ChangeContactArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err := service.ChangeContactService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func ChangeInforHandler(ctx *gin.Context) {
	params := http_param.ChangeInforArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email) //把email首字母改为大写
	if err := service.ChangeInforService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

//激活邮箱
func ActivateEmailHandler(ctx *gin.Context) {
	params := http_param.VerifyCodeArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//Verify_code is match with the specified email
	if redisclt.IsVerifyCodeMatchToRegisterAccount(params.VerifyCode, email) {
		//remove verify_code from redis where it stores if verify_code is matched with email
		redisclt.DeleteVerifyFromRedis(email)
		if err := service.ActiveEmailService(email); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "验证成功",
			})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	}
}

func CreateCompetitionHandler(ctx *gin.Context) {
	params := http_param.Competition{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Owner_email = email
	if err := service.CompetitionCreateService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetMyContactHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetMyContactService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func GetMyInforHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetMyInforService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func UpdateCompetitionHandler(ctx *gin.Context) {
	params := http_param.UpdateCompetition{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Owner_email = email
	if err := service.UpdateCompetitionService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func UpdateCompetitionNumHandler(ctx *gin.Context) {
	params := http_param.UpdateCompetitionNum{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Owner_email = email
	if err := service.UpdateCompetitionNumService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func DeleteCompetitionHandler(ctx *gin.Context) {
	params := http_param.DeleteCompetitionRequestArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	err = service.DeleteCompetitionService(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func CreateFollowPinpinHandler(ctx *gin.Context) {
	params := http_param.FollowPinpinArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err = service.CreateFollowPinpinService(params, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func DeleteFollowPinpinHandler(ctx *gin.Context) {
	params := http_param.FollowPinpinArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err = service.DeleteFollowPinpinService(params, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func CreateReplyHandler(ctx *gin.Context) {
	params := http_param.CreateReplyRequestArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	if err := service.CreateReplyService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func DeleteReplyHandler(ctx *gin.Context) {
	params := http_param.DeleteReplyRequestArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	err = service.DeleteReplyService(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetRepliesHandler(ctx *gin.Context) {
	params := http_param.ReplyListRequestArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetRepliesService(params, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func CreateThumbUpReplyHandler(ctx *gin.Context) {
	params := http_param.ThumbUp{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	if err := service.CreateThumbUpReplyService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func DeleteThumbUpReplyHandler(ctx *gin.Context) {
	params := http_param.ThumbUp{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	if err := service.DeleteThumbUpReplyService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetMyPinpinHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetMyPinpinService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func GetMyFollowHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetMyFollowService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func CreateReportHandler(ctx *gin.Context) {
	params := http_param.ReportParam{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	params.Email = email
	if err := service.CreateReportService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetNoticeHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetNoticeService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func ReadNoticeHandler(ctx *gin.Context) {
	params := http_param.ReadNoticeArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err := service.ReadNoticeService(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func CreateSystemNoticeHandler(ctx *gin.Context) {
	params := http_param.CreateSystemNotice{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if err := service.CreateSystemNoticeService(params); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetSystemNoticeHandler(ctx *gin.Context) {
	email, err := util.GetEmailFromAuthorization(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res, err := service.GetSystemNoticeService(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": res,
		})
		return
	}
}

func ReadSystemNoticeHandler(ctx *gin.Context) {
	params := http_param.ReadNoticeArgument{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	err := service.ReadSystemNoticeService(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}
