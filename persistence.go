package main

import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
	"log"
)

/*
get database connection 
*/
func connection()(db *sql.DB, err error){
	db, err = sql.Open("postgres", config.getUrlConnection("postgres"))
	if err != nil {
		fmt.Println("error")
		db.Close()
	}
	return
}

func FindAllHotels() map[int][]Hotel{
	db, _ := connection()
	var (
		id int
		name string
	)
	rows, err := db.Query("select id, name from hotel")
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	hotels := make(map[int][]Hotel)
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		hotels[id] = append(hotels[id], Hotel{id, name, nil})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return hotels
}

func FindAllHotelsContinent() map[int][]HotelContinent{
	db, _ := connection()
	var (
		id int
		hotelId int
		hotel string
		continentId int
		continent string
	)
	rows, err := db.Query(`select hc.id, h.id as hotel_id, h.name as hotel,
		                   c.id as continent_id, c.name as continent
		                   from hotel_continent hc
						   INNER JOIN hotel h ON h.id = hc.hotel
						   INNER JOIN continent c on c.id = hc.continent`)
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	hotelsContinent := make(map[int][]HotelContinent)
	for rows.Next() {
		err := rows.Scan(&id, &hotelId, &hotel, &continentId, &continent)
		if err != nil {
			log.Fatal(err)
		}
		hotelsContinent[id] = append(hotelsContinent[id], HotelContinent{id,
			 Hotel{hotelId, hotel, nil}, 
			 Continent{continentId, continent}})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return hotelsContinent
}

func FindAllHotelsRoom() map[int][]HotelRoom{
	db, _ := connection()
	var (
		id int
		hotelId int
		hotel string
		roomId int
		room string
		price float64
	)
	rows, err := db.Query(`select hr.id, h.id, h.name, r.id,  r.name, r.price
		                   from hotel_room hr
		                   INNER JOIN hotel h ON h.id = hr.hotel
		                   INNER JOIN room r on r.id = hr.room`)
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	hotelsRoom := make(map[int][]HotelRoom)
	for rows.Next() {
		err := rows.Scan(&id, &hotelId, &hotel, &roomId, &room, &price)
		if err != nil {
			log.Fatal(err)
		}
		hotelsRoom[id] = append(hotelsRoom[id], HotelRoom{id,
			 Hotel{hotelId, hotel, nil}, 
			 Room{roomId, room, price}})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return hotelsRoom
}