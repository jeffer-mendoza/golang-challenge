package main

import "net/http"

func loadRoutes(){
	http.Handle("/hotels", cached(config.CacheTime, hotels))
	http.Handle("/hotels-continent", cached(config.CacheTime, hotelsContinent))
}