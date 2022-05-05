package goden

type ResultType interface {
	int | string
}

type ResultOption[T ResultType] struct {
	Result 	T
	Err 	error
}

func Result[T ResultType] (result T, err error) *ResultOption[T] {
	return &ResultOption[T]{
		Result: result,
		Err: err,
	}
}

func (this *ResultOption[T]) Unwrap() T {
	if this.Err != nil {
		panic(this.Err)
	}
	return this.Result
}

func (this *ResultOption[T]) UnwrapOr(ret T) T {
	if this.Err != nil {
		return ret
	}
	return this.Result
}

