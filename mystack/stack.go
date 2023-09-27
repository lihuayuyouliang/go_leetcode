package mystack

type Stack struct {
	count int
	arr   []string
}

func NewStack() *Stack {
	return &Stack{
		count: 0,
		arr:   make([]string, 5),
	}
}
func (stack *Stack) push(str string) {
	stack.count++
	if stack.count == len(stack.arr) {
		stack.arr = append(stack.arr, str)
		return
	}
	stack.arr[stack.count] = str
}
func (stack *Stack) pop() string {
	if stack.count == 0 {
		return ""
	}
	item := stack.arr[stack.count]
	stack.count--
	return item
}
