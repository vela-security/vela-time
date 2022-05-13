package vtime

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func WithEnv(env assert.Environment) {
	xEnv = env
	uv := lua.NewUserKV()
	uv.Set("sec", lua.NewFunction(newLuaSecond))
	uv.Set("min", lua.NewFunction(newLuaMinute))
	uv.Set("hour", lua.NewFunction(newLuaHour))
	uv.Set("day", lua.NewFunction(newLuaDay))
	uv.Set("week", lua.NewFunction(newLuaWeek))
	uv.Set("month", lua.NewFunction(newLuaMonth))
	uv.Set("year", lua.NewFunction(newLuaYear))
	uv.Set("sleep", lua.NewFunction(newLuaTimeSleep))
	uv.Set("today", lua.NewFunction(newLuaTimeToday))
	uv.Set("now", lua.NewFunction(newLuaTimeNow))
	uv.Set("at", lua.NewFunction(newluaTimeAt))
	xEnv.Global("time", uv)
}
