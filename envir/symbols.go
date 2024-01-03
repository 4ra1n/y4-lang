package envir

import "github.com/4ra1n/y4-lang/base"

type Symbols struct {
	Outer *Symbols
	Table *base.Map[string, int]
}

func NewSymbolsNull() *Symbols {
	return NewSymbols(nil)
}

func NewSymbols(outer *Symbols) *Symbols {
	return &Symbols{
		Outer: outer,
		Table: base.NewMap[string, int](),
	}
}

func (s *Symbols) Size() int {
	return s.Table.Length()
}

func (s *Symbols) Append(other *Symbols) {
	for _, key := range other.Table.Keys() {
		val, ok := other.Table.Get(key)
		if ok {
			s.Table.Set(key, val)
		}
	}
}

func (s *Symbols) Find(key string) (int, bool) {
	value, exists := s.Table.Get(key)
	return value, exists
}

func (s *Symbols) Get(key string) *Location {
	return s.GetNest(key, 0)
}

func (s *Symbols) GetNest(key string, nest int) *Location {
	index, exists := s.Table.Get(key)
	if !exists {
		if s.Outer == nil {
			return nil
		}
		return s.Outer.GetNest(key, nest+1)
	}
	return NewLocation(nest, index)
}

func (s *Symbols) PutNew(key string) int {
	if index, exists := s.Find(key); exists {
		return index
	}
	return s.Add(key)
}

func (s *Symbols) Put(key string) *Location {
	loc := s.GetNest(key, 0)
	if loc == nil {
		return NewLocation(0, s.Add(key))
	}
	return loc
}

func (s *Symbols) Add(key string) int {
	index := s.Table.Length()
	s.Table.Set(key, index)
	return index
}
