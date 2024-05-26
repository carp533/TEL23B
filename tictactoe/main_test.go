package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPossibleMoves(t *testing.T) {
	//definiere das erwartete Ergebnis des Funktionsaufruf
	expected1 := [][]string{{PLAYER_MAX, EMPTY, EMPTY}, {EMPTY, EMPTY, EMPTY}, {EMPTY, EMPTY, EMPTY}}
	b := NewBoard()
	//rufe die Funktion auf
	moves, _ := getPossibleMoves(*b)
	fmt.Println(moves)
	found := false
	//vergeleiche das Ergebnis des Funktionsaufrufes mit dem erwarteten.
	for m := range moves {
		found = found || reflect.DeepEqual(moves[m], expected1)
	}
	if !found {
		t.Errorf("FEHLER!")
	}
}
