package gedis

type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

func (r *StringResult) Unwrap() string {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

func (r *StringResult) UnwrapOr(str string) string {
	if r.Err != nil {
		return str
	}
	return r.Result
}
