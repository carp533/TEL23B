package bubble

/* Bubblesort ist ein einfacher Sortieralgorithmus, der durch wiederholtes
Durchlaufen der Liste geht und benachbarte Elemente vertauscht, wenn sie
in der falschen Reihenfolge sind. Dieser Vorgang wird solange wiederholt,
bis die gesamte Liste sortiert ist.
Bubblesort hat eine quadratische Zeitkomplexität; in der Praxis werden oft
bessere Sortierverfahren eingesetzt.*/

/*Aufgabe: programmiere die Funktion Sort erneut, indem du für beide Schleifen
Zähler verwendest:
   for i:=0; i< len(array)-1; i++ {
      for j:=0; j < len(array)-i-1; j++ {
gib am Ende der äußeren Schleife das Array aus.
Führe den Algorithmus auf Papier aus und vergleiche das mit der Ausgabe deines
Programms.*/
func Sort(a []int) []int {
	var (
		n      = len(a)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return a
}

/*Die Binärsuche ist ein effizienter Suchalgorithmus, der auf einem sortierten
Array basiert.
Die Binärsuche hat eine logarithmische Laufzeitkomplexität von O(log n),
wodurch sie effizienter ist als lineare Suchalgorithmen, insbesondere für
große sortierte Datenmengen.

1. Initialisierung:
Die Suche beginnt mit einem sortierten Array. Es gibt zwei Indexe, low und high
(für Anfang und Ende des Bereichs).
Setze low auf den Index des ersten Elements und high auf den Index des letzten
Elements im Array.

2. Mittelpunkt bestimmen:
Berechne den Mittelpunkt des aktuellen Suchbereichs (mid = low + (high - low) / 2)

3. Vergleich:
Vergleiche das Element an der Position mid mit dem gesuchten Wert x.
Wenn a[mid] gleich x ist, wurde das Element gefunden, und die Suche ist erfolgreich.
	(return true)
Wenn a[mid] kleiner als x ist, liegt das gesuchte Element im rechten Teil des aktuellen
	Bereichs. Setze low = mid + 1 für den nächsten Durchlauf.
Wenn a[mid] größer als x ist, liegt das gesuchte Element im linken Teil des aktuellen
	Bereichs. Setze high = mid - 1 für den nächsten Durchlauf.

4. Wiederhole den Prozess:
Wiederhole die Schritte 2 und 3, bis low größer als high ist, in diesem Fall wurde das
Element nicht gefunden. (return false)

*/
/*Aufgabe: programmiere die BinarySearch anhand der obigen Beschreibung.
 */
func BinarySearch(a []int, x int) bool {
	return false
}
