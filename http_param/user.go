package http_param

import (
	"Pinpin/global"
	"Pinpin/util"
	"errors"
)

type AuthUserArguments struct {
	Email              	string `gorm:"column:'Email';primaryKey;"`
	Username           	string `form:"Username"binding:"required"`
	Password        	string `form:"Password"binding:"required"`
	Image				int64  `form:"Image"binding:"required"`
}

func (r *AuthUserArguments) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
		"Username":		"用户名",
		"Password":		"密码",
		"Image":		"头像",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
func (r *AuthUserArguments) GetUser(email string) (user AuthUserArguments, err error) {
	err = global.DB.Where("Email=?", email).Find(&user).Error
	return
}

func (r *AuthUserArguments) ResetUserPassword(emailAes string, passwordHash string) (err error) {
	var doc AuthUserArguments
	global.DB.Where("Email = ?", emailAes).Find(&doc)
	if doc.Email == "" {
		return errors.New("No such user in database")
	}
	doc.Password = passwordHash
	err = global.DB.Save(&doc).Error
	return
}

type SendVerificationCodeArguments struct {
	Email   			string `form:"Email" binding:"required"`
	IsResetPassword 	string `form:"IsResetPassword" binding:"required"`
}
func (v *SendVerificationCodeArguments) GetError(err error) string {
	m := map[string]string {
		"Email":			"电子邮件",
		"IsResetPassword":	"是否为重设密码",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type SigninArguments struct {
	Email   			string `form:"Email" binding:"required"`
	Password        	string `form:"Password"binding:"required"`
}

func (r *SigninArguments) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
		"Password":		"密码",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ForgetPasswordArgument struct {
	Email       string `form:"Email" binding:"required"`
	VerifyCode  string `form:"VerifyCode" binding:"required"`
	NewPassword string `form:"NewPassword" binding:"required"`
}

func (r *ForgetPasswordArgument) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
		"VerifyCode":	"验证码",
		"NewPassword":	"新密码",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ChangeUsernameArgument struct {
	Email       string `form:"Email" binding:"required"`
	NewUsername string `form:"NewUsername" binding:"required"`
}

func (r *ChangeUsernameArgument) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
		"NewUsername":	"新用户名",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ChangeImageArgument struct {
	Email       string `form:"Email" binding:"required"`
	NewImage 	int64 `form:"NewImage" binding:"required"`
}

func (r *ChangeImageArgument) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
		"NewImage":		"新头像",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ChangeContactArgument struct {
	Email       				string `form:"Email" binding:"required"`
	Contact_qq					string `form:"Contact_qq" binding:"omitempty"`
	Contact_wechat				string `form:"Contact_wechat" binding:"omitempty"`
	Contact_tel 				string `form:"Contact_tel" binding:"omitempty"`
}

func (r *ChangeContactArgument) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type ChangeInforArgument struct {
	Email       				string `form:"Email" binding:"required"`
	Master_name					string `form:"Master_name" binding:"omitempty"`
	Master_sex					string `form:"Master_sex" binding:"omitempty"`
	Master_gradeandmajor		string `form:"Master_gradeandmajor" binding:"omitempty"`
	Master_introduction			string `form:"Master_introduction" binding:"omitempty"`
}

func (r *ChangeInforArgument) GetError(err error) string {
	m := map[string]string{
		"Email":		"电子邮件",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type VerifyCodeArgument struct {
	VerifyCode  string `form:"VerifyCode" binding:"required"`
}

func (r *VerifyCodeArgument) GetError(err error) string {
	m := map[string]string{
		"VerifyCode":	"验证码",
	}
	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}