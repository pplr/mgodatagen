package generators_test

import (
	"fmt"
	"log"

	"github.com/feliixx/mgodatagen/datagen/generators"
	"github.com/globalsign/mgo/bson"
)

func Example() {

	var content = map[string]generators.Config{
		"key": {
			Type:      generators.TypeString,
			MinLength: 3,
			MaxLength: 5,
		},
	}
	ci := generators.NewCollInfo(1, nil, 1, nil, nil)
	docGenerator, err := ci.DocumentGenerator(content)
	if err != nil {
		log.Fatal(err)
	}

	var doc struct {
		Key string `bson:"key"`
	}
	err = bson.Unmarshal(docGenerator.Generate(), &doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", doc)
	// Output: {Key:1jUK}

}