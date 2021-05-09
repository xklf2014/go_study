package service

import (
	"gin_demo/study31_login_project/model"
	"gin_demo/study31_login_project/serializer"
)

type UserRegisterService struct {
	NickName string `json:"nick_name" db:"nick_name"`
	Password string `json:"password"`
	Age      uint32 `json:"age"`
	Gender   uint32 `json:"gender"`
	Email    string `json:"email"`
}

func (service *UserRegisterService) Register() serializer.Response {
	//验证邮箱是否已经存在
	sqlStr := `select count(id) from user where email = ?`
	var count int
	_ = model.Db.Get(&count, sqlStr, service.Email)
	if count > 0 {
		return serializer.Response{
			Code:  40001,
			Data:  nil,
			Msg:   "邮箱已经注册",
			Error: "",
		}
	}

	insertSql := `insert into user (nickname,password,age,gender,email) values(:nick_name,:password,:age,:gender,:email)`
	_, err := model.Db.NamedExec(insertSql, service)
	if err != nil {
		return serializer.Response{
			Code:  40002,
			Data:  nil,
			Msg:   "注册失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "注册成功",
		Error: "",
	}
}
