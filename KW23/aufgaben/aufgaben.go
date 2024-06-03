package main

/**********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:


Definiere ein Interface in Go, das Methoden enthält, um die linken und
rechten Kinder eines Knotens in einem Heap zurückzugeben. Implementiere
dieses Interface für einen Max-Heap, der als Array dargestellt wird.

Beschreibung:
Erstelle ein Interface HeapNode mit den Methoden Left() und Right().
Implementiere dieses Interface für eine Struktur Heap so, dass die Methoden
die Indizes der linken und rechten Kinder eines Knotens im Heap zurückgeben.
Die Kinder-Indizes sollten für den Fall, dass ein Kind existiert,
zurückgegeben werden; andernfalls sollt -1 (ein ungültiger Index) zurückgegeben
werden.
Beispiel:
Input:

Max-Heap: [42, 29, 18, 14, 7, 18, 12, 11, 5, 13]
Knoten-Index: 1 (Element: 29)
Output:

Linkes Kind: 14
Rechtes Kind: 7

**********************************************************************/

import (
	"fmt"
)

// HeapNode Interface mit Left und Right Methoden
type HeapNode interface {
	Left(index int) int
	Right(index int) int
}

// Heap Struktur, die das HeapNode Interface implementiert
type Heap struct {
	elements []int
}

// NewHeap initialisiert einen neuen Heap
func NewHeap(elements []int) *Heap {
	return &Heap{elements: elements}
}

// Left gibt das linke Kind eines Knotens zurück
func (h *Heap) Left(index int) int {
	leftIndex := 2*index + 1
	if leftIndex < len(h.elements) {
		return h.elements[leftIndex]
	}
	return -1 // Ungültiger Index
}

// Right gibt das rechte Kind eines Knotens zurück
func (h *Heap) Right(index int) int {
	rightIndex := 2*index + 2
	if rightIndex < len(h.elements) {
		return h.elements[rightIndex]
	}
	return -1 // Ungültiger Index
}

func main() {
	heap := NewHeap([]int{42, 29, 18, 14, 7, 18, 12, 11, 5, 13})
	index := 1

	leftChild := heap.Left(index)
	rightChild := heap.Right(index)

	fmt.Printf("Knoten %d (Wert: %d)\n", index, heap.elements[index])
	fmt.Printf("Linkes Kind: %d\n", leftChild)
	fmt.Printf("Rechtes Kind: %d\n", rightChild)
}
