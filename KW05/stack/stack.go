package main

import "fmt"

/*Versuche die Aufgaben zu lösen, OHNE in die Musterlösung zu schauen.
Die Fehlenden Funktionen werden im Labor gemeinsam programmiert.*/
/*Aufgabe:
(1) Vergleiche den Datentyp Stack mit dem Datentyp verkettete Liste.
(2) Programmiere den Stack mit
  a) der Methode Length ->Push +1, Pop -1
  b) allgemeinen Daten statt int -> interface{}. Teste das, indem du Push()
	 mit verschiedenen Datentypen aufrufst.
*/

type Stack struct {
	top *node
}
type node struct {
	value int /* oder interface{} */
	prev  *node
}

/*In go ist es üblich, sich für die Initialisierung eine eigene Funktion zu
schreiben.*/
//func New() *Stack {}

/*möchte man den Stack auf der Konsole ausgeben, muss man hierfür eine eigene
(rekursive) Methode programmiern, andernfalls wird nur die Adresse von top
ausgegeben. (z.B. {0xc00004e250})
Die Methode muss Signatur "String() string" haben, dann wird diese automatisch
bei fmt.Print* aufgerufen.*/

//func (s Stack) String() string {}

/*Aufgabe: a) teste das Programm, einmal mit und einmal ohne obige String Methode
b) programmiere die rekursive Methode String für die Ausgabe der Knoten. Die Abbruch
Bedingung ist hier prev == nil und rufe diese in der String() Methode für Stack auf.*/

//func (n node) String() string {}

//func (s *Stack) Push(val int) {}

//func (s *Stack) Pop() int {}

//func (s Stack) IsEmpty() bool {}

//func (s Stack) Top() int {}

/*Aufgabe: programmiere die Methode Pop(). Diese soll das oberste Element lesen,
ohne den Stack zu verändern.*/

/*Bemerkung: ruft man Pop() auf einem leeren Stack, bekommt man einen Fehler
"runtime error: invalid memory address or nil pointer dereference"
Man muss also immer zuerst IsEmpty() prüfen.
eine Alternative wäre, bei Pop() zwei Werte zurück zu geben (den Wert und einen Fehler):
func (s *Stack) Pop() (int, error)
der Aufruf wäre dann:
value, err := s.Pop()
if err != nil {
	//Fehlerbehandlung
}
*/
func main() {
	fmt.Println("Stack demo!")
	/*s := New()
	fmt.Println(*s)
	for i := 1; i < 5; i++ {
		s.Push(i * 10)
		fmt.Println(*s)
	}
	for !s.IsEmpty() {
		fmt.Println("Pop(): ", s.Pop())
	}
	fmt.Println(*s)*/
}
