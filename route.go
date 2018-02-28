package main

import "net/http"

func loadRoutes(){
	http.Handle("/hotels", cached(config.Cachetime, hotels))
	http.Handle("/hotels-continent", cached(config.Cachetime, hotelsContinent))
}