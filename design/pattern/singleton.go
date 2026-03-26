package pattern

import "sync"

// **单例模式**
// 保证一个类只有一个实例，并提供一个全局访问点

// 饿汉式（提前准备好实例）

type Obj1 struct{}

var obj1 *Obj1

func init() {
	obj1 = &Obj1{}
}

func GetObj1() *Obj1 {
	return obj1
}

// 懒汉式

type Obj2 struct{}

var (
	obj2 *Obj2
	once sync.Once
)

func GetObj2() *Obj2 {
	once.Do(func() {
		obj2 = &Obj2{}
	})
	return obj2
}
