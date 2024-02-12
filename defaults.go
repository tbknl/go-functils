package functils

func DefaultValueOnErr[In any, Out any](defaultValue Out, fn func(In) (Out, error)) func(In) Out {
	return func(input In) Out {
		val, err := fn(input)
		if err == nil {
			return val
		} else {
			return defaultValue
		}
	}
}

func DefaultOnErr[In any, Out any](fn func(In) (Out, error)) func(In) Out {
	var defaultValue Out
	return DefaultValueOnErr(defaultValue, fn)
}

func DefaultValueOnFail[In any, Out any](defaultValue Out, fn func(In) (Out, bool)) func(In) Out {
	return func(input In) Out {
		val, ok := fn(input)
		if ok {
			return val
		} else {
			return defaultValue
		}
	}
}

func DefaultOnFail[In any, Out any](fn func(In) (Out, bool)) func(In) Out {
	var defaultValue Out
	return DefaultValueOnFail(defaultValue, fn)
}

func DefaultValueOnNil[T any](defaultValue T) func(*T) T {
	return func(input *T) T {
		if input == nil {
			return defaultValue
		} else {
			return *input
		}
	}
}

func DefaultOnNil[T any](t *T) T {
	var defaultValue T
	return DefaultValueOnNil(defaultValue)(t)
}
