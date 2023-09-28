package observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	sub := &Subject{}
	observer1 := Observer1{}
	observer2 := Observer2{}
	sub.Register(&observer1)
	sub.Register(&observer2)
	sub.Notify("hello")
	sub.Remove(&observer1)
	sub.Notify("world")
}
