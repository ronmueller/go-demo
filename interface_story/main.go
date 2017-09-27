package main

import (
	"encoding/json"
	"fmt"
)

type animals []interface{}

type animal struct {
	Type string
	Name string
}

type zooMain struct {
	Name string `json:"name"`
}

type zooAlias zooInput

type zooInput struct {
	zooMain
	RawResult []json.RawMessage `json:"animals,omitempty"`
	Animals   animals           ``
}

/*
func (z *zooInput) UnmarshalJSON([]byte) error {

    return nil
}
*/

type zooOutput struct {
	zooMain
	RawResult []json.RawMessage `json:"-"`
	Animals   animals           `json:"animals"`
}

type dog struct {
	Type string
	Name string
	Legs int
}

type fish struct {
	Type string
	Name string
}

func main() {

	z := zooInput{}

	j := `{ "name": "zoo", "animals":[ { "type": "dog", "name": "foo", "legs": 4}, { "type": "fish", "name": "bar"} ] }`

	json.Unmarshal([]byte(j), &z)

	for _, ar := range z.RawResult {
		a := animal{}
		json.Unmarshal(ar, &a)
		switch a.Type {
		case "dog":
			d := dog{}
			json.Unmarshal(ar, &d)
			z.Animals = append(z.Animals, d)
		case "fish":
			f := fish{}
			json.Unmarshal(ar, &f)
			z.Animals = append(z.Animals, f)
		}
	}

	z.RawResult = nil

	fmt.Printf("%#v\n", z)

	b, _ := json.Marshal(zooOutput(z))
	fmt.Println(string(b))
}
