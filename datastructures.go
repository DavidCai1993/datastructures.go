package datastructures

// Comparable represents a type which is comparable.
type Comparable interface {
	Compare(Comparable) int
}

// Comparables represents []Comparable .
type Comparables []Comparable

// Len is to implements Sort.Interface .
func (cs Comparables) Len() int {
	return len(cs)
}

// Less is to implements Sort.Interface .
func (cs Comparables) Less(i, j int) bool {
	if cs[i].Compare(cs[j]) < 0 {
		return true
	}

	return false
}

// Shift removes the fisrt element and returns it.
func (cs *Comparables) Shift() Comparable {
	if len(*cs) == 0 {
		return nil
	}

	f := (*cs)[0]
	*cs = (*cs)[1:]
	return f
}

// Swap is to implements Sort.Interface .
func (cs Comparables) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// IntComparable represents a comparable int.
type IntComparable int

// Compare implements the Comparable interface.
func (ic IntComparable) Compare(i Comparable) int {
	other, ok := i.(IntComparable)

	if !ok {
		panic("param should be an intComparable")
	}

	if int(ic) == int(other) {
		return 0
	} else if int(ic) < int(other) {
		return -1
	} else {
		return 1
	}
}
