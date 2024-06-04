package template

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/osteele/liquid"
	"github.com/pluralsh/polly/algorithms"
)

var (
	liquidEngine           = liquid.NewEngine()
	excludedSprigFunctions = []string{"hello", "now", "uuidv4"}
	sprigFunctionNameMap   = map[string]string{
		"toJson":        "to_json",
		"fromJson":      "from_json",
		"semverCompare": "semver_compare",
		"sha256sum":     "sha26sum",
	}
)

func init() {
	fncs := sprig.TxtFuncMap()

	for name, fnc := range fncs {
		if algorithms.Index(excludedSprigFunctions, func(s string) bool { return s == name }) < 0 {
			liquidEngine.RegisterFilter(name, fnc)
		}
	}

	for key, name := range sprigFunctionNameMap {
		liquidEngine.RegisterFilter(name, fncs[key])
	}

	liquidEngine.RegisterFilter("ternary", ternary)
}

func RenderLiquid(input []byte, bindings map[string]interface{}) ([]byte, error) {
	return liquidEngine.ParseAndRender(input, bindings)
}
