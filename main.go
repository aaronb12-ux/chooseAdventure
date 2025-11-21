package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
)

type Options struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

type StoryArc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Options `json:"options"`
}

type MyHandler struct {
	message string
}



func unmarshalJSON() (map[string]StoryArc, error) {
	
	file, e := os.ReadFile("gopher.json")

	if e != nil {
		log.Fatal("error opening file")
	}

	var data map[string]StoryArc

	err := json.Unmarshal(file, &data)

    return data, err
	
}

func  (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from my handler: %s", h.message)
}

func main() {

	unmarshaledJSON, err := unmarshalJSON()
	

	if err != nil {
		log.Fatal("error unmarshaling JSON")
	}


	http.Handle("/", introHandler)

	http.ListenAndServe(":8080", nil)

}
