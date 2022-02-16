package middleware

import "github.com/kataras/iris/v12"

func AuthCheck(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("请先登录")
		ctx.Redirect("/user/login")
	}
	ctx.Application().Logger().Debug("已经登陆")
	ctx.Next()
}
