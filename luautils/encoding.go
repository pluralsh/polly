package luautils

import (
	"fmt"

	"dario.cat/mergo"
	lua "github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/util/json"
)

// RegisterEncodingModule registers the encoding module functions
func RegisterEncodingModule(L *lua.LState) {
	mod := L.RegisterModule("encoding", map[string]lua.LGFunction{
		"jsonEncode": jsonEncode,
		"jsonDecode": jsonDecode,
		"yamlEncode": yamlEncode,
		"yamlDecode": yamlDecode,
		"merge":      merge,
	})
	L.Push(mod)
}

func jsonEncode(L *lua.LState) int {
	value := L.CheckAny(1)
	goValue := ToGoValue(value)

	sanitized := sanitizeValue(goValue)

	jsonBytes, err := json.Marshal(sanitized)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(jsonBytes))
	return 1
}

func jsonDecode(L *lua.LState) int {
	jsonStr := L.CheckString(1)

	var goValue interface{}
	err := json.Unmarshal([]byte(jsonStr), &goValue)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(GoValueToLuaValue(L, goValue))
	return 1
}

func yamlEncode(L *lua.LState) int {
	value := L.CheckAny(1)
	goValue := ToGoValue(value)

	yamlBytes, err := yaml.Marshal(goValue)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(string(yamlBytes)))
	return 1
}

func yamlDecode(L *lua.LState) int {
	yamlStr := L.CheckString(1)

	var goValue interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &goValue)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(GoValueToLuaValue(L, goValue))
	return 1
}

func sanitizeValue(val interface{}) interface{} {
	switch v := val.(type) {
	case map[interface{}]interface{}:
		m := make(map[string]interface{})
		for key, value := range v {
			strKey := fmt.Sprintf("%v", key) // Convert key to string
			m[strKey] = sanitizeValue(value)
		}
		return m
	case []interface{}:
		for i := range v {
			v[i] = sanitizeValue(v[i])
		}
		return v
	default:
		return v
	}
}

func merge(L *lua.LState) int {
	// Get the destination (first argument)
	dst := L.CheckTable(1)
	// Get the source (second argument)
	src := L.CheckTable(2)

	// Convert Lua tables to Go maps
	dstMap := ToGoValue(dst).(map[interface{}]interface{})
	srcMap := ToGoValue(src).(map[interface{}]interface{})

	// Perform deep merge using mergo
	err := mergo.Merge(&dstMap, srcMap, mergo.WithOverride)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	// Convert back to Lua table and return
	L.Push(GoValueToLuaValue(L, sanitizeValue(dstMap)))
	return 1
}
