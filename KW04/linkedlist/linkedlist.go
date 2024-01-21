package main

import (
	"fmt"
)

/* Eine einfach verkettete Liste ist eine lineare Datensammlung,
in der jedes Element (oder Knoten, engl. Node) einen Wert speichert
und auf das nächste Element in der Liste verweist.
Das letzte Element hat keine "nächstes" Element, der Zeiger ist <nil>
(nicht referenzierter Nullwert).
*/

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) AddAtEnd(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	last := ll.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

/*
Aufgaben:
(1) implementiere die Methode "Size"

	func (ll LinkedList) Size() int

(2) implementiere die "IsInList" Methode

	wie sieht die Signatur dieser Methode aus?

(3) implementiere die Methode "RemoveValue"

	wie sieht die Signatur dieser Methode aus?
		a) lösche den ersten gefundenen Wert
		b) lösche allegefundene Werte
*/
func main() {
	ll := LinkedList{}
	ll.AddAtEnd(1)
	ll.AddAtEnd(2)
	ll.AddAtEnd(3)

	ll.Print() // Output: 1 2 3
}
