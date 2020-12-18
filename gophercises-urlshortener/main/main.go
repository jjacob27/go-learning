package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	urlshort ".."
)

func main() {
	yamlFile := flag.String("yaml", "urls.yaml", "Specify a file containing yaml of paths and URLS")
	jsonFile := flag.String("json", "urls.json", "Specify a file containing json of paths and URLS")

	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml, err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	json, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JsonHandler(json, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
