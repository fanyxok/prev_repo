package functional

func MapDiff[K comparable, V comparable](lhs map[K]V, rhs map[K]V) (map[K]Pair[V, V], map[K]V, map[K]V) {
	bothHave := make(map[K]Pair[V, V])
	lhsHave := make(map[K]V)
	rhsHave := make(map[K]V)
	for i, v := range lhs {
		if v2, ok := rhs[i]; ok {
			if v2 != v {
				bothHave[i] = Pair[V, V]{v, v2}
			}
		} else {
			lhsHave[i] = v
		}
	}

	for i, v := range rhs {
		if _, ok := lhs[i]; !ok {
			rhsHave[i] = v
		}
	}
	return bothHave, lhsHave, rhsHave
}

func MaxKey[K uint32 | int32, V any](map_ map[K]V) K {
	var max K
	for max = range map_ {
		break
	}
	for n := range map_ {
		if n > max {
			max = n
		}
	}
	return max
}
