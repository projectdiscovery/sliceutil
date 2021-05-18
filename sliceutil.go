package sliceutil

func PruneEmptyStrings(v []string) []string {
	return PruneEqual(v, "")
}

func PruneEqual(v []string, equalTo string) (r []string) {
	for i := range v {
		if v[i] != equalTo {
			r = append(r, v[i])
		}
	}
	return r
}
