package main

import (
	"encoding/json"
	"fmt"

	"github.com/pluralsh/polly/template"
	_ "github.com/pluralsh/polly/template"
)

func main() {
	marshal, err := json.MarshalIndent(template.RegisteredFilters(), "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}
