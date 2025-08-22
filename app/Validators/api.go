package Validators

import (
	"github.com/go-playground/validator/v10"
)

// 添加 or 编辑Api，验证错误自定义返回错误文案
func ApiGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "ParamName" {
			switch err[n].Tag() {
			case "param_name":
				return "Request parameter name length should be between 1 and 100"
			}
		}

		if err[n].Field() == "ParamType" {
			switch err[n].Tag() {
			case "param_type":
				return "Request parameter type length should be between 1 and 20"
			}
		}

		if err[n].Field() == "Num" {
			return "Api No length should be between 1 and 50"
		}

		if err[n].Field() == "Name" {
			return "Api name length should be between 1 and 200"
		}

		if err[n].Field() == "Url" {
			return "Api url length should be between 1 and 200"
		}

		if err[n].Field() == "Des" {
			return "Api desc length should be between 1 and 200"
		}

		if err[n].Field() == "Re" {
			return "Api response length should be between 1 and 100000"
		}

		if err[n].Field() == "St" {
			return "Api Request Note length should be between 1 and 100000"
		}

		if err[n].Field() == "Memo" {
			return "Api Memo length should be between 1 and 100000"
		}

	}
	return "Parameter error"
}

func ParamName(fl validator.FieldLevel) bool {
	if paraname, ok := fl.Field().Interface().([]string); ok {
		for _, v := range paraname {
			if v == "" || len(v) > 100 {
				return false
			}
		}
	}
	return true
}

func ParamType(fl validator.FieldLevel) bool {
	if paratype, ok := fl.Field().Interface().([]string); ok {
		for _, v := range paratype {
			if v == "" || len(v) > 20 {
				return false
			}
		}
	}
	return true
}
