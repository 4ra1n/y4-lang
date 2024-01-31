package envir

import (
	"github.com/4ra1n/y4-lang/base"
	"github.com/4ra1n/y4-lang/pool"
)

type BasicEnv struct {
	Values *base.Map[string, interface{}]
	Outer  Environment
	GoPool *pool.Pool
}

func NewBasicEnvWithEnv(e Environment) *BasicEnv {
	en := &BasicEnv{
		Values: base.NewMap[string, interface{}](),
		Outer:  e,
		GoPool: e.GetPool(),
	}
	return en
}

func (b *BasicEnv) GetPool() *pool.Pool {
	return b.GoPool
}

func (b *BasicEnv) Put(name string, value interface{}) {
	en := b.Where(name)
	if en == nil {
		en = b
	}
	en.PutNew(name, value)
}

func (b *BasicEnv) Get(name string) interface{} {
	v, _ := b.Values.Get(name)
	if v == nil && b.Outer != nil {
		return b.Outer.Get(name)
	} else {
		return v
	}
}

func (b *BasicEnv) SetOuter(e Environment) {
	b.Outer = e
}

func (b *BasicEnv) Symbols() *Symbols {
	return nil
}

func (b *BasicEnv) PutNest(nest, index int, value interface{}) {
}

func (b *BasicEnv) GetNest(nest, index int) interface{} {
	return nil
}

func (b *BasicEnv) PutNew(name string, value interface{}) {
	b.Values.Set(name, value)
}

func (b *BasicEnv) Where(name string) Environment {
	n, _ := b.Values.Get(name)
	if n != nil {
		return b
	}
	if b.Outer == nil {
		return nil
	}
	return b.Outer.Where(name)
}

func (b *BasicEnv) NewJob(fn func()) bool {
	b.GoPool.AddJob(fn)
	return true
}

func (b *BasicEnv) WaitJob() bool {
	b.GoPool.StopAll()
	b.GoPool.Wait()
	return true
}

func (b *BasicEnv) Clone() Environment {
	newValues := base.NewMap[string, interface{}]()
	for _, k := range b.Values.Keys() {
		v, _ := b.Values.Get(k)
		newValues.Set(k, v)
	}
	return &BasicEnv{
		Values: newValues,
		Outer:  b.Outer,
		GoPool: b.GoPool,
	}
}
