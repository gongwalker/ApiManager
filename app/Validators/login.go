package Validators

import (
	"github.com/go-playground/validator/v10"
)

// 登录参数校验，验证错误自定义返回错误文案
func LogInGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "Username" {
			switch err[n].Tag() {
			case "required":
				return "Please input username"
			}
		}
		if err[n].Field() == "Password" {
			switch err[n].Tag() {
			case "required":
				return "Please input password"
			case "min":
				return "The password is less than " + err[n].Param() + " characters in length"
			}
		}
	}

	return "Parameter error"
}
