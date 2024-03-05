package algorithms

func Merge(m1, m2 map[string]interface{}) map[string]interface{} {
	// lifted from helm's merge code
	out := make(map[string]interface{}, len(m1))
	for k, v := range m1 {
		out[k] = v
	}

	for k, v := range m2 {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = Merge(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}
