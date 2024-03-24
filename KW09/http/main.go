package main

import (
	"fmt"
	"net/http"
	"time"
)

/*Programm starten und im Browser http://localhost:5000/hello aufrufen
ggf. muss man noch eine Sicherheitsabfrage der Windows-Firewall/Virenscanner bestätigen
*/

func main() {
	//http.HandleFunc("/hello", hello)

	http.HandleFunc("/hello", protokolliereZeit(hello))
	http.ListenAndServe(":5000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
	time.Sleep(time.Millisecond * 100)
}
func protokolliereZeit(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hello(w, r)
		fmt.Printf("runtime: %v", time.Since(start))
		//TODO: rufe f auf und protokolliere die Zeit
	}
}

/*
	Wir wollen nun die Funktion hello "dekorieren",d.h.

Aufruf in main:

	http.HandleFunc("/hello", protokolliereZeit(hello))

func protokolliereZeit(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: rufe f auf und protokolliere die Zeit
	}
}
*/
