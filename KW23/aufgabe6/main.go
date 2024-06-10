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
	"strings"
)

// IsBalancedTags prüft, ob die Markup-Tags in der Zeichenkette korrekt ausbalanciert sind
func IsBalancedTags(s string) bool {
	stack := []string{}
	tags := map[string]string{
		"</open>":  "<open>",
		"</start>": "<start>",
	}

	i := 0
	for i < len(s) {
		if s[i] == '<' {
			j := i
			for j < len(s) && s[j] != '>' {
				j++
			}
			if j == len(s) {
				return false // Ungültiges Tag-Format
			}
			tag := s[i : j+1]
			if strings.HasPrefix(tag, "</") {
				// Schließendes Tag
				if len(stack) == 0 || stack[len(stack)-1] != tags[tag] {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				// Öffnendes Tag
				stack = append(stack, tag)
			}
			i = j + 1
		} else {
			i++
		}
	}

	return len(stack) == 0
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
