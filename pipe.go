package functils

type PipeFunc[I any, O any] func(I) O

func Pipe2[I any, O any, X any](
	fn1 PipeFunc[I, X],
	fn2 PipeFunc[X, O],
) PipeFunc[I, O] {
	return func(input I) O {
		return fn2(fn1(input))
	}
}

func Pipe3[I any, O any, X1 any, X2 any](
	fn1 PipeFunc[I, X1],
	fn2 PipeFunc[X1, X2],
	fn3 PipeFunc[X2, O],
) PipeFunc[I, O] {
	return func(input I) O {
		return fn3(fn2(fn1(input)))
	}
}

func Pipe4[I any, O any, X1 any, X2 any, X3 any](
	fn1 PipeFunc[I, X1],
	fn2 PipeFunc[X1, X2],
	fn3 PipeFunc[X2, X3],
	fn4 PipeFunc[X3, O],
) PipeFunc[I, O] {
	return func(input I) O {
		return fn4(fn3(fn2(fn1(input))))
	}
}

func Pipe5[I any, O any, X1 any, X2 any, X3 any, X4 any](
	fn1 PipeFunc[I, X1],
	fn2 PipeFunc[X1, X2],
	fn3 PipeFunc[X2, X3],
	fn4 PipeFunc[X3, X4],
	fn5 PipeFunc[X4, O],
) PipeFunc[I, O] {
	return func(input I) O {
		return fn5(fn4(fn3(fn2(fn1(input)))))
	}
}

// Type identity function.
// Use this function with its generic type filled in when otherwise the first
// function in the pipe is a generic without its type parameters specified.
func PipeInputType[I any](input I) I {
	return input
}
