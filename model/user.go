package model

type AuthUser struct {
	Email              			string `gorm:"primary_key"`
	Username           			string
	Password        			string
	Image						int64//存头像图片的序号
	IsEmailActivated    		bool//邮箱是否被验证
	Contact_qq					string
	Contact_wechat				string
	Contact_tel 				string
	Master_name					string
	Master_sex					string
	Master_gradeandmajor		string
	Master_introduction			string
}
