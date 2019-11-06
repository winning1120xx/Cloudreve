package controllers

import (
	"Cloudreve/serializer"
	"encoding/json"
	"gopkg.in/go-playground/validator.v8"
)

// ParamErrorMsg 根据Validator返回的错误信息给出错误提示
func ParamErrorMsg(filed string, tag string) string {
	// 未通过验证的表单域与中文对应
	fieldMap := map[string]string{
		"UserName": "用户名",
		"Password": "密码",
	}
	// 未通过的规则与中文对应
	tagMap := map[string]string{
		"required": "不能为空",
		"min":      "太短",
		"max":      "太长",
	}
	fieldVal, findField := fieldMap[filed]
	tagVal, findTag := tagMap[tag]
	if findField && findTag {
		// 返回拼接出来的错误信息
		return fieldVal + tagVal
	}
	return ""
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	// 处理 Validator 产生的错误
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.ParamErr(
				ParamErrorMsg(e.Field, e.Tag),
				err,
			)
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}