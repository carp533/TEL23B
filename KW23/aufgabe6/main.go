package main

// Aufgabe: Prüfe auf balancierte Markup-Tags
// Ziel:
// Implementiere eine Funktion in Go, die prüft, ob eine gegebene Zeichenkette korrekt ausbalancierte Markup-Tags enthält. Die Zeichenkette kann Tags wie <open> und </open>, <start> und </start> enthalten.

// Beschreibung:
// Schreibe eine Funktion IsBalancedTags, die eine Zeichenkette als Eingabeparameter nimmt und einen booleschen Wert zurückgibt.
// Die Funktion soll true zurückgeben, wenn die Markup-Tags in der Zeichenkette korrekt ausbalanciert sind, andernfalls false.
// Ein Markup-Tag ist korrekt ausbalanciert, wenn jedes öffnende Tag ein entsprechendes schließendes Tag hat und die Tags richtig verschachtelt sind.
// Beispiel:
// Input:

// Zeichenkette: "<open><start></start></open>"
// Output:

// true
// Input:

// Zeichenkette: "<open><start></open></start>"
// Output:

// false
import (
	"fmt"
)

// IsBalancedTags prüft, ob die Markup-Tags in der Zeichenkette korrekt ausbalanciert sind
func IsBalancedTags(s string) bool {
	stack := []string{}
	tags := map[string]string{
		"</open>":  "<open>",
		"</start>": "<start>",
	}
	_, _ = stack, tags

	// suche '<'
	//     suche '>'
	// 	bei Fehlern: stop
	// 	lese Text zwischen < und >
	// 	prüfe ob es sich um ein schließendes tag handelt '</'; du kannst z.B. strings.HasPrefix(tag, "</") verwenden
	// 	  ja: prüfe, ob das oberste element des stack zum schließenden tag passt
	// 	  nein: füge das element zum stack hinzu
	// am ende muss stack leer sein

	return false
}

func main() {
	// Beispielzeichenketten
	examples := []string{
		"<open><start></start></open>",
		"<open><start></open></start>",
		"<open></open>",
		"<start><open></open></start>",
		"<start><open></start></open>",
		"<open><start></start>",
		"</open><start></start>",
	}

	for _, example := range examples {
		fmt.Printf("Zeichenkette: %s, Ausbalanciert: %v\n", example, IsBalancedTags(example))
	}
}
