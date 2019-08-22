package engine

func isSingleAndEmpty(a []string) bool {
	return len(a) == 1 && a[0] == ""
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
