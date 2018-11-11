package mongo

import (
  "_internal/read_file"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Vet struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name    					  string
	Email    					  string
	Tel_1      					string
	Tel_2     				  string
	Latitude   					string
	Longitude    				string
	Address_street    	string
	Address_number   	 	string
	Zip_code   			 		string
}

func Fetch() map[int]map[string]string {

	vetsArray := make(map[int]map[string]string)
	vetsList := make([]Vet, 20);
	dbConfig := readFile.Read("/Users/valerio/go/src/_internal/mongo/config.json")

	session, err := mgo.Dial(dbConfig["address"].(string))
	if err != nil {	
		panic(err)
	}
	defer session.Close()

	// session.SetMode(mgo.Monotonic, true)

		// Drop Database
		// if IsDrop {
		// 	err = session.DB("vets").DropDatabase()
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }

	c := session.DB(dbConfig["name"].(string)).C("hotel_roma")

	err = c.Find(bson.M{}).All(&vetsList)
	if err != nil {	
		panic(err)
	}
	
	vetTemp := make(map[string]string)
	for i, vet := range vetsList {

		vetTemp["name"] = vet.Name
		vetTemp["email"] = vet.Email
		vetTemp["tel_1"] = vet.Tel_1
		vetTemp["tel_2"] = vet.Tel_2
		vetTemp["latitude"] = vet.Latitude
		vetTemp["longitude"] = vet.Longitude
		vetTemp["address_street"] = vet.Address_street
		vetTemp["address_number"] = vet.Address_number
		vetTemp["zip_code"] = vet.Zip_code

		vetsArray[i] = vetTemp;
	}

	return vetsArray;
}