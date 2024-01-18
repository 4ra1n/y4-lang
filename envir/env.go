package envir

import "github.com/4ra1n/y4-lang/pool"

const (
	TRUE  = 1
	FALSE = 0
)

type Environment interface {
	Put(name string, value interface{})
	Get(name string) interface{}
	SetOuter(e Environment)
	Symbols() *Symbols
	PutNest(nest, index int, value interface{})
	GetNest(nest, index int) interface{}
	PutNew(name string, value interface{})
	Where(name string) Environment
	GetPool() *pool.Pool
	NewJob(fn func()) bool
	WaitJob() bool
}
