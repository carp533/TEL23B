package main

import "fmt"

/* Alle Tests sollen funktionieren - Testfunktionen nicht ändern!

Aufgabe 1a:
berechne die "Hailstone sequence", definiert durch:
Starte mit einer positiven Zahl n:
Wenn n == 1 Ende.
Wenn n   gerade ist: n -> n/2
Wenn n ungerade ist: n -> (3 * n) + 1
*/
func Hs(n int) []int {
	result := make([]int, 0)
	result = append(result, n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		result = append(result, n)
	}
	return result
}

func ExampleHs() {
	//Output:
	//[5 16 8 4 2 1]
	seq := Hs(5)
	fmt.Println(seq)
}

/*Aufgabe 1b:
berechne alle Hailstone Sequenzen < 100.000 und gib den Startwert und Länge der
längsten Sequenz aus.
*/
func HsMax(max int) (maxN, maxLen int) {
	maxN, maxLen = 1, 1
	for i := 1; i <= max; i++ {
		l := len(Hs(i))
		if l > maxLen {
			maxN, maxLen = i, l
		}
	}
	return maxN, maxLen
}

func ExampleHsMax() {
	//Output:
	//hs(77031): 351 Elemente
	maxN, maxLen := HsMax(100000)
	fmt.Printf("hs(%d): %d Elemente\n", maxN, maxLen)
}
