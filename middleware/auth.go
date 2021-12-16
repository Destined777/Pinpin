package middleware

import (
	"Pinpin/service"
	"Pinpin/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

//检查用户是否登录的中间件
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未认证，请先登录",
			})
			return
		}
		token := strings.Fields(authHeader)
		if len(token) != 2 || token[0] != "Bearer" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "认证信息无效,请先登录",
			})
			return
		}
		claim, err := util.GetClaimFromToken(token[1])
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "认证信息无效,请先登录",
			})
			return
		}
		email := claim.(jwt.MapClaims)["email"].(string)

		//检查用户是不是已经注册
		//防止数据库中删除该用户信息但仍用原有token登录
		if _, err := service.GetUserByEmailService(email); err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "认证信息无效,请先登录",
			})
			return
		} else {
			context.Next()
		}
	}
}


//用户在激活状态才可以做的一些事情

func Activated() gin.HandlerFunc {
	return func(context *gin.Context) {
		email, _ := util.GetEmailFromAuthorization(context)
		if u, _ := service.GetUserByEmailService(email); u.IsEmailActivated {
			context.Next()
		} else {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{
				"msg": "账号未激活，请激活后使用",
			})
			return
		}
	}
}
