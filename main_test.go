package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Gaardsholt/go-gitguardian/types"
	"github.com/getkin/kin-openapi/openapi2"
)

func Test_main(t *testing.T) {
	input, err := os.ReadFile("swagger.json")
	if err != nil {
		panic(err)
	}

	var doc openapi2.T
	if err = json.Unmarshal(input, &doc); err != nil {
		panic(err)
	}

	for k, v := range doc.Paths {
		operations := v.Operations()

		for o := range operations {
			t.Run(fmt.Sprintf("%s - %s", o, k), func(t *testing.T) {
				if !checkIfPathIsThere(k, o) {
					t.Fatalf("Path '%s' with operation '%s' is not there", k, o)
				}
			})
		}
	}

}

func checkIfPathIsThere(path string, operation string) bool {
	for _, v := range types.Endpoints {
		if v.GetApiPath() == path && v.Operation == operation {
			return true
		}
	}
	return false
}
