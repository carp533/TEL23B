package main

import (
	"fmt"
	"strings"
)

/**********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:
Implementiere ein Lindenmayer-System (L-System) in Go, das eine bestimmte Anzahl von Iterationen durchführt und die resultierende Zeichenkette ausgibt.

Beschreibung:
Ein L-System besteht aus einem Startaxiom (einer Anfangszeichenkette) und Produktionsregeln, die angeben, wie jede Zeichenkette transformiert werden soll.
Deine Aufgabe ist es, das L-System zu implementieren, sodass es eine bestimmte Anzahl von Iterationen durchführt und die resultierende Zeichenkette ausgibt.
Beispiel:
Startaxiom: "A"
Produktionsregeln:
"A" -> "AB"
"B" -> "A"
Iterationen: 4
Erwartetes Ergebnis nach 4 Iterationen:

Iteration 0: "A"
Iteration 1: "AB"
Iteration 2: "ABA"
Iteration 3: "ABAAB"
Iteration 4: "ABAABABA"
Bemerkung: in koch.go ist ein Programm, mit dem man aus dem Ergebnis Turtle-Graphics Bilder erzeugen kann.
Wer das testen möchte, muss ein externes modul installieren (go get / go mod tidy)
/***********************************************************************/
// LSystem struct, das das Startaxiom und die Produktionsregeln enthält
type LSystem struct {
	axiom  string
	rules  map[rune]string
	result string
}

// NewLSystem erstellt ein neues L-System
func NewLSystem(axiom string, rules map[rune]string) *LSystem {
	return &LSystem{
		axiom:  axiom,
		rules:  rules,
		result: axiom,
	}
}

// Iterate führt eine bestimmte Anzahl von Iterationen des L-Systems durch
// Wert der letzten Iteration soll in ls.result
func (ls *LSystem) Iterate(n int) {
	//Hinweis: eine Möglichkeit ist es, strings.Builder zu verwenden:
	// var newResult strings.Builder
	// string schreiben:
	// func (b *strings.Builder) WriteString(s string) (int, error)
	// zeichen schreiben:
	// func (b *strings.Builder) WriteRune(r rune) (int, error)
	// ergebenis:
	// func (b *strings.Builder) String() string
	// newResult.WriteString(replacement)
	// newResult.WriteRune(char)
	// newResult.String()
	var newResult strings.Builder
	_ = newResult
}

// Result gibt das aktuelle Ergebnis des L-Systems zurück
func (ls *LSystem) Result() string {
	return ls.result
}

func main() {
	// Beispiel L-System
	axiom := "A"
	rules := map[rune]string{
		'A': "AB",
		'B': "A",
	}
	iterations := 4

	lsystem := NewLSystem(axiom, rules)
	lsystem.Iterate(iterations)

	fmt.Println("Ergebnis nach", iterations, "Iterationen:")
	fmt.Println(lsystem.Result())
}
