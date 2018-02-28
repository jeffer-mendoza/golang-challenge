package main

import "net/http"

func loadRoutes(){
	http.Handle("/hotels", cached("10s", hotels))
	http.HandleFunc("/hotels-continent", hotelsContinent)
}