package observer

import "fmt"

type IObserver interface {
	Update(msg string)
}
type Observer1 struct{}

type Observer2 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}
func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}
type Subject struct {
	arr []IObserver
}

func (sub *Subject) Register(observer IObserver) {
	sub.arr = append(sub.arr, observer)
}
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.arr {
		if ob == observer {
			sub.arr = append(sub.arr[:i], sub.arr[i+1:]...)
			return
		}
	}
}
func (sub *Subject) Notify(msg string) {
	for _, ob := range sub.arr {
		ob.Update(msg)
	}
}
