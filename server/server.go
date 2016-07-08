package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Mapping dsda
type Mapping struct {
	URI  string `json:"uri" yaml:"uri"`
	File string `json:"file" yaml:"file"`
}

// Configuration sdad
type Configuration struct {
	Port     int       `json:"port" yaml:"port"`
	Mappings []Mapping `json:"mappings" yaml:"mappings"`
}

var configuration *Configuration

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	uri := fmt.Sprintf("%v", r.URL.Path)
	//fmt.Printf("URI = %v\n", uri)

	for _, mapping := range configuration.Mappings {
		if strings.Compare(mapping.URI, uri) == 0 {
			data, err := ioutil.ReadFile(mapping.File)
			if err != nil {
				fmt.Printf("URI = %v -> ERROR %v \n", uri, http.StatusInternalServerError)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				fmt.Printf("URI = %v -> %v %v \n", uri, mapping.File, http.StatusOK)
				http.ServeContent(w, r, mapping.File, time.Now(), bytes.NewReader(data))
				return
			}
		}
	}
	// If Here not found
	fmt.Printf("URI = %v -> ERROR %v \n", uri, http.StatusNotFound)
	http.Error(w, "No content for "+uri, http.StatusNotFound)
}

func Start(config *Configuration) {
	configuration = config
	http.HandleFunc("/", downloadHandler) // set router
	port := fmt.Sprintf(":%v", configuration.Port)
	err := http.ListenAndServe(port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
