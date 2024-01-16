package envir

import "github.com/4ra1n/y4-lang/log"

type ResizableEnv struct {
	Names *Symbols
	env   *ArrayEnv
}

func NewResizableEnv(size int, poolSize int) *ResizableEnv {
	e := NewArrayEnv(size, poolSize, nil)
	n := NewSymbolsNull()
	r := &ResizableEnv{
		Names: n,
		env:   e,
	}
	log.Debugf("new env with %d - %d", size, poolSize)
	return r
}

func (r *ResizableEnv) Put(name string, value interface{}) {
	e := r.Where(name)
	if e == nil {
		e = r
	}
	e.PutNew(name, value)
}

func (r *ResizableEnv) Get(name string) interface{} {
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

func (r *ResizableEnv) SetOuter(e Environment) {
	r.env.SetOuter(e)
}

func (r *ResizableEnv) Symbols() *Symbols {
	return r.Names
}

func (r *ResizableEnv) PutNest(nest, index int, value interface{}) {
	r.env.PutNest(nest, index, value)
}

func (r *ResizableEnv) GetNest(nest, index int) interface{} {
	return r.env.GetNest(nest, index)
}

func (r *ResizableEnv) PutNew(name string, value interface{}) {
	r.Assign(r.Names.PutNew(name), value)
}

func (r *ResizableEnv) Where(name string) Environment {
	_, ok := r.Names.Find(name)
	if ok {
		return r
	} else if r.env.Outer == nil {
		return nil
	} else {
		return r.env.Outer.Where(name)
	}
}

func (r *ResizableEnv) Assign(index int, value interface{}) {
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

func (r *ResizableEnv) NewJob(fn func()) bool {
	return r.env.NewJob(fn)
}

func (r *ResizableEnv) WaitJob() bool {
	return r.env.WaitJob()
}
