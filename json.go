package graphql

import (
	"fmt"
)

type Json map[string]interface{}

func (_ Json) ImplementsGraphQLType(name string) bool {
	return name == "Json"
}

func (j *Json) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case Json:
		j = &input
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}
