package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func searchInObject(obj interface{}, searchString string) interface{} {
	for _, value := range obj.(map[string]interface{}) {
		if reflect.TypeOf(value).String() == "string" {
			if strings.Contains(strings.ToLower(value.(string)), strings.ToLower(searchString)) {
				return obj
			}
		}

		if reflect.TypeOf(value).String() == "map[string]interface {}" {
			searchResult := searchInObject(value, searchString)
			if searchResult != nil {
				return searchResult
			}
		}
	}
	return nil
}

func search(obj []interface{}, searchString string) []interface{} {
	var searchResults []interface{}
	for _, item := range obj {
		searchResult := searchInObject(item, searchString)
		if searchResult != nil {
			searchResults = append(searchResults, searchResult)
		}
	}
	return searchResults
}

func main() {
	file, err := os.Open("../../../data/movies.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)

	var datas []interface{}

	err = decoder.Decode(&datas)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	searchResults := search(datas, "Science Fiction")
	elapsed := time.Since(start)

	resultsJSON, err := json.MarshalIndent(searchResults, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resultsJSON))
	fmt.Printf("Found %d results in %s\n", len(searchResults), elapsed)

}
