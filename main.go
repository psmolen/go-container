package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(port(), nil))
}

func echo(w http.ResponseWriter, r *http.Request) {

	var o interface{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(b, &o)

	if err != nil {
		http.Error(w, "Error converting request to json",
			http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func port() string {
	var port string
	var found bool
	port, found = os.LookupEnv("PORT")
	if found {
		return fmt.Sprintf(":%s", port) // prepend a colon
	}
	return ":12345"
}