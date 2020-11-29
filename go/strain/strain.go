package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep filters Ints, keeping elements that return true when run through f.
func (i Ints) Keep(f func(n int) bool) Ints {
	var ret Ints
	for _, val := range i {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// Discard filters Ints, keeping elements that return false when run through f.
func (i Ints) Discard(f func(n int) bool) Ints {
	var ret Ints
	for _, val := range i {
		if !f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// Keep filters Lists, keeping elements that return true when run through f.
func (l Lists) Keep(f func(l []int) bool) Lists {
	var ret Lists
	for _, val := range l {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

// Keep filters Strings, keeping elements that return true when run through f.
func (s Strings) Keep(f func(s string) bool) Strings {
	var ret Strings
	for _, val := range s {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}
