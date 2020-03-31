package set

import (
	"fmt"
	"strings"
)

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(xs ...interface{}) {
	for _, x := range xs {
		s[x] = struct{}{}
	}
}

func (s Set) Remove(xs ...interface{}) {
	for _, x := range xs {
		if _, ok := s[x]; ok {
			delete(s, x)
		}
	}
}

func (s Set) Contains(x interface{}) bool {
	_, ok := s[x]
	return ok
}

func (s Set) Clear() {
	for x := range s {
		delete(s, x)
	}
}

func (s Set) Intersection(other Set) Set {
	zs := NewSet()
	for x := range s {
		if other.Contains(x) {
			zs.Add(x)
		}
	}

	return zs
}

func (s Set) Complement(other Set) Set {
	zs := NewSet()
	for x := range s {
		if !other.Contains(x) {
			zs.Add(x)
		}
	}

	return zs
}

func (s Set) Union(other Set) Set {
	zs := NewSet()

	for x := range s {
		zs.Add(x)
	}
	for x := range other {
		zs.Add(x)
	}

	return zs
}

func (s Set) Difference(other Set) Set {
	return s.Complement(other).Union(other.Complement(s))
}

func (s Set) IsSubsetOf(other Set) bool {
	for x := range s {
		if !other.Contains(x) {
			return false
		}
	}

	return true
}

func (s Set) IsSupersetOf(other Set) bool {
	return other.IsSubsetOf(s)
}

func (s Set) IsEqual(other Set) bool {
	return s.IsSubsetOf(other) && s.IsSupersetOf(other)
}

func (s Set) String() string {
	xs, i := make([]string, len(s)), 0

	for x := range s {
		xs[i] = fmt.Sprintf("%#v", x)
		i++
	}

	return fmt.Sprintf("{%s}", strings.Join(xs, ", "))

}
