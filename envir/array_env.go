package envir

import (
	"github.com/4ra1n/y4-lang/log"
	"github.com/4ra1n/y4-lang/pool"
)

// ArrayEnv
// 适用于自定义函数内部的环境
// 为什么这里敢直接写死数组大小
// 在创建函数内环境之前会 Lookup 实际需求大小
// 根据需求 size 创建
type ArrayEnv struct {
	Values []interface{}
	Outer  Environment
	GoPool *pool.Pool
}

func NewArrayEnv(size int, p *pool.Pool, out Environment) *ArrayEnv {
	a := &ArrayEnv{
		Values: make([]interface{}, size),
		Outer:  out,
		GoPool: p,
	}
	return a
}

func (a *ArrayEnv) GetPool() *pool.Pool {
	return a.GoPool
}

func (a *ArrayEnv) error(name string) {
	log.Errorf("数组环境禁止操作: %s", name)
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
	log.Error("数组环境不允许查询符号")
	return nil
}

func (a *ArrayEnv) PutNest(nest, index int, value interface{}) {
	if nest == 0 {
		a.Values[index] = value
	} else if a.Outer == nil {
		log.Error("数组环境没有外部环境")
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

func (a *ArrayEnv) Clone() Environment {
	// 注意直接复制的是指针
	// 可能被后续操作修改
	newValues := make([]interface{}, len(a.Values))
	copy(newValues[:], a.Values[:])
	return &ArrayEnv{
		Values: newValues,
		Outer:  a.Outer,
		GoPool: a.GoPool,
	}
}
