package events

import "sync"

type Handler func(args ...interface{})

//EventInterface Events接口
type EventInterface interface {
	Emit(event string, data ...interface{})
	On(event string, fn interface{})
	Once(event string, fn interface{})
	Off(event string, fn interface{})
}

type EventEmitter struct {
	handlers sync.Map
}

//Emit 发送消息
func (e *EventEmitter) Emit(event string, data ...interface{}) {
	val, ok := e.handlers.Load(event)
	if !ok {
		return
	}
	handlers := val.(*sync.Map)
	handlers.Range(func(key, value interface{}) bool {
		callback := key.(Handler)
		callback(data...)
		//处理仅订阅一次
		once := value.(bool)
		if once {
			handlers.Delete(key)
		}
		return true
	})
}

//On 监听
func (e *EventEmitter) On(event string, fn interface{}) {
	val, ok := e.handlers.Load(event)
	if !ok {
		val = new(sync.Map)
		e.handlers.Store(event, val)
	}
	handlers := val.(*sync.Map)
	handlers.Store(fn, false)
}

//Once 监听一次
func (e *EventEmitter) Once(event string, fn interface{}) {
	val, ok := e.handlers.Load(event)
	if !ok {
		val = new(sync.Map)
		e.handlers.Store(event, val)
	}
	handlers := val.(*sync.Map)
	handlers.Store(fn, true)
}

//Off 取消监听
func (e *EventEmitter) Off(event string, fn interface{}) {
	val, ok := e.handlers.Load(event)
	if !ok {
		return
	}
	handlers := val.(*sync.Map)
	handlers.Delete(fn)
}
