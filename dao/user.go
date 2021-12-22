package dao

import (
	"Pinpin/global"
	"Pinpin/http_param"
	"Pinpin/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(user model.AuthUser) (err error) {
	isExist, err := ExistUser(user.Email)
	if err != nil { //查询出错
		return err
	}
	if isExist == true { //邮箱已经存在
		return errors.New("此邮箱已经被注册过了")
	}
	err = global.DB.Create(user).Error
	return
}

//处理查询出错，对未查询到与查询到用户均不报错而只是分类
func ExistUser(email string) (isExist bool, err error) {
	user := model.AuthUser{}
	err = global.DB.Where("Email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func IsEmailActivated(email string) (isEmailActivated bool) {
	isExist, _ := ExistUser(email)
	if isExist == false {
		return false
	}
	user := model.AuthUser{}
	err := global.DB.Where("Email = ?", email).First(&user).Error
	if err != nil && user.IsEmailActivated == true {
		return true
	} else {
		return false
	}
}

func IsEmailAndPasswordMatch(param http_param.SigninArguments) bool {
	u, _ := GetUser(param.Email)
	if u.Password != param.Password {
		return false
	}
	return true
}

func ResetPassword(email string, newPassword string) (err error) {
	var user model.AuthUser
	global.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" { //保证查找到了该用户
		return errors.New("No such user in database")
	}
	user.Password = newPassword
	if !user.IsEmailActivated {
		user.IsEmailActivated = true
	}
	err = global.DB.Save(&user).Error
	return
}

func ResetUsername(param http_param.ChangeUsernameArgument) (err error) {
	var user model.AuthUser
	global.DB.Where("email = ?", param.Email).Find(&user)
	if user.Email == "" { //保证查找到了该用户
		return errors.New("No such user in database")
	}
	user.Username = param.NewUsername
	err = global.DB.Save(&user).Error
	return
}

func ResetImage(param http_param.ChangeImageArgument) (err error) {
	var user model.AuthUser
	global.DB.Where("email = ?", param.Email).Find(&user)
	if user.Email == "" { //保证查找到了该用户
		return errors.New("No such user in database")
	}
	user.Image = param.NewImage
	err = global.DB.Save(&user).Error
	return
}

func ResetContact(param http_param.ChangeContactArgument) (err error) {
	var user model.AuthUser
	global.DB.Where("email = ?", param.Email).Find(&user)
	if user.Email == "" { //保证查找到了该用户
		return errors.New("No such user in database")
	}
	user.Contact_wechat = param.Contact_wechat
	user.Contact_qq = param.Contact_qq
	user.Contact_tel = param.Contact_tel
	err = global.DB.Save(&user).Error
	return
}

func ResetInfor(param http_param.ChangeInforArgument) (err error) {
	var user model.AuthUser
	global.DB.Where("email = ?", param.Email).Find(&user)
	if user.Email == "" { //保证查找到了该用户
		return errors.New("No such user in database")
	}
	user.Master_name = param.Master_name
	user.Master_sex = param.Master_sex
	user.Master_gradeandmajor = param.Master_gradeandmajor
	user.Master_introduction = param.Master_introduction
	err = global.DB.Save(&user).Error
	return
}

func GetMyContact(email string) (res http_param.ChangeContactArgument, err error) {
	err = global.DB.Model(&model.AuthUser{}).Select(
		"Email",
		"Contact_wechat",
		"Contact_qq",
		"Contact_tel").Where("Email=?", email).Take(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("联系方式为空")
	}
	return
}

func GetMyInfor(email string) (res http_param.ChangeInforArgument, err error) {
	err = global.DB.Model(&model.AuthUser{}).Select(
		"Email",
		"Master_name",
		"Master_sex",
		"Master_gradeandmajor",
		"Master_introduction").Where("Email=?", email).Take(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("个人经历为空")
	}
	fmt.Println(res)
	return
}

func ActiveEmail(email string) (err error) {
	var user model.AuthUser
	global.DB.Where("Email=?", email).Find(&user)
	if user.Email == "" {
		return errors.New("No such user in database")
	}
	if user.IsEmailActivated {
		return errors.New("This emailAes is activated")
	}
	user.IsEmailActivated = true
	return global.DB.Save(&user).Error
}

func GetUser(email string) (user model.AuthUser, err error) {
	err = global.DB.Where("Email=?", email).Find(&user).Error
	return
}
