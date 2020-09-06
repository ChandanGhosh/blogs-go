package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Pretty(value interface{}) {
	b, err := json.MarshalIndent(value, "", " ")
	if err != nil {
		log.Printf("Marshaling error, %v", err)
		return
	}
	fmt.Println(string(b))
}
