package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//ein Zeichen (einfache Anführungszeichen ') - Datentyp: rune (=int32)
	//Unicode code-point (ein bis vier byte Länge)
	//ch := '😀'

	ch := '\U0001F600' // \u 4 hex-Stellen \U 8 hex-Stellen
	fmt.Printf("Binär:  \t%b\n", ch)
	fmt.Printf("Unicode:\t%U\n", ch)
	fmt.Printf("Zeichen:\t%c\n", ch)
	fmt.Printf("Datentyp:\t%T\n", ch)
	fmt.Println("-------------------------------------")

	exampleString := "Hello, 世界" // Ein String mit ASCII und Nicht-ASCII Zeichen

	fmt.Println(exampleString)
	// Anzahl der Bytes
	byteLength := len(exampleString)
	fmt.Printf("Byte-Länge: %d\n", byteLength)

	// Anzahl der Runen (UTF-8 Zeichen)
	runeCount := utf8.RuneCountInString(exampleString)
	fmt.Printf("Rune-Länge: %d\n", runeCount)

	// Iterieren über jeden Rune (korrekt für UTF-8)
	fmt.Println("Iterieren über Runen:")
	for i, runeValue := range exampleString {
		fmt.Printf("%d:\t%c\n", i, runeValue)
	}

	fmt.Println("Iterieren über Bytes:")
	bytelist := []byte(exampleString)
	for i, byteValue := range bytelist {
		fmt.Printf("%d:\t%c\t%x\n", i, byteValue, byteValue)
	}

	// Extrahieren eines Substrings (korrekt für UTF-8)
	// Start bei Rune-Index 7 (beginnend mit 0), was dem Beginn des "世界" Teils entspricht
	runeIndexStart := 7
	runeIndexEnd := runeIndexStart + utf8.RuneCountInString(exampleString[runeIndexStart:])
	substring := string([]rune(exampleString)[runeIndexStart:runeIndexEnd])
	fmt.Printf("Substring: %s\n", substring)

}
