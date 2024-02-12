package functils

type KV[K comparable, V any] struct {
	Key   K
	Value V
}

func MapEntries[M ~map[K]V, K comparable, V any](m M) []KV[K, V] {
	r := make([]KV[K, V], 0, len(m))
	for k, v := range m {
		r = append(r, KV[K, V]{k, v})
	}
	return r
}

func MapFromEntries[Slice ~[]KV[K, V], K comparable, V any](s Slice) map[K]V {
	r := make(map[K]V, len(s))
	for _, kv := range s {
		r[kv.Key] = kv.Value
	}
	return r
}

func MapFilter[M ~map[K]V, K comparable, V any](fn func(K, V) bool) func(M) M {
	return func(input M) M {
		r := make(map[K]V)
		for k, v := range input {
			if fn(k, v) {
				r[k] = v
			}
		}
		return r
	}
}

func MapFilterKeys[M ~map[K]V, K comparable, V any](fn func(K) bool) func(M) M {
	return func(input M) M {
		r := make(map[K]V)
		for k, v := range input {
			if fn(k) {
				r[k] = v
			}
		}
		return r
	}
}

func MapFilterValues[M ~map[K]V, K comparable, V any](fn func(V) bool) func(M) M {
	return func(input M) M {
		r := make(map[K]V)
		for k, v := range input {
			if fn(v) {
				r[k] = v
			}
		}
		return r
	}
}

func MapTransform[IM ~map[IK]IV, OM map[OK]OV, IK comparable, IV any, OK comparable, OV any](
	fn func(IK, IV) (OK, OV),
) func(IM) OM {
	return func(input IM) OM {
		r := make(map[OK]OV, len(input))
		for ik, iv := range input {
			ok, ov := fn(ik, iv)
			r[ok] = ov
		}
		return r
	}
}

func MapTransformKeys[IM ~map[IK]V, OM map[OK]V, IK comparable, OK comparable, V any](
	fn func(IK) OK,
) func(IM) OM {
	return func(input IM) OM {
		r := make(map[OK]V, len(input))
		for ik, v := range input {
			ok := fn(ik)
			r[ok] = v
		}
		return r
	}
}

func MapTransformValues[IM ~map[K]IV, OM map[K]OV, K comparable, IV any, OV any](
	fn func(IV) OV,
) func(IM) OM {
	return func(input IM) OM {
		r := make(map[K]OV, len(input))
		for k, iv := range input {
			ov := fn(iv)
			r[k] = ov
		}
		return r
	}
}
