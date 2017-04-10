package main
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Entertainment Data
}

type Data struct {
	Film Films
	Play Theater
}

type Films map[string]int
type Theater map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(response))

}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}

func getJsonResponse() ([]byte, error) {
	films := make(map[string]int)
	films["Little Big Man"] = 1970
	films["Bull Durham"] = 1988

	theater := make(map[string]int)
	theater["Hamlet"] = 1603
	theater["Death of a Salesman"] = 1949

	d := Data{films, theater}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
