package HotelByCoords

import (
	"fmt"
  "net/http"
  "_internal/mongo"	
)
  // "fmt"
  // "gopkg.in/mgo.v2/bson"	

func Get(w http.ResponseWriter, r *http.Request) { 

	// mongo.Fetch()

	fmt.Fprintln(w, "Results All: ", mongo.Fetch())
}