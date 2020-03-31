package set

import (
	"fmt"
	"strings"
)

type MultiSet map[interface{}]int

func NewMultiSet() MultiSet {
	return make(MultiSet)
}

func (m MultiSet) Add(xs ...interface{}) {
	for _, x := range xs {
		m[x]++
	}
}

func (m MultiSet) Remove(xs ...interface{}) {
	for _, x := range xs {
		if n, ok := m[x]; ok {
			if n > 1 {
				m[x]--
			} else {
				delete(m, x)
			}
		}
	}
}

func (m MultiSet) Contains(x interface{}) bool {
	_, ok := m[x]
	return ok
}

func (m MultiSet) Clear() {
	for x := range m {
		delete(m, x)
	}
}

func (m MultiSet) Intersection(other MultiSet) MultiSet {
	zs := NewMultiSet()

	for x, c := range m {
		if other.Contains(x) {
			if c < other[x] {
				zs[x] = c
			} else {
				zs[x] = other[x]
			}
		}
	}
	return zs
}

func (m MultiSet) Complement(other MultiSet) MultiSet {
	zs := NewMultiSet()

	for x, c := range m {
		if !other.Contains(x) {
			zs[x] = c
		} else {
			if c > other[x] {
				zs[x] = c - other[x]
			}
		}
	}

	return zs
}

func (m MultiSet) Union(other MultiSet) MultiSet {
	zs := NewMultiSet()

	for x, c := range m {
		zs[x] += c
	}

	for x, c := range other {
		zs[x] += c
	}

	return zs
}

func (m MultiSet) Difference(other MultiSet) MultiSet {
	return m.Complement(other).Union(other.Complement(m))
}

func (m MultiSet) IsSubsetOf(other MultiSet) bool {
	for x, c := range m {
		if !other.Contains(x) || c > other[x] {
			return false
		}
	}

	return true
}

func (m MultiSet) IsSupersetOf(other MultiSet) bool {
	return other.IsSubsetOf(m)
}

func (m MultiSet) IsEqual(other MultiSet) bool {
	return m.IsSubsetOf(other) && m.IsSupersetOf(other)
}

func (m MultiSet) String() string {
	xs, i := make([]string, len(m)), 0

	for x, c := range m {
		xs[i] = fmt.Sprintf("(key: %#v, count: %d)", x, c)
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(xs, ", "))
}
