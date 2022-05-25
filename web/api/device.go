package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-master/database"
	"github.com/zgwit/iot-master/influx"
	"github.com/zgwit/iot-master/log"
	"github.com/zgwit/iot-master/master"
	"github.com/zgwit/iot-master/model"
	"github.com/zgwit/iot-master/tsdb"
	"github.com/zgwit/storm/v3/q"
	"golang.org/x/net/websocket"
	"regexp"
	"strconv"
	"time"
)

func deviceRoutes(app *gin.RouterGroup) {
	app.POST("list", deviceList)
	app.POST("create", deviceCreate)

	app.GET("alarm/clear", deviceAlarmClearAll)

	app.Use(parseParamId)
	app.GET(":id", deviceDetail)
	app.POST(":id", deviceUpdate)
	app.GET(":id/delete", deviceDelete)
	app.GET(":id/start", deviceStart)
	app.GET(":id/stop", deviceStop)
	app.GET(":id/enable", deviceEnable)
	app.GET(":id/disable", deviceDisable)
	app.GET(":id/context", deviceContext)
	app.GET(":id/refresh", deviceRefresh)
	app.GET(":id/refresh/:name", deviceRefreshPoint)
	app.POST(":id/execute", deviceExecute)
	app.GET(":id/watch", deviceWatch)
	app.POST(":id/alarm/list", deviceAlarm)
	app.GET(":id/alarm/clear", deviceAlarmClear)

	app.GET(":id/value/:name/history", deviceValueHistory)
}

