package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pluralsh/polly/luautils"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
)

type TestProcessor struct {
	*luautils.Processor
}

func NewTestProcessor(filePath string) *TestProcessor {
	p := luautils.NewProcessor(filePath)
	return &TestProcessor{
		Processor: p,
	}
}

// Process takes a Lua script as a string and returns values and file paths
func (p *TestProcessor) Process(luaScript string) (map[string]interface{}, []string, error) {
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

	if err := luautils.MapLua(p.L.GetGlobal("values").(*lua.LTable), &values); err != nil {
		return nil, nil, err
	}

	if err := luautils.MapLua(p.L.GetGlobal("valuesFiles").(*lua.LTable), &valuesFiles); err != nil {
		return nil, nil, err
	}

	return values, valuesFiles, nil
}

func TestGenerateOutput(t *testing.T) {
	// Test Lua script
	luaScript := `
		values = {}
		values["key1"] = "value1"
		values["key2"] = 42
		
		valuesFiles = {"config.json", "data.txt"}
	`

	// Process the Lua script
	p := NewTestProcessor("../files")
	values, valuesFiles, err := p.Process(luaScript)
	assert.NoError(t, err)

	// Check values

	assert.NotNil(t, valuesFiles)
	assert.NotNil(t, values)

	assert.Equal(t, values["key1"], `value1`)
	assert.Equal(t, values["key2"], float64(42))

	assert.Equal(t, len(valuesFiles), 2)
	assert.Equal(t, valuesFiles[0], `config.json`)
	assert.Equal(t, valuesFiles[1], `data.txt`)
}

func TestComplex(t *testing.T) {
	// Test Lua script
	luaScript := `
		local jsonStr = encoding.jsonEncode(fs.read("simple.json"))
		local data = encoding.jsonDecode(jsonStr)
		
		local yamlStr = encoding.yamlEncode({
		  user = {
			name = "Alice",
			roles = {"admin", "user"}
		  }
		})
		local yamlData = encoding.yamlDecode(yamlStr)
		
		-- Define values
		values = {}
		values["name"] = "John Doe"
		values["age"] = 30
		values["isActive"] = true
		values["encoded"] = {
		  json = jsonStr,
		  yaml = yamlStr
		}
		
		-- Define an array
		values["tags"] = {"personal", "important", "urgent"}
		
		-- Define a nested table
		values["settings"] = {
			theme = "dark",
			notifications = true,
			display = {
				fontSize = 14,
				colorScheme = "monokai"
			}
		}
		
		local textFile = fs.read("text.txt")
 		values["text"] = textFile

		-- Define values files
		valuesFiles = {}
		local files = fs.walk(".")
		for i, file in ipairs(files) do
   	 		table.insert(valuesFiles, file)
		end
	`

	// Process the Lua script
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fullPath := filepath.Join(dir, "files")

	p := NewTestProcessor(fullPath)
	values, valuesFiles, err := p.Process(luaScript)
	assert.NoError(t, err)

	assert.NotNil(t, valuesFiles)
	assert.NotNil(t, values)

	assert.Equal(t, values["name"], `John Doe`)
	assert.Equal(t, values["text"], `hello`)
	assert.Equal(t, len(valuesFiles), 2)
}

func TestUnsafeOSLib(t *testing.T) {
	// Test Lua script
	luaScript := `
		values = {}
		valuesFiles = {}

		local files = fs.walk(".")
		for i, file in ipairs(files) do
			os.execute("rm -f " .. filename)
		end

	`

	// Process the Lua script
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fullPath := filepath.Join(dir, "files")

	p := NewTestProcessor(fullPath)
	_, _, err = p.Process(luaScript)

	// Check values
	assert.Error(t, err)

}

func TestUnsafeReadFile(t *testing.T) {
	// Test Lua script
	luaScript := `
		values = {}
		valuesFiles = {}

		local filename = "text.txt"
		
		-- Open the file for reading
		local file = io.open(filename, "r")
		
		if file then
			-- Read the entire contents
			local content = file:read("*all")
			file:close()
		
			print("File contents:")
			print(content)
		else
			print("Failed to open file: " .. filename)
		end

	`

	// Process the Lua script
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fullPath := filepath.Join(dir, "files")

	p := NewTestProcessor(fullPath)
	_, _, err = p.Process(luaScript)

	// Check values
	assert.Error(t, err)

}

func TestFileOutsideTheBaseDir(t *testing.T) {
	// Test Lua script
	luaScript := `
		values = {}
		valuesFiles = {}
		local content, err = fs.walk("../")
		if not content then
			values["error"] = err
		else
			values["content"] = content
		end
		
	`

	// Process the Lua script
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fullPath := filepath.Join(dir, "files")

	p := NewTestProcessor(fullPath)
	values, _, err := p.Process(luaScript)

	// Check values
	assert.NoError(t, err)
	assert.NotNil(t, values)
	assert.Equal(t, map[string]interface{}{"access denied": "path outside base directory"}, values["error"])
}

func TestMerge(t *testing.T) {
	type SSLConfig struct {
		Enabled bool   `json:"enabled"`
		Cert    string `json:"cert"`
		Key     string `json:"key"`
	}

	type ServerConfig struct {
		Host string    `json:"host"`
		Port int       `json:"port"`
		SSL  SSLConfig `json:"ssl"`
	}

	// Test Lua script
	luaScript := `
		values = {}	
		valuesFiles = {}
		local baseConfig = {
			server = {
				host = "localhost",
				port = 8080,
				ssl = {enabled = false, cert = "default.crt"}
			}
		}
		
		local prodOverrides = {
			server = {
				host = "0.0.0.0",
				ssl = {enabled = true, key = "prod.key"}
			}
		}
		
		local finalConfig = encoding.merge(baseConfig, prodOverrides)
		values["config"] = finalConfig

		-- Result: {
		--   server = {
		--     host = "0.0.0.0",        -- overridden
		--     port = 8080,             -- preserved from base
		--     ssl = {
		--       enabled = true,        -- overridden
		--       cert = "default.crt",  -- preserved from base
		--       key = "prod.key"       -- added from override
		--     }
		--   }
		-- }
	`
	// Process the Lua script
	p := NewTestProcessor("../files")
	values, _, err := p.Process(luaScript)
	assert.NoError(t, err)

	// Check values
	assert.NotNil(t, values)
	var config ServerConfig
	rawConfig, ok := values["config"].(map[interface{}]interface{})
	assert.True(t, ok)
	serverRaw, ok := rawConfig["server"].(map[interface{}]interface{})
	assert.True(t, ok)

	err = mapstructure.Decode(serverRaw, &config)
	assert.NoError(t, err)

	assert.Equal(t, "0.0.0.0", config.Host)
	assert.Equal(t, 8080, config.Port)
	assert.True(t, config.SSL.Enabled)
	assert.Equal(t, "default.crt", config.SSL.Cert)
	assert.Equal(t, "prod.key", config.SSL.Key)
}
