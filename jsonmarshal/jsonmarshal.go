package main 

import (  
    "encoding/json"
    "fmt"
    "bytes"
)

/* -----------------------------------------
Definition of the internal Book object (read from input)
--------------------------------------------*/
type Book struct {
	Id    					json.Number			`json:"id"`
	Revision				int					`json:"revision"`
	ISBN					Everstring			`json:"isbn"`
	Title					string      		`json:"title"`
}

/* -----------------------------------------
Definition of the public Book object
--------------------------------------------*/
type AliasBook Book
type omit 			*struct{}

type PublicBook struct {
    Id 			string 			`json:"id"`
    Revision 	omit 			`json:"revision,omitempty"`
    *AliasBook
}

/* -----------------------------------------
Rendering functions
--------------------------------------------*/
func (bb *Book) MarshalJSON() ([]byte, error) {
	fmt.Println("---------------MarschalJSON---------------")
	aux := PublicBook{
		Id:			bb.Id.String(),
		AliasBook:	(*AliasBook)(bb),
	}
	
	return json.Marshal(&aux)
}

type Everstring int64
func (isbn *Everstring) UnmarshalJSON(b []byte) error {
	var n json.Number
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	val, _ := n.Int64()
	*isbn = Everstring(val)
	return nil
}

func (isbn *Everstring) MarshalJSON() ([]byte, error) {
	s := fmt.Sprint(*isbn)
	return json.Marshal(s)
}

func main() {
	var jsonStreams[2][]byte
	// Input ISBN as string
	jsonStreams[0] = []byte(`{"id":"123","revision":1234,"isbn":"9","title":"Go for dummies"}`)
	// Input ISBN as int
	jsonStreams[1] = []byte(`{"id":123,"revision":1234,"isbn":9,"title":"Go for dummies"}`)
	
	// For each stream
	for i := range jsonStreams {
		fmt.Print("stream: ")
		fmt.Println(string(jsonStreams[i]))
	
		// Read Input
		b := Book{}
		err := json.Unmarshal(jsonStreams[i], &b)
		if err == nil {
	        fmt.Printf("%+v\n", b)
	    } else {
	        fmt.Println(err)
	        fmt.Printf("%+v\n", b)
	    }
	    
	    // Output as JSON
		response := new(bytes.Buffer)
		enc := json.NewEncoder(response)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "    ")
		err = enc.Encode(&b)
		if err == nil {
	        fmt.Printf("%+v\n", response)
	    } else {
	        fmt.Println(err)
	        fmt.Printf("%+v\n", response)
	    }
	}
}