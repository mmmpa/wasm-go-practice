package main

import (
	"os"
	"encoding/json"
	"github.com/mmmpa/iso-go/data"
)

func main() {
	var message data.Message
	err := json.Unmarshal([]byte(os.Args[1]), &message)

	if err != nil {
		j, _ := json.Marshal(
			data.ValidationResult{
				Valid:   false,
				Message: "Invalid JSON.",
			},
		)
		println(string(j))
		return
	}

	j, _ := json.Marshal(data.Validate(message))
	println(string(j))
}
