package dao

import (
	"Pinpin/http_param"
	"Pinpin/model"
	"Pinpin/util"
)

func CreateReport(param http_param.ReportParam) (err error) {
	var temp model.Report
	now := util.GetTimeStamp()
	err = temp.CreateReport(param.Email, param.PinpinId, param.Content, now)
	return
}
