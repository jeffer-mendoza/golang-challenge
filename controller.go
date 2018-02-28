package main

import (
	"encoding/json"
	"net/http"
)


/**
get hotels
**/
func hotels(response http.ResponseWriter, request *http.Request) {
	//get post request data
	//start_date := request.FormValue("start_date")
	//due_date := request.FormValue("due_date")
	hoteles := FindAllHotels()
	result, err := json.Marshal(hoteles)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(result)
}

/**
get hotels with continent
**/
func hotelsContinent(response http.ResponseWriter, request *http.Request){
	hoteles := FindAllHotelsContinent()
	result, err := json.Marshal(hoteles)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(result)
}