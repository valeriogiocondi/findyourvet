package foo

import (
  "net/http"
  "projects/findyourvet/third_part/gorilla/mux"
)

func SayHello(w http.ResponseWriter, request *http.Request) { 

	vars := mux.Vars(request)
	name := vars["name"]

  // message := r.URL.Path
  // message = strings.TrimPrefix(message, "/")
  var message = "Hello " + name

  w.Write([]byte(message))
}