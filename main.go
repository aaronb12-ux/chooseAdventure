package main

import (
	//"encoding/json"
	"encoding/json"

	"log"
	"net/http"
	"os"

	"html/template"
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
	UnmarshaledJSON map[string]StoryArc
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
	
	var path string
	tmpl := template.Must(template.ParseFiles("page.html"))
	if r.URL.Path == "/" {
	
		path = "intro"

		data := StoryArc{
			Title:  h.UnmarshaledJSON[path].Title,
			Story: h.UnmarshaledJSON[path].Story,
			Options: h.UnmarshaledJSON[path].Options,
		}

		tmpl.Execute(w, data)
	} else {
	
		path :=  r.URL.Path[1:] 

		data := StoryArc{
			Title:  h.UnmarshaledJSON[path].Title,
			Story: h.UnmarshaledJSON[path].Story,
			Options: h.UnmarshaledJSON[path].Options,
		}

		tmpl.Execute(w, data)
	}
}

func main() {

	unmarshaledJSON, err := unmarshalJSON()
	
	if err != nil {
		log.Fatal("error unmarshaling JSON")
	}

	handler := &MyHandler{UnmarshaledJSON: unmarshaledJSON}
	http.Handle("/", handler)


	http.ListenAndServe(":8080", nil)

}
