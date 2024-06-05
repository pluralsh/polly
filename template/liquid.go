package template

import (
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/osteele/liquid"
	"github.com/pluralsh/polly/algorithms"
)

var (
	liquidEngine = liquid.NewEngine()

	// excludedSprigFunctions contains names of Spring functions that will be excluded.
	excludedSprigFunctions = []string{"hello", "now", "uuidv4"}

	// sprigFunctionNameAliases contains additional aliases for Sprig functions.
	sprigFunctionNameAliases = map[string]string{
		"toJson":        "to_json",
		"fromJson":      "from_json",
		"semverCompare": "semver_compare",
		"sha256sum":     "sha26sum",
	}

	// internalFunctions to register. These will override Sprig functions.
	internalFunctions = map[string]any{
		"indent":  indent,
		"nindent": nindent,
		"replace": strings.ReplaceAll,
		"default": dfault,
		"ternary": ternary,
	}
)

func init() {
	fncs := sprig.TxtFuncMap()

	excludedFunctions := excludedSprigFunctions
	for k := range internalFunctions {
		excludedFunctions = append(excludedFunctions, k)
	}

	for name, fnc := range fncs {
		if algorithms.Index(excludedFunctions, func(s string) bool { return s == name }) < 0 {
			liquidEngine.RegisterFilter(name, fnc)
		}
	}

	for key, name := range sprigFunctionNameAliases {
		liquidEngine.RegisterFilter(name, fncs[key])
	}

	for name, fnc := range internalFunctions {
		liquidEngine.RegisterFilter(name, fnc)
	}
}

func RenderLiquid(input []byte, bindings map[string]interface{}) ([]byte, error) {
	return liquidEngine.ParseAndRender(input, bindings)
}
