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
	return
}

func Dedupe(v []string) (r []string) {
	seen := make(map[string]struct{})
	for _, vv := range v {
		if _, ok := seen[vv]; !ok {
			seen[vv] = struct{}{}
			r = append(r, vv)
		}
	}
	return
}
