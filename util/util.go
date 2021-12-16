package util

import (
	"Pinpin/consts"
	"Pinpin/email"
	"Pinpin/global"
	"Pinpin/redisclt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"strings"
	"time"
)

func ErrorHandler(err error, m map[string]string) (msg string) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return "输入参数错误"
	}
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErrs {
			switch val.Tag() {
			case "required":
				return m[val.Field()] + "不能为空哦"
			}
		}
	}
	return ""
}

//生成6位数验证码
func GenerateVerificationCode() (code string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func SendEmail(Email string) (err error) {
	emailClient := global.Config.EmailSettings
	to := Email + "@hust.edu.cn"
	subject := "Pinpin验证码"
	vcode := GenerateVerificationCode()
	redisclt.StoreEmailAndVerifyCodeInRedis(vcode, Email)

	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>拼拼验证码</title>
		</head>
		<body>
			验证码为：` + vcode + `.
		</body>
		</html>`

	sendUserName := "Pinpin"//发送邮件的人名称
	fmt.Println("send email")
	err = email.SendEmail(emailClient.User, sendUserName, emailClient.Password, emailClient.Host, to, subject, body, "html")
	return
}

func GetTimeStamp() (t int64) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t = time.Now().In(loc).Unix()
	return
}

//GenerateTokenByJwt 使用Jwt生成token作身份认证使用
//其中存储了登录的时间，登陆的email
func GenerateTokenByJwt(email string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"email":     email,
		"timeStamp": GetTimeStamp(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(consts.TOKEN_SCRECT_KEY))
	return
}

func GetEmailFromAuthorization(c *gin.Context) (email string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("您未登录，请登陆后查看")
		}
	}()
	authHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authHeader)[1]
	if err != nil {
		err = errors.New("您未登录，请登陆后查看")
		return
	}
	claim, _ := GetClaimFromToken(token)
	email = claim.(jwt.MapClaims)["email"].(string)
	return
}

func GetClaimFromToken(tokenString string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(consts.TOKEN_SCRECT_KEY), err
	})
	if err != nil {
		return nil, err
	} else {
		claims = token.Claims.(jwt.MapClaims)
		return claims, nil
	}
}