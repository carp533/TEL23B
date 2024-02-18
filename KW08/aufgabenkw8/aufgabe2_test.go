package main

import (
	"fmt"
)

/*Aufgabe 2a: analysiere einen String auf doppelte Zeichen
Hinweis: caste den String in eine Liste mit rune:
chars := []rune(s)
du kannst
return fmt.Errorf(" %q is duplicated at positions %d and %d.\n", chars[pos1], pos1+1, pos2+1)
zurück geben, falls ein Zeichen doppelt ist.
*/

func Analyze(s string) error {
	return nil
}

func ExampleAnalyze() {
	//Output:
	//Analyze(""): ok
	//Analyze("abcABC"): ok
	//Analyze("XYZ ZYX"):  'X' is duplicated at positions 1 and 7.
	//Analyze("01234567890ABCDEFGHIJKLMN0PQRSTUVWXYZ0X"):  '0' is duplicated at positions 1 and 11.
	//Analyze("😍😀🙌💃😍🙌🐬🐳🐋🐡"):  '😍' is duplicated at positions 1 and 5.
	strings := []string{
		"",
		"abcABC",
		"XYZ ZYX",
		"01234567890ABCDEFGHIJKLMN0PQRSTUVWXYZ0X",
		"😍😀🙌💃😍🙌🐬🐳🐋🐡",
	}
	for _, s := range strings {
		fmt.Printf("Analyze(%q): ", s)
		if err := Analyze(s); err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("ok\n")
		}

	}
}

/*
Aufgabe 2b:
zähle die Wort-Häufigkeiten eines Strings. Als Rückgabe verwendest du eine map,
mit den Wörten als Schlüssel und Anzahl als Wert.
Hinweis: verwende die Funktion strings.Fields und strings.ToLower
*/
func WordFrequencyCounter(input string) map[string]int {
	return nil
}
func ExampleWordFrequencyCounter() {
	//Output:
	// Word Frequencies:
	// hello: 2
	// world: 2
	// this: 4
	// is: 2
	text := "Hello world this is this this hello this is world"
	frequencies := WordFrequencyCounter(text)

	fmt.Println("Word Frequencies:")
	for word, freq := range frequencies {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
