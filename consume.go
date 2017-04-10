package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main(){
	url := "http://localhost:1337"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Entertainment.Film, "\n", p.Entertainment.Play)

 }
