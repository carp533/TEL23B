/*
*********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:
// Implementiere eine Funktion in Go, die einen Zähler erzeugt. Jedes Mal, wenn der Zähler aufgerufen wird, soll er um 1 erhöht werden und den aktuellen Zählerstand zurückgeben.

// Beschreibung:
// Schreibe eine Funktion NewCounter, die eine Zählerfunktion zurückgibt.
// Die Zählerfunktion soll jedes Mal, wenn sie aufgerufen wird, den Zähler um 1 erhöhen und den aktuellen Zählerstand zurückgeben.
// Nutze ein Closure, um den Zustand des Zählers zwischen den Aufrufen zu speichern.
// counter := NewCounter()
// fmt.Println(counter()) // Ausgabe: 1
// fmt.Println(counter()) // Ausgabe: 2
// fmt.Println(counter()) // Ausgabe: 3
*/
package main

import (
	"fmt"
)

// NewCounter gibt eine Zählerfunktion zurück
func NewCounter() func() int {

}

func main() {
	counter := NewCounter()

	fmt.Println(counter()) // Ausgabe: 1
	fmt.Println(counter()) // Ausgabe: 2
	fmt.Println(counter()) // Ausgabe: 3

	// Ein weiterer Zähler zur Demonstration, dass jeder eigene Zustand speichert
	anotherCounter := NewCounter()
	fmt.Println(anotherCounter()) // Ausgabe: 1
	fmt.Println(anotherCounter()) // Ausgabe: 2

	// Der erste Zähler bleibt unverändert
	fmt.Println(counter()) // Ausgabe: 4
}
