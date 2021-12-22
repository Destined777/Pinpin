package redisclt

import (
	"Pinpin/consts"
	"fmt"
	"time"
)

func StoreEmailAndVerifyCodeInRedis(verifyCode string, email string) {
	Set(consts.REDIS_VERIFY_CODE_SUFFIX+email, verifyCode, consts.VERIFYCODE_VALID_TIME*time.Second)
	re := Get(consts.REDIS_VERIFY_CODE_SUFFIX + email)
	fmt.Println(re.Val())
	if re.Val() == verifyCode {
		fmt.Println("store succeeded")
	}
}

//IsVerifyCodeMatchToRegisterAccount Check if verify code and user specified by email inputs are match
func IsVerifyCodeMatchToRegisterAccount(verifyCode string, email string) (IsMatch bool) {
	re := Get(consts.REDIS_VERIFY_CODE_SUFFIX + email)
	fmt.Println("\n\n", re.Val())
	if re.Val() == "" {
		fmt.Println("get failed")
	}
	if re.Val() == verifyCode && re.Val() != "" {
		IsMatch = true
	} else {
		IsMatch = false
	}
	fmt.Println("IsMatch: ", IsMatch)
	return
}
func DeleteVerifyFromRedis(email string) {
	Del(consts.REDIS_VERIFY_CODE_SUFFIX + email)
}
