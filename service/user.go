package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/redisclt"
	"Pinpin/util"
	"errors"
	"log"
)

func SendVerificationCodeService(email string) (err error) {
	isExist, err := dao.ExistUser(email)
	if err != nil {
		return errors.New("内部服务器错误")
	}
	if !isExist {
		return errors.New("不存在该成员")
	}
	go func() {
		err := util.SendEmail(email)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("send success")
		}
	}()
	return nil
}

func SigninService(param http_param.SigninArguments) (err error, token string, name string, image int64) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return
	}
	if !isExist { //此邮箱尚未注册
		err = errors.New("此邮箱尚未注册")
		return
	}
	isMatch := dao.IsEmailAndPasswordMatch(param)
	if !isMatch {
		err = errors.New("邮箱或密码错误")
		return
	}
	token, _ = util.GenerateTokenByJwt(param.Email)
	res, err := dao.GetUser(param.Email)
	name = res.Username
	image = res.Image
	return
}

func ForgetPasswordService(param http_param.ForgetPasswordArgument) (err error) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == false { //此邮箱尚未注册
		return errors.New("此邮箱尚未注册")
	}
	if !redisclt.IsVerifyCodeMatchToRegisterAccount(param.VerifyCode, param.Email) {
		return errors.New("验证码过期或不正确")
	}
	err = dao.ResetPassword(param.Email, param.NewPassword)
	return
}

func ChangeUsernameService(param http_param.ChangeUsernameArgument) (err error) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == false { //此邮箱尚未注册
		return errors.New("此邮箱尚未注册")
	}
	err = dao.ResetUsername(param)
	return
}

func ChangeImageService(param http_param.ChangeImageArgument) (err error) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == false { //此邮箱尚未注册
		return errors.New("此邮箱尚未注册")
	}
	err = dao.ResetImage(param)
	return
}

func ChangeContactService(param http_param.ChangeContactArgument) (err error) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == false { //此邮箱尚未注册
		return errors.New("此邮箱尚未注册")
	}
	err = dao.ResetContact(param)
	return
}

func ChangeInforService(param http_param.ChangeInforArgument) (err error) {
	isExist, err := dao.ExistUser(param.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == false { //此邮箱尚未注册
		return errors.New("此邮箱尚未注册")
	}
	err = dao.ResetInfor(param)
	return
}

func ActiveEmailService(email string) (err error) {
	err = dao.ActiveEmail(email)
	return
}

func GetUserByEmailService(email string) (user model.AuthUser, err error) {
	user, err = dao.GetUser(email)
	return
}

func CreateUserService(user http_param.AuthUserArguments) (err error) {
	authUser := model.AuthUser{
		Email:            user.Email,
		Username:         user.Username,
		Password:         user.Password,
		Image:            user.Image,
		IsEmailActivated: false,
	}
	err = dao.CreateUser(authUser)
	if err != nil {
		switch err.Error() {
		case "此邮箱已经被注册过了":
			return
		default:
			return errors.New("服务器内部发生错误")
		}
	}
	return nil
}

func GetMyContactService(email string) (res http_param.ChangeContactArgument, err error) {
	res, err = dao.GetMyContact(email)
	if err == nil {
		return
	} else if err.Error() == "联系方式为空" {
		return
	} else {
		err = errors.New("服务器内部发生错误")
	}
	return
}

func GetMyInforService(email string) (res http_param.ChangeInforArgument, err error) {
	res, err = dao.GetMyInfor(email)
	if err == nil {
		return
	} else if err.Error() == "个人经历为空" {
		return
	} else {
		err = errors.New("服务器内部发生错误")
	}
	return
}
