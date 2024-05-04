package template

import (
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/osteele/liquid"
)

var (
	liquidEngine   = liquid.NewEngine()
	sprigFunctions = map[string]string{
		"toJson":        "to_json",
		"fromJson":      "from_json",
		"b64enc":        "b64enc",
		"b64dec":        "b64dec",
		"semverCompare": "semver_compare",
		"sha256sum":     "sha26sum",
		"quote":         "quote",
		"squote":        "squote",
		"replace":       "replace",
		"coalesce":      "coalesce",
	}
)

func init() {
	fncs := sprig.TxtFuncMap()
	for key, name := range sprigFunctions {
		liquidEngine.RegisterFilter(name, fncs[key])
	}
	liquidEngine.RegisterFilter("indent", indent)
	liquidEngine.RegisterFilter("nindent", nindent)
	liquidEngine.RegisterFilter("replace", strings.ReplaceAll)

	liquidEngine.RegisterFilter("default", dfault)
	liquidEngine.RegisterFilter("ternary", ternary)
}

func RenderLiquid(input []byte, bindings map[string]interface{}) ([]byte, error) {
	return liquidEngine.ParseAndRender(input, bindings)
}
