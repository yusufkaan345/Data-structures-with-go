package stack

import "fmt"

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	// Son öğeyi al ve kaldır
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Stack'in üstündeki öğeyi döndürür ama kaldırmaz
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Stack'in içeriğini yazdırır
func (s *Stack) Print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Print(s.items[i], " ")
	}
	fmt.Println()
}

func StackCall() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack after pushes:")
	stack.Print() // Output: 30 20 10

	item, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Popped:", item) // Output: Popped: 30
	}

	fmt.Println("Stack after pop:")
	stack.Print() // Output: 20 10

	item, err = stack.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Peeked:", item) // Output: Peeked: 20
	}

	fmt.Println("Is the stack empty?", stack.IsEmpty()) // Output: Is the stack empty? false
}
