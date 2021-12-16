package api

import (
	"Pinpin/http_param"
	"Pinpin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreateUserHandler(ctx *gin.Context) {
	params := http_param.AuthUserArguments{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	params.Email = strings.Title(params.Email)//把email首字母改为大写
	if err := service.CreateUserService(params); err != nil {
		switch err.Error() {
		case "此邮箱已经被注册过了":
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func GetCompetitionDetailsHandler(ctx *gin.Context) {
	res, err := service.GetCompetitionDetailsService()
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

func SearchCompetitionsHandler(ctx *gin.Context) {
	params := http_param.SearchCompetitions{}
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": params.GetError(err),
		})
		return
	}
	res, err := service.SearchCompetitionsService(params.Title)
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