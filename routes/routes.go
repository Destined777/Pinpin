package routes

import (
	"Pinpin/middleware"
	"Pinpin/routes/api"
	"github.com/gin-gonic/gin"
)

func UsePinpinRouter(r *gin.Engine) {
	Pinpinapi := r.Group("/api")
	{
		welcome := Pinpinapi.Group("/welcome")
		{
			welcome.POST("/createUser", api.CreateUserHandler)
			welcome.GET("/getCompetitions", api.GetCompetitionDetailsHandler)
			welcome.GET("/searchCompetitions", api.SearchCompetitionsHandler)
			welcome.GET("/getUserInfor", api.GetUserInforHandler)
		}
		manage := Pinpinapi.Group("/manage")
		{
			manage.POST("/sendVerifyCode", api.SendVerificationCodeHandler)
			manage.POST("/signin", api.SigninHandler)
			manage.POST("/activation", middleware.Auth(), api.ActivateEmailHandler)
			manage.GET("/getMyContact", middleware.Auth(), api.GetMyContactHandler)
			manage.GET("/getMyInfor", middleware.Auth(), api.GetMyInforHandler)

			change := manage.Group("/change", middleware.Auth())
			{
				change.PUT("/changePassword", api.ForgetPasswordHandler)
				change.PUT("/changeUsername", api.ChangeUsernameHandler)
				change.PUT("/changeImage", api.ChangeImageHandler)
				change.PUT("/changeContact", api.ChangeContactHandler)
				change.PUT("/changeInfor", api.ChangeInforHandler)
			}
			recruit := manage.Group("/recruit", middleware.Auth(), middleware.Activated())
			{
				recruit.POST("/createCompetition", api.CreateCompetitionHandler)
				recruit.PUT("/updateCompetition", api.UpdateCompetitionHandler)
				recruit.PUT("/updateCompetitionNum", api.UpdateCompetitionNumHandler)
				recruit.DELETE("/deleteCompetition", api.DeleteCompetitionHandler)
			}
			follow := manage.Group("/follows", middleware.Auth())
			{
				follow.POST("/createFollow", api.CreateFollowPinpinHandler)
				follow.DELETE("/deleteFollow", api.DeleteFollowPinpinHandler)
			}
			reply := manage.Group("/replies", middleware.Auth())
			{
				reply.POST("/createReply", middleware.Activated(), api.CreateReplyHandler)
				reply.GET("/getReplies", api.GetRepliesHandler)
				reply.DELETE("/deleteReply", api.DeleteReplyHandler)
			}
			thumbup := manage.Group("/thumbups", middleware.Auth())
			{
				thumbup.POST("/createThumbUp", api.CreateThumbUpReplyHandler)
				thumbup.DELETE("/deleteThumbUp", api.DeleteThumbUpReplyHandler)
			}
			report := manage.Group("/reports", middleware.Auth())
			{
				report.POST("/createReport", api.CreateReportHandler)
			}
			notice := manage.Group("/notices", middleware.Auth())
			{
				notice.GET("/getNotice", api.GetNoticeHandler)
				notice.POST("/readNotice", api.ReadNoticeHandler)
			}
			sys_notice := manage.Group("/systemNotice", middleware.Auth())
			{
				sys_notice.POST("/createSystemNotice", api.CreateSystemNoticeHandler)
				sys_notice.GET("/getSystemNotice", api.GetSystemNoticeHandler)
				sys_notice.POST("/readSystemNotice", api.ReadSystemNoticeHandler)
			}
			myself := manage.Group("/myself", middleware.Auth())
			{
				myself.GET("/getMypinpin", api.GetMyPinpinHandler)
				myself.GET("/getMyfollow", api.GetMyFollowHandler)
			}
		}
	}
}
