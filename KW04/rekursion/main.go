package main

import "fmt"

/*
Rekursion in der Informatik bezeichnet den Prozess, bei
dem eine Funktion sich selbst aufruft. Diese Technik wird
häufig verwendet, um Probleme zu lösen, die in kleinere,
ähnliche Probleme aufgeteilt werden können. Rekursive
Funktionen benötigen eine Abbruchbedingung, um zu
verhindern, dass sie unendlich oft ausgeführt werden.
*/
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/*Aufgabe: programmiere die Multiplikation zweier Zahlen
mit hilfe einer rekursiven Funktion und Addition
func mult(a,b int) int
Hinweis: nutze aus, dass (y+1)*x = y*x+x, also
	mult(0,b)   = 0
	mult(a+1,b) = mult(a,b)+b
*/

/*Aufgabe programmiere die Ackermann Funktion.
  A(m, n) =
  n+1                              if m = 0
  A(m-1, 1)                        if m > 0 and n = 0
  A(m-1, A(m, n-1))                if m > 0 and n > 0
  Bonusaufgabe: zähle die Anzahl der rekursiven Aufrufe
  Hinweis erweitere die Funktion um einen int-Pointer, der
  bei jedem Aufruf inkrementiert wird.

*/
func ackermann(m, n int) int {
	//TODO
	return 0
}

func main() {
	fmt.Printf("fact(7)=%v\n", fact(7))
	fmt.Printf("fib(7)=%v\n", fib(7))
	fmt.Printf("ackermann(3, 2)=%v\n", ackermann(3, 2)) //->29
	//fmt.Printf("ackermann(4, 1)=%v\n", ackermann(4, 1)) //->65533
}
