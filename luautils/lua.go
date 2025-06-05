package luautils

import (
	lua "github.com/yuin/gopher-lua"
)

// Processor handles Lua script processing
type Processor struct {
	L *lua.LState
}

// NewProcessor creates a new Lua script processor
func NewProcessor(path string) *Processor {
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

	// register custom modules
	RegisterEncodingModule(L)
	RegisterFSModule(L)

	// Set base path
	SetBasePath(path)

	return &Processor{
		L: L,
	}
}

// Process takes a Lua script as a string and returns values and file paths
func (p *Processor) Process(luaScript string) (map[string]interface{}, []string, error) {
	defer p.L.Close()

	// Initialize empty results
	values := make(map[string]interface{})
	var valuesFiles []string

	// Register global values and valuesFiles in Lua
	valuesTable := p.L.NewTable()
	p.L.SetGlobal("values", valuesTable)

	valuesFilesTable := p.L.NewTable()
	p.L.SetGlobal("valuesFiles", valuesFilesTable)

	// Execute the Lua script
	err := p.L.DoString(luaScript)
	if err != nil {
		return nil, nil, err
	}

	// Extract values from the Lua state
	values = extractValues(p.L, valuesTable)

	// Extract valuesFiles from the Lua state
	valuesFiles = extractValuesFiles(p.L, valuesFilesTable)

	return values, valuesFiles, nil
}

// extractValues extracts values from Lua table and converts them to Go types
func extractValues(L *lua.LState, table *lua.LTable) map[string]interface{} {
	values := make(map[string]interface{})

	L.ForEach(table, func(k, v lua.LValue) {
		// Convert lua.LValue to Go types
		keyStr := k.String()
		values[keyStr] = ToGoValue(v)
	})

	return values
}

// extractValuesFiles extracts valuesFiles from Lua table
func extractValuesFiles(L *lua.LState, table *lua.LTable) []string {
	var valuesFiles []string

	L.ForEach(table, func(k, v lua.LValue) {
		// Only consider string values in valuesFiles
		if lv, ok := v.(lua.LString); ok {
			valuesFiles = append(valuesFiles, string(lv))
		}
	})

	return valuesFiles
}
