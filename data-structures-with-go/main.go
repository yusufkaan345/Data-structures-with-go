package main

import (
	doublelinkedlist "data-structures/doblelinkedlist"
	"data-structures/fibonacci"
	linkedList "data-structures/linkedlist"
	"data-structures/stack"
	"fmt"
)

//fibonacci "data-structures/fibonacci"
//linkedList "data-structures/linkedlist"
//"fmt"

func main() {
	fmt.Println(fibonacci.Fibonacci(6, nil))
	linkedList.LinkedList()
	doublelinkedlist.DoubleLinkedList()
	stack.StackCall()
}
