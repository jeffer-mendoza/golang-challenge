package main

import "net/http"

func loadRoutes(){
	http.Handle("/hotels", cached("10m", hotels))
	http.Handle("/hotels-continent", cached("10m",hotelsContinent))
}