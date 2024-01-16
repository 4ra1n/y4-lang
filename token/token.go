package token

type Token interface {
	IsIdentifier() bool
	IsNumber() bool
	IsString() bool

	GetLineNumber() int
	GetNumber() (interface{}, error)
	GetText() (string, error)

	String() string
}
