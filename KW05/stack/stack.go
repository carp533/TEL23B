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
	top    *node
	length int
}
type node struct {
	value interface{}
	prev  *node
}

/*In go ist es üblich, sich für die Initialisierung eine eigene Funktion zu
schreiben.*/
func New() *Stack {
	return &Stack{top: nil}
}

/*möchte man den Stack auf der Konsole ausgeben, muss man hierfür eine eigene
(rekursive) Methode programmiern, andernfalls wird nur die Adresse von top
ausgegeben. (z.B. {0xc00004e250})
Die Methode muss Signatur "String() string" haben, dann wird diese automatisch
bei fmt.Print* aufgerufen.*/

func (s Stack) String() string {
	if s.top == nil {
		return "empty stack!"
	}
	return fmt.Sprintf("Stack, Länge: %d %v.", s.length, s.top.String())
}
func (n node) String() string {
	if n.prev == nil {
		return fmt.Sprintf("%v -> nil", n.value)
	}
	return fmt.Sprintf("%v ->", n.value) + n.prev.String()
}

/*Aufgabe: a) teste das Programm, einmal mit und einmal ohne obige String Methode
b) programmiere die rekursive Methode String für die Ausgabe der Knoten. Die Abbruch
Bedingung ist hier prev == nil und rufe diese in der String() Methode für Stack auf.*/

func (s *Stack) Push(val interface{}) {
	top := &node{val, s.top}
	s.top = top
	s.length++
	//s.length += 1
}

func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return 0, fmt.Errorf("can't Pop() with empty stack!")
	}
	val := s.top.value
	s.top = s.top.prev
	s.length--
	return val, nil
}

func (s Stack) IsEmpty() bool {
	return s.length == 0
}

func (s Stack) Top() interface{} {
	return s.top.value
}

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

func liste(li []int) {
	li[0] = 99
	li = append(li, 6)
}

func main() {

	fmt.Println("Stack demo!")
	s := New()
	fmt.Println(s)
	s.Push("hello hello!")
	for i := 1; i < 5; i++ {
		s.Push(i * 10)
		fmt.Println(*s)
	}
	//for !s.IsEmpty() {
	val, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	fmt.Println(*s)

}
