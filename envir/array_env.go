package envir

import (
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pool"
)

type ArrayEnv struct {
	Values []interface{}
	Outer  Environment
	GoPool *pool.Pool
}

func NewArrayEnv(size int, out Environment) *ArrayEnv {
	a := &ArrayEnv{
		Values: make([]interface{}, size),
		Outer:  out,
		GoPool: pool.NewPool(10),
	}
	return a
}

func (a *ArrayEnv) error(name string) {
	log.Errorf("cannot access: %s", name)
}

func (a *ArrayEnv) Put(name string, _ interface{}) {
	a.error(name)
}

func (a *ArrayEnv) Get(name string) interface{} {
	a.error(name)
	return nil
}

func (a *ArrayEnv) SetOuter(e Environment) {
	a.Outer = e
}

func (a *ArrayEnv) Symbols() *Symbols {
	log.Error("array envir not allow symbols")
	return nil
}

func (a *ArrayEnv) PutNest(nest, index int, value interface{}) {
	if nest == 0 {
		a.Values[index] = value
	} else if a.Outer == nil {
		log.Error("array envir no outer environment")
	} else {
		a.Outer.PutNest(nest-1, index, value)
	}
}

func (a *ArrayEnv) GetNest(nest, index int) interface{} {
	if nest == 0 {
		return a.Values[index]
	} else if a.Outer == nil {
		return nil
	} else {
		return a.Outer.GetNest(nest-1, index)
	}
}

func (a *ArrayEnv) PutNew(name string, _ interface{}) {
	a.error(name)
}

func (a *ArrayEnv) Where(name string) Environment {
	a.error(name)
	return nil
}

func (a *ArrayEnv) NewJob(fn func()) bool {
	a.GoPool.AddJob(fn)
	return true
}

func (a *ArrayEnv) WaitJob() bool {
	a.GoPool.StopAll()
	a.GoPool.Wait()
	return true
}
