package main

import (
	"fmt"
)

/**********************************************************************
* Aufgabe A (10 Punkte):
* AUFGABENSTELLUNG:
Implementiere eine Funktion in Go, die eine Liste von Funktionswerten
(Floats) glättet, indem sie das arithmetische Mittel über ein Fenster der
Länge 𝑘 berechnet.

Beschreibung:
Schreibe eine Funktion Smooth, die eine Liste von Gleitkommazahlen und eine Fensterlänge
𝑘
k als Eingabeparameter nimmt.
Die Funktion soll die Liste glätten, indem sie für jedes Element das arithmetische Mittel der
𝑘
k benachbarten Elemente (inklusive des Elements selbst) berechnet.
Gib die geglättete Liste zurück.
Beispiel:
Input:

Liste: [1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0]
Fensterlänge: 3
Output:

Glättete Liste: [1.5, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 9.5]
**********************************************************************/

// Smooth glättet eine Liste von Floats mit einem gleitenden Durchschnitt
func Smooth(values []float64, k int) []float64 {
	n := len(values)
	if n == 0 || k <= 0 {
		return nil
	}

	// Initialisiere die geglättete Liste
	smoothed := make([]float64, n)

	// Gleite über die Liste und berechne das arithmetische Mittel für jedes Fenster
	for i := 0; i < n; i++ {
		start := max(0, i-k/2)
		end := min(n, i+k/2+1)
		sum := 0.0
		for j := start; j < end; j++ {
			sum += values[j]
		}
		smoothed[i] = sum / float64(end-start)
	}

	return smoothed
}

// Hilfsfunktionen für min und max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Beispiel-Liste von Funktionswerten
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	k := 3

	smoothedValues := Smooth(values, k)

	fmt.Println("Originalwerte:", values)
	fmt.Println("Geglättete Werte:", smoothedValues)
}
