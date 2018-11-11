package HotelByCity

import (
  "net/http"
  // "strings"
)

func Get(w http.ResponseWriter, r *http.Request) { 

  // message := r.URL.Path
  // message = strings.TrimPrefix(message, "/")
  // message = "Hello " + message

  w.Write([]byte("HotelByCity"))
}