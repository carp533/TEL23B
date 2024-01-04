# Arrays
Der Typ [n]T ist ein Array mit n Werten des Typs T.

Der Ausdruck

var a [10]int

deklariert eine Variable a als ein Array von zehn int. Die Indizes der Elemente sind 
0,1,...8,9

Die Länge eines Arrays ist Teil seines Typs, so dass Arrays nicht in ihrer Größe verändert 
werden können. Dies scheint einschränkend zu sein, aber keine Sorge, Go bietet eine bequeme 
Möglichkeit, mit Arrays zu arbeiten.

# Slices
Ein Array hat eine feste Größe. Ein Slice hingegen ist eine dynamische, flexible Sicht auf 
die Elemente eines Arrays. In der Praxis sind Slices viel häufiger anzutreffen als Arrays.

Der Typ []T ist ein Slice mit Elementen des Typs T.

Ein Slice wird durch die Angabe von zwei Indizes gebildet, einer unteren und einer oberen 
Grenze, die durch einen Doppelpunkt getrennt sind:

a[low : high]
Damit wird ein halboffener Bereich ausgewählt, der das erste Element einschließt, aber das 
letzte ausschließt.

Der folgende Ausdruck erzeugt ein Slice, das die Elemente 1 bis 3 von a einschließt:

a[1:4]

## Slices sind Referenzen auf Arrays

Ein Slice speichert keine Daten, sondern beschreibt lediglich einen Abschnitt eines zugrunde 
liegenden Arrays.

Werden die Elemente eines Slices geändert, so werden die entsprechenden Elemente des zugrunde 
liegenden Arrays geändert.

Andere Slices, die dasselbe zugrunde liegende Array nutzen, sehen diese Änderungen.