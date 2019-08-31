package stack

type Stack struct {
	top  *Item
	size int
}

type Item struct {
	value interface{}
	next  *Item
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.size > 0 {
		value := stack.top.value
		stack.top = stack.top.next
		stack.size--
		return value
	}

	return nil
}

func (stack *Stack) Get() interface{} {
	if stack.size > 0 {
		return stack.top.value
	}
	return nil
}

func (stack *Stack) GetSize() int {
	return stack.size
}
