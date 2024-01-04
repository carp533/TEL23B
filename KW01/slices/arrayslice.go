package main

import (
	"fmt"
)

func main() {

	//definiert einen Array mit String der Länge 2.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Definiere eine Liste von Zahlen.
	// Man nennt das eine "int-Slice".
	l1 := []int{10, 20, 30, 40, 50}

	// Den Typ von l1 ausgeben:
	fmt.Printf("Typ von l1: %T\n", l1)
	// Den Wert von l1 ausgeben:
	fmt.Printf("Wert von l1: %v\n", l1)

	// Das Element an Stelle 0 in l1 ausgeben:
	fmt.Println(l1[0])

	// Alle weiteren Elemente ausgeben:
	fmt.Println(l1[1])
	fmt.Println(l1[2])
	fmt.Println(l1[3])
	fmt.Println(l1[4])
	// Es gibt 5 Elemente an den Positionen 0,1,2,3,4.

	// Die Länge von l1 ausgeben:
	fmt.Println(len(l1))

	// Alle Elemente der Liste in einer Schleife ausgeben:
	for i := 0; i < len(l1); i++ {
		fmt.Printf("%v: %v\n", i, l1[i])
	}

	// Bessere Version der obigen Schleife:
	for index, value := range l1 {
		fmt.Printf("%v: %v\n", index, value)
	}

	// Liste an Stelle 3 schreiben:
	l1[3] = 500
	fmt.Println(l1)

	matrix := [][]int{
		{1, 2, 3}, // Typ: []int
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	// Das Element in der Mitte verändern und ausgeben:
	matrix[1][1] = 500
	fmt.Println(matrix)
	fmt.Println(matrix[1][1])

	// Vorsicht Falle:
	row := []int{1, 2, 3}
	matrix2 := [][]int{row, row, row}
	fmt.Println(matrix2)
	matrix2[0][2] = 42
	fmt.Println(matrix2)

	//Listen können mit dem eingbauten Befehl make erstellt werden
	//make(<Datentyp>, <Länge>, <Kapazität>)
	e := make([]int, 5)
	printSlice("a", e)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	//mit dem eingebauten Befehl append können Elemente zu der Liste hinzugefügt werden.
	for i := 0; i < 5; i++ {
		d = append(d, i)
	}
	printSlice("d", d)

}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
