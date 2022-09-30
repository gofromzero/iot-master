package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/v2/internal/core"
	"net"
	"net/http"
	"net/url"
)

func ServiceProxy(ctx *gin.Context) {
	svc := core.Services.Load(ctx.Param("name"))
	if svc == nil {
		replyFail(ctx, "not found")
		return
	}

	req := ctx.Request.Clone(ctx)
	req.URL, _ = url.Parse(svc.Addr + req.RequestURI)
	req.RequestURI = ""

	//req, _ := http.NewRequest(ctx.Request.Method, ctx.Request.RequestURI, ctx.Request.Body)
	cli := http.Client{}
	if svc.Net == "unix" {
		cli.Transport = &http.Transport{DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		}}
	}

	resp, err := cli.Do(req)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//返回结果
	//ctx.Request.Response = resp
	_ = resp.Write(ctx.Writer)
	ctx.Abort()
}
