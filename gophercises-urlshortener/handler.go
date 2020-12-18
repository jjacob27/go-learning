package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		dest, keyPresent := pathsToUrls[req.URL.Path]
		if keyPresent {
			http.Redirect(w, req, dest, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, req)
	}
}

type yamlUrlMap struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var yamlUrls []yamlUrlMap
	err := yaml.Unmarshal(yml, &yamlUrls)
	mapUrls := make(map[string]string)
	if err == nil {
		for _, yaml := range yamlUrls {
			mapUrls[yaml.Path] = yaml.Url
		}
	}
	return MapHandler(mapUrls, fallback), err
}

func JsonHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var jsonUrls []yamlUrlMap
	err := json.Unmarshal(jsonBytes, &jsonUrls)
	mapUrls := make(map[string]string)
	if err == nil {
		for _, json := range jsonUrls {
			mapUrls[json.Path] = json.Url
		}
	}
	return MapHandler(mapUrls, fallback), err
}
