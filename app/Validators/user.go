package Validators

import (
	"github.com/go-playground/validator/v10"
)

// 添加用户，验证错误自定义返回错误文案
func AddUserGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "LoginName" {
			switch err[n].Tag() {
			case "required":
				return "Please input login_name"
			case "max":
				return "The login_name is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The login_name is more than " + err[n].Param() + " characters in length"
			}

		}
		if err[n].Field() == "LoginPwd" {
			switch err[n].Tag() {
			case "required":
				return "Please input password"
			case "max":
				return "The password is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The password is more than " + err[n].Param() + " characters in length"
			}
		}

		if err[n].Field() == "Role" {
			switch err[n].Tag() {
			case "required":
				return "Please input role"
			case "max":
				return "The role is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The role is more than " + err[n].Param() + " characters in length"
			}
		}
	}

	return "Parameter error"
}

// 编辑用户，验证错误自定义返回错误文案
func EditUserGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "LoginName" {
			switch err[n].Tag() {
			case "required":
				return "Please input login_name"
			case "max":
				return "The login_name is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The login_name is more than " + err[n].Param() + " characters in length"
			}

		}

		if err[n].Field() == "Role" {
			switch err[n].Tag() {
			case "required":
				return "Please input role"
			case "max":
				return "The role is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The role is more than " + err[n].Param() + " characters in length"
			}
		}
	}

	return "Parameter error"
}

// 重置密码，验证错误自定义返回错误文案
func RestPwdGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "LoginPwd" {
			switch err[n].Tag() {
			case "required":
				return "Please input password"
			case "max":
				return "The password is more than " + err[n].Param() + " characters in length"
			case "min":
				return "The password is more than " + err[n].Param() + " characters in length"
			}
		}
	}

	return "Parameter error"
}
