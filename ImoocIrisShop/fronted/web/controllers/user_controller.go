package controllers

import (
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name :"user/register.html",
	}
}


func (u *UserController) PostRegister()  {
	var (
		nickname = u.Ctx.FormValue("nickname")
		username = u.Ctx.FormValue("username")
		password = u.Ctx.FormValue("password")
	)
	// 数据校验
	user := &datamodels.User{
		UserName: username,
		NickName: nickname,
		HashPassword: password,
	}
	_, err :=u.Service.AddUser(user)
	u.Ctx.Application().Logger().Debug(err)
	if err != nil {
		u.Ctx.Redirect("/user/error")
		return
	}
	u.Ctx.Redirect("/user/login")
	return
}

