package envir

import (
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pool"
)

type GlobalEnv struct {
	Names *Symbols
	env   *ArrayEnv
	pool  *pool.Pool
}

func NewResizableEnv(size int, poolSize int) *GlobalEnv {
	p := pool.NewPool(poolSize)
	e := NewArrayEnv(size, p, nil)
	n := NewSymbolsNull()
	r := &GlobalEnv{
		Names: n,
		env:   e,
		pool:  p,
	}
	log.Debugf("new env with %d - %d", size, poolSize)
	return r
}

func (r *GlobalEnv) GetPool() *pool.Pool {
	return r.pool
}

func (r *GlobalEnv) Put(name string, value interface{}) {
	e := r.Where(name)
	if e == nil {
		e = r
	}
	e.PutNew(name, value)
}

func (r *GlobalEnv) Get(name string) interface{} {
	i, ok := r.Names.Find(name)
	if !ok {
		if r.env.Outer == nil {
			return nil
		} else {
			return r.env.Outer.Get(name)
		}
	} else {
		return r.env.Values[i]
	}
}

func (r *GlobalEnv) SetOuter(e Environment) {
	r.env.SetOuter(e)
}

func (r *GlobalEnv) Symbols() *Symbols {
	return r.Names
}

func (r *GlobalEnv) PutNest(nest, index int, value interface{}) {
	r.env.PutNest(nest, index, value)
}

func (r *GlobalEnv) GetNest(nest, index int) interface{} {
	return r.env.GetNest(nest, index)
}

func (r *GlobalEnv) PutNew(name string, value interface{}) {
	r.Assign(r.Names.PutNew(name), value)
}

func (r *GlobalEnv) Where(name string) Environment {
	_, ok := r.Names.Find(name)
	if ok {
		return r
	} else if r.env.Outer == nil {
		return nil
	} else {
		return r.env.Outer.Where(name)
	}
}

func (r *GlobalEnv) Assign(index int, value interface{}) {
	if index > len(r.env.Values) {
		newLen := len(r.env.Values) * 2
		if index > newLen {
			newLen = index + 1
		}
		newValues := make([]interface{}, newLen)
		copy(newValues, r.env.Values)
		r.env.Values = newValues
	}
	r.env.Values[index] = value
}

func (r *GlobalEnv) NewJob(fn func()) bool {
	return r.env.NewJob(fn)
}

func (r *GlobalEnv) WaitJob() bool {
	return r.env.WaitJob()
}
