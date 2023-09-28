package di

import (
	"fmt"
	"testing"
)

// A 依赖关系 A -> B -> C
type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{
		B: b,
	}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{
		C: c}
}

type C struct {
	Num int
}

// 注意这里是没有入参的
func NewC() *C {
	return &C{
		Num: 1,
	}
}

func TestContainer_Invoke(t *testing.T) {
	container := New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%+v: %d", a, a.B.C.Num)
	})
	fmt.Println("")
	if err != nil {
		panic(err)
	}
	if true != true {
		t.Errorf("错误")
	}
}
