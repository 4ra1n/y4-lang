package envir

type Location struct {
	Nest  int
	Index int
}

func NewLocation(nest, index int) *Location {
	return &Location{Nest: nest, Index: index}
}
