package template

func ternary(v bool, vt interface{}, vf interface{}) interface{} {
	if v {
		return vt
	}

	return vf
}