func deviceList(ctx *gin.Context) {
	records, cnt, err := normalSearch(ctx, database.Master, &model.Device{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	devices := records.(*[]*model.Device)
	devs := make([]*model.DeviceEx, 0) //len(devices)

	for _, d := range *devices {
		dev := &model.DeviceEx{Device: *d}
		devs = append(devs, dev)
		d := master.GetDevice(dev.Id)
		if d != nil {
			dev.Running = d.Running()
		}
		if dev.ElementId != "" {
			var element model.Element
			err := database.Master.One("Id", dev.ElementId, &element)
			if err == nil {
				dev.Element = element.Name
				dev.DeviceContent = element.DeviceContent
			}
		}
		var link model.Link
		err := database.Master.One("Id", dev.LinkId, &link)
		if err == nil {
			if link.Name != "" {
				dev.Link = link.Name
			} else if link.SN != "" {
				dev.Link = link.SN
			} else if link.Remote != "" {
				dev.Link = link.Remote
			}
		}
	}

	replyList(ctx, devs, cnt)
}

func deviceCreate(ctx *gin.Context) {
	var device model.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	err = database.Master.Save(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, device)

	//启动
	if !device.Disabled {
		go func() {
			_, err := master.LoadDevice(device.Id)
			if err != nil {
				log.Error(err)
				return
			}
		}()
	}
}

func deviceDetail(ctx *gin.Context) {
	var device model.Device
	err := database.Master.One("Id", ctx.GetInt64("id"), &device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	//补充信息
	dev := model.DeviceEx{Device: device}
	d := master.GetDevice(dev.Id)
	if d != nil {
		dev.Running = d.Running()
	}
	if dev.ElementId != "" {
		var element model.Element
		err := database.Master.One("Id", dev.ElementId, &element)
		if err == nil {
			dev.Element = element.Name
			dev.DeviceContent = element.DeviceContent
		}
		var link model.Link
		err = database.Master.One("Id", dev.LinkId, &link)
		if err == nil {
			if link.Name != "" {
				dev.Link = link.Name
			} else if link.SN != "" {
				dev.Link = link.SN
			} else if link.Remote != "" {
				dev.Link = link.Remote
			}
		}
	}

	replyOk(ctx, dev)
}

func deviceUpdate(ctx *gin.Context) {
	var device model.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}
	device.Id = ctx.GetInt64("id")

	err = database.Master.Update(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, device)

	//重新启动
	go func() {
		_ = master.RemoveDevice(device.Id)
		_, err = master.LoadDevice(device.Id)
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func deviceDelete(ctx *gin.Context) {
	device := model.Device{Id: ctx.GetInt64("id")}
	err := database.Master.DeleteStruct(&device)
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, device)

	//关闭
	go func() {
		err := master.RemoveDevice(device.Id)
		if err != nil {
			log.Error(err)
		}
	}()
}

func deviceStart(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Start()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceStop(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "not found")
		return
	}
	err := device.Stop()
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceEnable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Device{Id: ctx.GetInt64("id")}, "Disabled", false)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//启动
	go func() {
		_, err := master.LoadDevice(ctx.GetInt64("id"))
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func deviceDisable(ctx *gin.Context) {
	err := database.Master.UpdateField(&model.Device{Id: ctx.GetInt64("id")}, "Disabled", true)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)

	//关闭
	go func() {
		device := master.GetDevice(ctx.GetInt64("id"))
		if device == nil {
			return
		}
		err := device.Stop()
		if err != nil {
			log.Error(err)
			return
		}
	}()
}

func deviceContext(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	replyOk(ctx, device.Context)
}

func deviceRefresh(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err := device.Refresh()
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func deviceRefreshPoint(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	val, err := device.RefreshPoint(ctx.Param("name"))
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, val)
}

type executeBody struct {
	Command   string    `json:"command"`
	Arguments []float64 `json:"arguments"`
}

func deviceExecute(ctx *gin.Context) {
	var body executeBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		replyError(ctx, err)
		return
	}

	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	err = device.Execute(body.Command, body.Arguments)
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyOk(ctx, nil)
}

func deviceWatch(ctx *gin.Context) {
	device := master.GetDevice(ctx.GetInt64("id"))
	if device == nil {
		replyFail(ctx, "找不到设备")
		return
	}
	websocket.Handler(func(ws *websocket.Conn) {
		watchAllEvents(ws, device)
	}).ServeHTTP(ctx.Writer, ctx.Request)
}

func deviceAlarm(ctx *gin.Context) {
	alarms, cnt, err := normalSearchById(ctx, database.History, "DeviceId", ctx.GetInt64("id"), &model.DeviceAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}
	replyList(ctx, alarms, cnt)
}

func deviceAlarmClear(ctx *gin.Context) {
	err := database.History.Select(q.Eq("DeviceId", ctx.GetInt64("id"))).Delete(&model.DeviceAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

func deviceAlarmClearAll(ctx *gin.Context) {
	err := database.History.Drop(&model.DeviceAlarm{})
	if err != nil {
		replyError(ctx, err)
		return
	}

	replyOk(ctx, nil)
}

var timeReg *regexp.Regexp

func init() {
	timeReg = regexp.MustCompile(`^(-?\d+)(h|m|s)$`)
}

func parseTime(tm string) (int64, error) {
	ss := timeReg.FindStringSubmatch(tm)
	if ss == nil || len(ss) != 3 {
		return 0, errors.New("错误时间")
	}
	val, _ := strconv.ParseInt(ss[1], 10, 64)
	switch ss[2] {
	case "d":
		val *= 24 * 60 * 60 * 1000
	case "h":
		val *= 60 * 60 * 1000
	case "m":
		val *= 60 * 1000
	case "s":
		val *= 1000
	}
	return val, nil
}

func deviceValueHistory(ctx *gin.Context) {
	id := ctx.Param("id")
	key := ctx.Param("name")
	start := ctx.DefaultQuery("start", "-5h")
	end := ctx.DefaultQuery("end", "0h")
	window := ctx.DefaultQuery("window", "10m")

	//优先查询InfluxDB
	if influx.Opened() {
		values, err := influx.Query(map[string]string{"id": id}, key, start, end, window)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, values)
		return
	}

	//查询内部数据库
	if tsdb.Opened() {
		//相对时间转化为时间戳
		s, err := parseTime(start)
		if err != nil {
			replyError(ctx, err)
			return
		}
		s += time.Now().UnixMilli()

		e, err := parseTime(end)
		if err != nil {
			replyError(ctx, err)
			return
		}
		e += time.Now().UnixMilli()

		w, err := parseTime(window)
		if err != nil {
			replyError(ctx, err)
			return
		}
		values, err := tsdb.Query(id, key, s, e, w)
		if err != nil {
			replyError(ctx, err)
			return
		}
		replyOk(ctx, values)
		return
	}

	replyFail(ctx, "没有开启历史数据库")
}
