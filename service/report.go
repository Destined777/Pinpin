package service

import (
	"Pinpin/dao"
	"Pinpin/http_param"
)

func CreateReportService(param http_param.ReportParam) (err error) {
	err = dao.CreateReport(param)
	return
}
