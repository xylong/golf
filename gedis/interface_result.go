package gedis

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (r *InterfaceResult) Unwrap() interface{} {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

func (r *InterfaceResult) UnwrapOr(any interface{}) interface{} {
	if r.Err != nil {
		return any
	}
	return r.Result
}
