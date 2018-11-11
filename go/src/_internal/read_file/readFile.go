package readFile
	
import (
  "fmt"
  "encoding/json"
  "os"
  "io/ioutil"
 )

func Read(path string) map[string]interface{} {

	var dbConfig map[string]interface{}

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(byteValue), &dbConfig)

	// JsonInterface2ArrayString(dbConfig)

	return dbConfig;
}

/*
func JsonInterface2ArrayString(jsonArray map[string]interface{}) map[string]interface{} {

	res := make(map[string]interface{})

	fmt.Println(res);

	return res;
}
func FetchMapInterface(jsonArray map[string]interface{}, res map[string]interface{}) map[string]interface{} {

	for k, i := range jsonArray {

		switch nestedArray := i.(type) {

			case string: {
				
				res[k] = i
				fmt.Println(k)
				FetchMapInterface(nestedArray, res)
			}
			case interface{}: {
				
				// CREATE A NEW MAP
				fmt.Println(nestedArray)
				FetchMapInterface(nestedArray, make(map[string]string))
			}
			case []interface{}: {
				
				// CREATE A NEW MAP OF MAP
				fmt.Println(nestedArray)
				FetchMapInterface(nestedArray, make(map[int]interface{}))
			}
			default: {}
		}
	}

	return res

}
*/