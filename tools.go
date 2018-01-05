package querycute

// filterFields by it names listed in include, but if include is empty nothing changed
func filterFields(fields []string, values []interface{}, include ...string) ([]string, []interface{}) {
	if (len(fields) != len(values)) || len(include) == 0 {
		return fields, values
	}

	filter := make(map[string]struct{})
	for _, field := range include {
		filter[field] = struct{}{}
	}

	var rNames []string
	var rFields []interface{}
	for i, name := range fields {
		if _, ok := filter[name]; ok {
			rFields[i] = values[i]
		}
	}

	return rNames, rFields
}
