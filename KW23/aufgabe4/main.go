package main

/**********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:
* In der Datei avl.go findest du eine Implementierung eines AVL-Baumes.
* Du sollst die Implementierung so anpassen, dass keine Werte doppelt in
* den AVL Baum eingefügt werden.
* schreibe eine Test- oder Beispielfunktion dazu.
**********************************************************************/

import (
	"encoding/json"
	"fmt"
	"log"
)

type intKey int

// satisfy Key
func (k intKey) Less(k2 Key) bool { return k < k2.(intKey) }
func (k intKey) Eq(k2 Key) bool   { return k == k2.(intKey) }

// use json for cheap tree visualization
func dump(tree *Node) {
	b, err := json.MarshalIndent(tree, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func main() {
	var tree *Node
	fmt.Println("Empty tree:")
	dump(tree)

	fmt.Println("\nInsert test:")
	Insert(&tree, intKey(3))
	Insert(&tree, intKey(1))
	Insert(&tree, intKey(4))
	Insert(&tree, intKey(1))
	Insert(&tree, intKey(5))
	dump(tree)

	fmt.Println("\nRemove test:")
	Remove(&tree, intKey(3))
	Remove(&tree, intKey(1))
	dump(tree)
}
