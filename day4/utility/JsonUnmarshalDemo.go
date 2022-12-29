package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	users, _ := ioutil.ReadFile("data.json")
	var result []map[string]interface{}

	json.Unmarshal([]byte(users), &result)
    
}
