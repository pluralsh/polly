package luautils

import (
	lua "github.com/yuin/gopher-lua"
)

// Processor handles Lua script processing
type Processor struct {
	BasePath string
}

func NewProcessor(path string) *Processor {
	return &Processor{BasePath: path}
}

func (p *Processor) NewLuaState() *lua.LState {
	L := lua.NewState(lua.Options{
		SkipOpenLibs: true,
	})

	// Load only safe standard libraries
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
		{lua.StringLibName, lua.OpenString},
		{lua.MathLibName, lua.OpenMath},
	} {
		if err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}

	// Register custom modules
	RegisterEncodingModule(p, L)
	RegisterFSModule(p, L)
	RegisterUtilsModule(L)

	return L
}
