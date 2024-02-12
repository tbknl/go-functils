package functils

func SliceTransform[Elem any, Res any](fn func(Elem) Res) func([]Elem) []Res {
	return func(input []Elem) []Res {
		result := make([]Res, len(input))
		for i := range input {
			result[i] = fn(input[i])
		}
		return result
	}
}

func SliceFilter[Elem any](fn func(Elem) bool) func([]Elem) []Elem {
	return func(input []Elem) []Elem {
		var result []Elem
		for _, val := range input {
			if fn(val) {
				result = append(result, val)
			}
		}
		return result
	}
}
