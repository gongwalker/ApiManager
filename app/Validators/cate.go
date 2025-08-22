package Validators

import (
	"github.com/go-playground/validator/v10"
)

// 添加分类，验证错误自定义返回错误文案
func AddCateGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "Cname" {
			switch err[n].Tag() {
			case "required":
				return "Please input collection"
			case "max":
				return "The description is more than " + err[n].Param() + " characters in length!"
			}

		}
		if err[n].Field() == "Cdesc" {
			switch err[n].Tag() {
			case "required":
				return "Please input description!"
			case "max":
				return "The description is more than " + err[n].Param() + " characters in length!"
			}
		}
	}

	return "Parameter error"
}

// 添加分类，验证错误自定义返回错误文案
func EditCateGetError(err validator.ValidationErrors) string {
	for n := range err {
		if err[n].Field() == "Aid" {
			switch err[n].Tag() {
			case "required":
				return "Aid is a required parameter"
			}
		}

		if err[n].Field() == "Cname" {
			switch err[n].Tag() {
			case "required":
				return "Please input collection"
			case "max":
				return "The description is more than " + err[n].Param() + " characters in length"
			}

		}
		if err[n].Field() == "Cdesc" {
			switch err[n].Tag() {
			case "required":
				return "Please input description"
			case "max":
				return "The description is more than " + err[n].Param() + " characters in length"
			}
		}

		if err[n].Field() == "Csort" {
			switch err[n].Tag() {
			case "gte":
				fallthrough
			case "lte":
				return "Sort value should be between 0 and 9999"
			case "number":
				return "Sort value should be is number"
			}

		}
	}

	return "Parameter error"
}
