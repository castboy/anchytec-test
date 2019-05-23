package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{
    "tagA" : "json string",  
    "tagB" : 1024, 
    "tagC" : null, 
    "tagD" : {
        "tagE" : 2147483648, 
        "tagF" : 3.14159265358979
    },
    "tagG" : [
        "json array",
       1024,
        {"tagH" : "json object"}
    ]
}`

	var i interface{}
	err := json.Unmarshal([]byte(s), &i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	Assert(i)

}

func Assert(i interface{}) {
	switch t := i.(type) {
	case map[string]interface{}:
		for k, v := range t {
			switch t1 := v.(type) {
			case map[string]interface{}:
				fmt.Println(k, " : ")
				Assert(t1)
			case []interface{}:
				fmt.Println(k, " : ")
				for k1, v1 := range t1 {
					switch t2 := v1.(type) {
					case map[string]interface{}:
						fmt.Println(k1, " : ")
						Assert(t2)
					default:
						fmt.Println(k1, " : ", v1)
					}
				}
			default:
				fmt.Println(k, " : ", v)
			}
		}
	}
}
