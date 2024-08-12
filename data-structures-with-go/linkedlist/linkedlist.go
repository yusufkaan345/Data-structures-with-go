package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	head *Node
}

func (ll *LinkList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
func (ll *LinkList) Search(data int) *Node {
	current := ll.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil // Veri bulunamadı
}
func (ll *LinkList) Delete(data int) {
	current := ll.head
	var previous *Node

	// Baş düğümü kontrol et
	if current != nil && current.data == data {
		ll.head = current.next
		return
	}

	// Orta veya son düğümü kontrol et
	for current != nil && current.data != data {
		previous = current
		current = current.next
	}

	// Düğüm bulunamadı
	if current == nil {
		return
	}

	// Düğüm bulunursa
	previous.next = current.next
}

func LinkedList() {
	ll := &LinkList{}
	ll.InsertAtBeginning(10)
	ll.InsertAtBeginning(20)
	ll.InsertAtBeginning(30)

	ll.PrintList() // Output: 30 -> 20 -> 10 -> nil
}
