package main

import (
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
	

func loadRoutes(){
	http.Handle("/hotels", cached(config.Cachetime, hotels))
	http.Handle("/hotels-continent", cached(config.Cachetime, hotelsContinent))
	http.Handle("/hotels-room", cached(config.Cachetime, hotelsRooms))
	//let resource monitoring with prometheus
	http.Handle("/metrics", promhttp.Handler())
}