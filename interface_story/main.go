package main

import (
	"encoding/json"
	"fmt"
	"net/url"
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
	*zooMain
	RawResult []json.RawMessage `json:"animals,omitempty"`
	Animals   animals           ``
}

type zooOutput struct {
	*zooMain
	RawResult []json.RawMessage `json:"-"`
	Animals   animals           `json:"animals"`
}

type WebURLType string

type dog struct {
	Type           string
	Name           string
	Legs           int
	WebURL         WebURLType
	WebPageContent string
}

type fish struct {
	Type           string
	Name           string
	WebURL         WebURLType
	WebPageContent string
}

func main() {

	z := zooInput{}

	j := `{ "name": "zoo", "animals":[ { "type": "dog", "name": "foo", "legs": 4, "weburl": "http://www.dog.com"}, { "type": "fish", "name": "bar", "weburl": "www.fish.com"} ] }`

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

	fmt.Printf("%#v\n\n", z)

	marsh := zooOutput(z)
	b, _ := json.Marshal(&marsh)

	fmt.Println(string(b))
}

func (w WebURLType) MarshalJSON() ([]byte, error) {
	s := fmt.Sprint(w)
	fmt.Printf("%s: %s\n", "WebURL", s)

	_, err := url.ParseRequestURI(s)
	if err != nil {
		// delete if URL is not valid
		s = ""
	}

	return json.Marshal(&s)
}
