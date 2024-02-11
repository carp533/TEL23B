Ein Interface I ist definiert durch eine Menge von Methodensignaturen.
Das Interface ist nicht direkt mit einem Datentypen gekoppelt.
Wir sagen: ein Datentyp T ist vom Typ Interface I, wenn T alle Methoden 
des Interfaces I bereit stellt.

In go ist es üblich, möglichst wenig Methoden in einem Interface zusammen zu
fassen.

typische Beispiele aus der go-Standard Bibliothek:

in go eingebaut: 
type error interface {
	Error() string
}

package sort: 
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
(anschaulich: wenn ich die 
 - Länge kenne, 
 - zwei Elemente vergleichen kann
 - zwei Elemente tauschen kann
 kann ich den Datentyp sortieren)

package io: 
(durch "kleine" interfaces fördert man das Schreiben modularer, wiederverwendbarer 
und flexibler Quelltexte.)
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}