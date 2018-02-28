package main

type Room struct {
	Id   int
	name string
	price float64
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

type HotelRoom struct{
	Id int
	Hotel Hotel
	Room Room
}