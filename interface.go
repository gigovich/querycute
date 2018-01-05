package querycute

// Mapper definition
type Mapper interface {
	// GetMapping object and values for bind
	GetMapping() (Mapping, []interface{})
}

// ValuesBinder defines hook method which will be called after query return values
type ValuesBinder interface {
	// OnValuesBind recieve fields names present in query and associated values received from query
	OnValuesBind(fields []string, vars []interface{}) error
}
