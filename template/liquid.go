package template

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/osteele/liquid"
	"github.com/samber/lo"
)

type FilterFunction struct {
	Name            string   `json:"name"`
	Aliases         []string `json:"aliases,omitempty"`
	FunctionPath    string   `json:"functionPath,omitempty"`
	FunctionPackage string   `json:"functionPackage,omitempty"`
	FunctionName    string   `json:"functionName,omitempty"`
}

var (
	liquidEngine = liquid.NewEngine()

	// excludedSprigFunctions contains names of Spring functions that will be excluded.
	excludedSprigFunctions = []string{"hello", "now", "uuidv4"}

	// sprigFunctionNameAliases contains additional aliases for Sprig functions.
	sprigFunctionNameAliases = map[string][]string{
		"toJson":        {"to_json"},
		"fromJson":      {"from_json"},
		"semverCompare": {"semver_compare"},
		"sha256sum":     {"sha26sum"},
	}

	// internalFunctions to register. These will override Sprig functions if same names are used.
	internalFunctions = map[string]any{
		"indent":  indent,
		"nindent": nindent,
		"replace": strings.ReplaceAll,
		"default": dfault,
		"ternary": ternary,
	}

	// registeredFunctions contains information about all registered template functions.
	registeredFunctions = map[string]FilterFunction{}
)

func init() {
	sprigFunctions := sprig.TxtFuncMap()
	for name, fnc := range sprigFunctions {
		_, hasInternalFunctionNameConflict := internalFunctions[name]
		if !lo.Contains(excludedSprigFunctions, name) && !hasInternalFunctionNameConflict {
			registerFilter(name, sprigFunctionNameAliases[name], fnc)
		}
	}

	for name, fnc := range internalFunctions {
		registerFilter(name, nil, fnc)
	}
}

func registerFilter(name string, aliases []string, fn any) {
	liquidEngine.RegisterFilter(name, fn)

	for _, alias := range aliases {
		liquidEngine.RegisterFilter(alias, fn)
	}

	fnPath := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	lastDot := strings.LastIndexByte(fnPath, '.')
	fnPackage := ""
	fnName := ""
	if lastDot >= 0 {
		fnPackage = fnPath[:lastDot]
		fnName = fnPath[lastDot+1:]
	}

	registeredFunctions[name] = FilterFunction{
		Name:            name,
		Aliases:         aliases,
		FunctionPath:    fnPath,
		FunctionPackage: fnPackage,
		FunctionName:    fnName,
	}
}

func RegisteredFilters() map[string]FilterFunction {
	return registeredFunctions
}

func RenderLiquid(input []byte, bindings map[string]interface{}) ([]byte, error) {
	return liquidEngine.ParseAndRender(input, bindings)
}
