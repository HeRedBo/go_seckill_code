package controllers

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/datamodels"
	"ImoocIrisShop/services"
	"ImoocIrisShop/tool"
	"fmt"
	"log"
	"strconv"

	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (u *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (u *UserController) PostRegister() {

	var (
		nickname = u.Ctx.PostValue("nickname")
		username = u.Ctx.PostValue("username")
		password = u.Ctx.PostValue("password")
	)
	// 数据校验
	user := &datamodels.User{
		Username: username,
		Nickname: nickname,
		Password: password,
	}
	_, err := u.Service.AddUser(user)
	u.Ctx.Application().Logger().Debug(err)
	if err != nil {
		u.Ctx.Redirect("/user/error")
		return
	}
	u.Ctx.Redirect("/user/login")
}

func (u *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (u *UserController) PostLogin() mvc.Response {
	//1.获取用户提交的表单信息
	var (
		userName = u.Ctx.FormValue("userName")
		password = u.Ctx.FormValue("password")
	)

	fmt.Println(userName, password)
	//2、验证账号密码正确
	user, isOk := u.Service.IsPwdSuccess(userName, password)
	fmt.Println(user, "isOk", isOk)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	// 将用户ID 写入到 Cookie中
	tool.GlobalCookie(u.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	u.Session.Set("userID", strconv.FormatInt(user.ID, 10))
	log.Println("login success")

	// 跳转到 产品首页
	return mvc.Response{
		Path: "/product/",
	}
}

func (u *UserController) PostTest() {
	u.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	fmt.Println("asdfasd", dec)

	//return fmt.Sprintf("%s is %d years old\n", form.Username, form.Age)
}

func (u *UserController) HandleError(ctx iris.Context, err error) {
	if iris.IsErrPath(err) {
		fmt.Println(ctx.Request().RequestURI)
		return // continue.
	}
	ctx.StopWithError(iris.StatusBadRequest, err)
}
