package main

import (
  "net/http"
  "_internal/router"
)

func main() {

	var r = router.Router()
 
  // http.HandleFunc("/", foo.SayHello)
  if err := http.ListenAndServe(":8081", r); err != nil {
    panic(err)
  }

}