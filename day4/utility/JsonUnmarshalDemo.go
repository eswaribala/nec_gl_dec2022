package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	users, _ := ioutil.ReadFile("data.json")
	var result []map[string]interface{}

	json.Unmarshal([]byte(users), &result)
	for _, value := range result {

		for innerKey, innerValue := range value {

			if reflect.TypeOf(innerValue).String() == "map[string]interface {}" {
				fmt.Printf("%+v", innerValue)

			} else {
				fmt.Printf("%s->%s\n", innerKey, innerValue)
			}
		}
	}
}
