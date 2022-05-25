package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/model"
)

func userRoutes(app *gin.RouterGroup) {

	app.GET("me", userMe)


	app.POST("list", userList)
	app.POST("create", userCreate)

	app.Use(parseParamId)

	app.GET(":id", userDetail)
	app.POST(":id", userUpdate)
	app.GET(":id/delete", userDelete)
	app.GET(":id/password", userPassword)
	app.GET(":id/enable", userEnable)
	app.GET(":id/disable", userDisable)

}

func userMe(ctx *gin.Context) {
	replyOk(ctx, ctx.MustGet("user"))
}

func userList(ctx *gin.Context) {
	users, cnt, err := normalSearch(ctx, database.Master, &model.User{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, users, cnt)
}

func userCreate(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//TODO 默认密码

	replyOk(ctx, user)
}

func userDetail(ctx *gin.Context) {
	var user model.User
	err := database.Master.One("Id", ctx.GetInt64("id"), &user)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, user)
}

func userUpdate(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}
	user.Id = ctx.GetInt64("id")

	err = database.Master.Update(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, user)
}

func userDelete(ctx *gin.Context) {
	user := model.User{Id: ctx.GetInt64("id")}
	err := database.Master.DeleteStruct(&user)
	if err != nil {
		replyError(ctx, err)
		return
	}

	_ = database.Master.DeleteStruct(model.Password{Id: user.Id})

	replyOk(ctx, user)
}

func userPassword(ctx *gin.Context) {

	replyOk(ctx, nil)
}

func userEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(model.User{Id: ctx.GetInt64("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func userDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(model.User{Id: ctx.GetInt64("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}
