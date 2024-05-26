package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(0)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(9) - rand.Intn(9)
	}
	return slice
}

func ExampleTopKFrequent() {
	//Output:
	//...
	s := generateSlice(20)
	fmt.Println(s)
	fmt.Println(TopKFrequent(s, 3))
}

func TestTopK(t *testing.T) {
	// input := ?
	// expected := ?
	// Funktion aufrufen
	// falls f(input!=expected) -> Fehler:
	// 	t.Errorf("FEHLER!")
	// für einen Vergleich von Listen kannst du reflect.DeepEqual verwenden
}
