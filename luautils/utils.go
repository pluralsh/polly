package luautils

import (
	"fmt"
	"path/filepath"
	"strings"

	"dario.cat/mergo"
	lua "github.com/yuin/gopher-lua"
)

// RegisterUtilsModule registers the utils module functions
func RegisterUtilsModule(L *lua.LState) {
	mod := L.RegisterModule("utils", map[string]lua.LGFunction{
		"merge":       merge,
		"splitString": splitString,
		"pathJoin":    pathJoin,
	})
	L.Push(mod)
}

func pathJoin(L *lua.LState) int {
	parts := L.CheckTable(1)

	converted := ToGoValue(parts).([]interface{})
	res := make([]string, 0, len(converted))
	for _, part := range converted {
		res = append(res, part.(string))
	}

	joined := filepath.Join(res...)
	L.Push(GoValueToLuaValue(L, joined))
	return 1
}

func splitString(L *lua.LState) int {
	str := L.CheckString(1)
	delim := L.CheckString(2)

	parts := strings.Split(str, delim)
	fmt.Printf("parts: %v\n", parts)
	L.Push(GoValueToLuaValue(L, parts))
	return 1
}

func merge(L *lua.LState) int {
	// Get the destination (first argument)
	dst := L.CheckTable(1)
	// Get the source (second argument)
	src := L.CheckTable(2)

	override := L.OptString(3, "override")

	strategy := mergo.WithOverride
	if override == "append" {
		strategy = mergo.WithAppendSlice
	}

	// Convert Lua tables to Go maps
	dstMap := ToGoValue(dst).(map[interface{}]interface{})
	srcMap := ToGoValue(src).(map[interface{}]interface{})

	// Perform deep merge using mergo
	err := mergo.Merge(&dstMap, srcMap, strategy)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	// Convert back to Lua table and return
	L.Push(GoValueToLuaValue(L, sanitizeValue(dstMap)))
	return 1
}
