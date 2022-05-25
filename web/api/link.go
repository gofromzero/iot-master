package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"golang.org/x/net/websocket"
)

func linkRoutes(app *gin.RouterGroup) {
	app.POST("list", linkList)
	app.Use(parseParamId)
	app.GET(":id", linkDetail)
	app.POST(":id", linkUpdate)
	app.GET(":id/delete", linkDelete)
	app.GET(":id/close", linkClose)
	app.GET(":id/enable", linkEnable)
	app.GET(":id/disable", linkDisable)
	app.GET(":id/watch", linkWatch)
}

func linkList(ctx *gin.Context) {
	records, cnt, err := normalSearch(ctx, database.Master, &model.Link{})
	if err != nil {
		replyError(ctx, err)
		return
	}


	//补充信息
	links := records.(*[]*model.Link)
	ls := make([]*model.LinkEx, 0) //len(links)

	for _, d := range *links {
		l := &model.LinkEx{Link: *d}
		ls = append(ls, l)
		d := master.GetLink(l.Id)
		if d != nil {
			l.Running = d.Instance.Running()
		}

		var tunnel model.Tunnel
		err := database.Master.One("Id", l.TunnelId, &tunnel)
		if err == nil {
			l.Tunnel = tunnel.Name
		}
	}

	replyList(ctx, ls, cnt)
}



func linkDetail(ctx *gin.Context) {
	var link model.Link
	err := database.Master.One("Id", ctx.GetInt64("id"), &link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	l := &model.LinkEx{Link: link}
	d := master.GetLink(l.Id)
	if d != nil {
		l.Running = d.Instance.Running()
	}

	var tunnel model.Tunnel
	err = database.Master.One("Id", l.TunnelId, &tunnel)
	if err == nil {
		l.Tunnel = tunnel.Name
	}

	replyOk(ctx, l)
}

func linkUpdate(ctx *gin.Context) {
	var link model.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}
	link.Id = ctx.GetInt64("id")

	err = database.Master.Update(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, link)
}

func linkDelete(ctx *gin.Context) {
	link := model.Link{Id: ctx.GetInt64("id")}
	err := database.Master.DeleteStruct(&link)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, link)
	//关闭
	go func() {
		err := master.RemoveLink(link.Id)
		if err != nil {
			log.Error(err)
		}
	}()
}

func linkClose(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt64("id"))
	if link == nil {
		replyFail(ctx, "link not found")
		return
	}
	err := link.Instance.Close()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func linkEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Link{Id: ctx.GetInt64("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func linkDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Link{Id: ctx.GetInt64("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		link := master.GetLink(ctx.GetInt64("id"))
		if link == nil {
			return
		}
		err := link.Instance.Close()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func linkWatch(ctx *gin.Context) {
	link := master.GetLink(ctx.GetInt64("id"))
	if link == nil {
		replyFail(ctx, "找不到链接")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, link.Instance)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}
