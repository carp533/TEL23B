package main

import "container/heap"

/* Aufgabe: ermittle die K häufigsten Elemente einer Liste
hierfür möchten wir einen Heap verwenden.
A) Schaue dir zunächst die Fehlermeldung an und korrigiere diese, indem
du die fehlenden Methoden programmierst. Du kannst die Methoden manuell
anlegen oder den Quick-Fix (bei der Fehlermeldung) verwenden
B) Schaue dir Methode Push an. Was bewirkt die Zeile
	m := num.(Num)
   ? Welchen Datentyp hat num?
C) programmiere die Funktion topKFrequent
D) programmiere eine Beispielfunktion und eine Testfunktion
*/

func TopKFrequent(nums []int, k int) []int {
	// count := make(map[int]int)
	// TODO: erstelle eine map mit den Elementen aus nums und deren Häufigkeit
	nums_ := make(Nums, 0)
	// TODO: wandle die map in die Liste nums_ um
	heap.Init(&nums_)
	var res []int
	// TODO: rufe k mal Pop auf des Heaps und schreibe das Ergebnis in die Liste res
	return res
}

// Num stores its value and frequency as Count.
type Num struct {
	Val   int
	Count int
}

// Nums struct for impl Interface
type Nums []Num

// Push heap Interface
func (n *Nums) Push(num interface{}) {
	m := num.(Num)
	*n = append(*n, m)
}

// Pop heap Interface
func (n *Nums) Pop() interface{} {
	res := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return res
}
