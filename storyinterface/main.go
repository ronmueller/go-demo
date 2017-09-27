package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Story struct; reading from couchbase
type Story struct {
	ID              json.Number      `json:"id"`
	Title           string           `json:"title"`
	ArticleElements []ArticleElement `json:"article_elements"`
}

type ArticleElement struct {
	BoxType string
	ID      json.Number
}

type VideoAd struct {
	BoxType string
}

type Video struct {
	ID     int
	CBID   string
	ShowAd bool
}

func (v *Video) getdata() {
	fmt.Println("videointerface")
}

type Diashow struct {
	ID      int
	CBID    string
	AdIndex int
}

func (v Diashow) getdata() {
	fmt.Println("diashowinterface")
}

type cbinterface interface {
	getdata()
}

// AliasArticleElement ArticleElement
type AliasArticleElement ArticleElement

func (ae *ArticleElement) UnmarshalJSON(b []byte) error {
	fmt.Println("unmarshal articleelements")
	var element AliasArticleElement
	json.Unmarshal(b, &element)
	switch bt := element.BoxType; bt {
	case "video_ad":
		*ae = VideoAd(element)
	}
	return nil
}

// AliasStory Story
type AliasStory Story

// PublicStory struct; Definition of the public story object
type PublicStory struct {
	//    Revision 	omit			`json:"revision,omitempty"`
	*AliasStory
}

// MarshalJSON Story;
func (s *Story) MarshalJSON() ([]byte, error) {
	aux := PublicStory{
		AliasStory: (*AliasStory)(s),
	}
	return json.Marshal(&aux)
}

func main() {
	var jsonStreams [1][]byte
	jsonStreams[0] = []byte(`{"id":123,"title":"Tadaaa","article_elements":[{"boxtype":"video_ad"},{"boxtype":"video", "id":123},{"boxtype":"dia", "id":456}]}`)
	// jsonStreams[1] = []byte(`{"id":123,"title":"Go for dummies"}`)

	// For each stream
	for i := range jsonStreams {
		fmt.Print("stream: ")
		fmt.Println(string(jsonStreams[i]))

		// Read Input
		b := Story{}
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
