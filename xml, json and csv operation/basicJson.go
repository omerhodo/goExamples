package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
		{
			"data":{
				"object":"card",
				"id":"card_32423423432",
				"first_name":"omer",
				"last_name":"hodo",
				"balance":"34,112"
			}
		}
	`

	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("----------------")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
