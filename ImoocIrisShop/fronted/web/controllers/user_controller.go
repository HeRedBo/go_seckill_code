package controllers

import (
	"ImoocIrisShop/services"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
	//Session *sessions.Session
}

func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name :"user/register.html",
	}
}

//func (c *UserController) PostRegister() {
func (u *UserController) PostRegister()  {

	fmt.Println(u.Ctx.Request())
	fmt.Println(u.Ctx.FormValue("nickname"))
	return
	//var (
	//	nickname = u.Ctx.FormValue("nickname")
	//	username = u.Ctx.FormValue("username")
	//	password = u.Ctx.FormValue("password")
	//)
	// 数据校验
	//fmt.Println(nickname,username,password)

	//user := &datamodels.User{
	//	//UserName: username,
	//	//NickName: nickname,
	//	//HashPassword: password,
	//}
	//fmt.Println(user)
	//_, err :=u.Service.AddUser(user)
	//u.Ctx.Application().Logger().Debug(err)
	//if err != nil {
	//	u.Ctx.Redirect("/user/error")
	//	return
	//}
	//u.Ctx.Redirect("/user/login")
}

