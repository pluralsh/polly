package luautils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

var basePath string

// SetBasePath sets the base path for file operations
func SetBasePath(path string) {
	basePath = path
}

// RegisterFSModule registers the fs module functions
func RegisterFSModule(L *lua.LState) {
	mod := L.RegisterModule("fs", map[string]lua.LGFunction{
		"read": fsRead,
		"walk": fsWalk,
	})
	L.Push(mod)
}

func fsRead(L *lua.LState) int {
	filePath := L.CheckString(1)

	// Validate and clean the path
	cleanPath, err := validatePath(filePath)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	content, err := os.ReadFile(cleanPath)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(string(content)))
	return 1
}

func fsWalk(L *lua.LState) int {
	dir := L.CheckString(1)

	// Validate and clean the path
	cleanPath, err := validatePath(dir)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	var files []string
	err = filepath.Walk(cleanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Convert absolute path to relative path from base directory
			relPath, err := filepath.Rel(basePath, path)
			if err != nil {
				return err
			}
			files = append(files, relPath)
		}
		return nil
	})

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	// Convert files slice to Lua table
	filesTable := L.NewTable()
	for i, file := range files {
		L.RawSetInt(filesTable, i+1, lua.LString(file))
	}

	L.Push(filesTable)
	return 1
}

func validatePath(path string) (string, error) {
	if basePath == "" {
		return "", fmt.Errorf("base path not set")
	}
	// Clean the path and resolve relative components
	cleanPath := filepath.Clean(filepath.Join(basePath, path))

	// Ensure the path is within the base directory
	if !strings.HasPrefix(cleanPath, basePath) {
		return "", fmt.Errorf("access denied: path outside base directory")
	}

	return cleanPath, nil
}
