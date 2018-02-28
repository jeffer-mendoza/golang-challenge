package main

type Room struct {
	Name        string
	PricePerDay int
}

type Hotel struct {
	Id int
	Name     string
	Rooms []Room
}

type Continent struct {
	Id int
	Name     string
}

type HotelContinent struct {
	Id int
	Hotel Hotel
	Continent Continent
}