package service

import (
	"gin_demo/study31_login_project/model"
	"gin_demo/study31_login_project/serializer"
)

type UserLoginService struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (service *UserLoginService) Login() serializer.Response {
	loginSql := `select count(id) from user where email = ? and password = ?`
	var count int
	_ = model.Db.Get(&count, loginSql, service.UserName, service.Password)

	if count == 0 {
		return serializer.Response{
			Code: 40003,
			Msg:  "账号或者密码错误",
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "登录成功",
	}
}
