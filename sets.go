package main

type Set struct {
	elements map[any]bool
}

func NewSet() *Set {
	return &Set{
		make(map[any]bool),
	}
}

func (s *Set) Add(el any) {
	s.elements[el] = false
}

func (s *Set) Remove(el any) {
	delete(s.elements, el)
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) List() []any {
	var list []any
	for k := range s.elements {
		list = append(list, k)
	}
	return list
}

func (s *Set) Has(el any) bool {
	_, ok := s.elements[el]
	return ok
}

func (s *Set) Copy() *Set {
	u := NewSet()
	for k := range s.elements {
		u.Add(k)
	}
	return u
}

func Union(sets ...Set) *Set {
	u := sets[0].Copy()
	for _, set := range sets[1:len(sets)] {
		for k := range set.elements {
			u.Add(k)
		}
	}
	return u
}

func Intersect(sets ...Set) *Set {
	i := sets[0].Copy()
	for _, set := range sets[1:len(sets)] {
		for k := range set.elements {
			if !set.Has(k) {
				i.Remove(k)
			}
		}
	}
	return i
}

func (s *Set) Difference(t Set) *Set {
	for k := range t.elements {
		if s.Has(k) {
			delete(s.elements, k)
		}
	}
	return s
}

func (s *Set) IsSubset(t Set) bool {
	for k := range s.elements {
		if !t.Has(k) {
			return false
		}
	}
	return true
}
