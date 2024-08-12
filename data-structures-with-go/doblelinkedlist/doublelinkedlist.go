package doublelinkedlist

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}
type doublelinkedlist struct {
	head *Node
	tail *Node
}

func (dll *doublelinkedlist) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	if dll.head == nil {
		// Liste boşsa, hem head hem de tail yeni düğüm olmalı
		dll.head = newNode
		dll.tail = newNode
	} else {
		// Liste doluysa, yeni düğümü başa ekleyin
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *doublelinkedlist) InsertAtTheEnd(data int) {
	newNode := &Node{data: data}
	if dll.tail == nil {
		// Liste boşsa, hem head hem de tail yeni düğüm olmalı
		dll.head = newNode
		dll.tail = newNode
	} else {
		// Liste doluysa, yeni düğümü sona  ekleyin
		newNode.prev = dll.head
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *doublelinkedlist) PrintListForward() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")

}

func (dll *doublelinkedlist) PrintListBackwards() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")

}

func (dll *doublelinkedlist) Search(data int) *Node {
	current := dll.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil // Veri bulunamadı
}
func (dll *doublelinkedlist) Delete(data int) {
	current := dll.head

	// Baş düğümü kontrol et
	if current != nil && current.data == data {
		dll.head = current.next
		if dll.head != nil {
			dll.head.prev = nil
		} else {
			dll.tail = nil
		}
		return
	}

	for current != nil && current.data != data {
		current = current.next
	}

	// Düğüm bulunamadı
	if current == nil {
		return
	}

	// Düğüm bulunursa
	if current.prev != nil {
		current.prev.next = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	}
	if current == dll.tail {
		dll.tail = current.prev
	}
}
func DoubleLinkedList() {
	dll := &doublelinkedlist{}
	dll.InsertAtBeginning(10)
	dll.InsertAtBeginning(20)
	dll.InsertAtBeginning(30)
	dll.Search(30)
	dll.PrintListBackwards() // Output: 30 -> 20 -> 10 -> nil
}
