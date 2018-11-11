package router

import (
  "gorilla/mux"
  "_components/HotelByCity"
  "_components/HotelByCoords"

)

func Router() *mux.Router {

  r := mux.NewRouter()
  r.HandleFunc("/hotelbycity/", HotelByCity.Get).Methods("GET")
  r.HandleFunc("/hotelbycoords/", HotelByCoords.Get).Methods("GET")
  return r
}