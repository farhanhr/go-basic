package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0283","name":"Thinkplus T-80", "price" : 180000}`
	jsonByte := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonByte, &result)
	
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id": "P0283",
		"name": "Thinkplus T-80",
		"price": 180000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}