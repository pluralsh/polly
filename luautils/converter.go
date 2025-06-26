package luautils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	lua "github.com/yuin/gopher-lua"
)

// Mapper maps a lua table to a Go struct pointer.
type Mapper struct {
}

// MapLua maps the lua table to the given struct pointer with default options.
func MapLua(tbl *lua.LTable, st interface{}) error {
	return NewMapper().Map(tbl, st)
}

// NewMapper returns a new mapper.
func NewMapper() *Mapper {

	return &Mapper{}
}

// Map maps the lua table to the given struct pointer.
func (mapper *Mapper) Map(tbl *lua.LTable, st interface{}) error {
	goValue := ToGoValue(tbl)

	stVal := reflect.ValueOf(st)
	if stVal.Kind() != reflect.Ptr || stVal.IsNil() {
		return errors.New("st must be a non-nil pointer")
	}

	stElem := stVal.Elem()
	stKind := stElem.Kind()

	var config = &mapstructure.DecoderConfig{
		Result: st,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	switch v := goValue.(type) {
	case map[interface{}]interface{}:
		if stKind != reflect.Struct && stKind != reflect.Map {
			stElem.Set(reflect.Zero(stElem.Type()))
			return nil
		}
		return decoder.Decode(v)
	case []interface{}:
		if stKind != reflect.Slice && stKind != reflect.Array {
			stElem.Set(reflect.Zero(stElem.Type()))
			return nil
		}
		return decoder.Decode(v)
	default:
		return errors.New("unsupported table format: expected map or array")
	}
}

// ToGoValue converts the given LValue to a Go object.
func ToGoValue(lv lua.LValue) interface{} {
	switch v := lv.(type) {
	case *lua.LNilType:
		return nil
	case lua.LBool:
		return bool(v)
	case lua.LString:
		return string(v)
	case lua.LNumber:
		return float64(v)
	case *lua.LTable:
		maxn := v.MaxN()
		if maxn == 0 { // table
			ret := make(map[interface{}]interface{})
			v.ForEach(func(key, value lua.LValue) {
				keystr := fmt.Sprint(ToGoValue(key))
				ret[keystr] = ToGoValue(value)
			})
			return ret
		} else { // array
			ret := make([]interface{}, 0, maxn)
			for i := 1; i <= maxn; i++ {
				ret = append(ret, ToGoValue(v.RawGetInt(i)))
			}
			return ret
		}
	default:
		return v
	}
}

// GoValueToLuaValue converts a Go value to a Lua value
func GoValueToLuaValue(L *lua.LState, value interface{}) lua.LValue {
	if value == nil {
		return lua.LNil
	}

	rType := reflect.TypeOf(value)

	if rType.Kind() == reflect.Slice || rType.Kind() == reflect.Array {
		table := L.NewTable()
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			L.RawSetInt(table, i+1, GoValueToLuaValue(L, s.Index(i).Interface()))
		}
		return table
	}

	if rType.Kind() == reflect.Map {
		table := L.NewTable()
		s := reflect.ValueOf(value)
		for _, k := range s.MapKeys() {
			L.RawSet(table, lua.LString(k.String()), GoValueToLuaValue(L, s.MapIndex(k).Interface()))
		}
		return table
	}

	switch v := value.(type) {
	case bool:
		return lua.LBool(v)
	case string:
		return lua.LString(v)
	case int:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	default:
		return lua.LString("<unknown>")
	}
}
