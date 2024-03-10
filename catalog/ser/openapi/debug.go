package openapi

import (
	"fmt"

	"github.com/jsightapi/jsight-api-core/catalog"
)

func debugExchangeSchema(es catalog.ExchangeSchema) {
	fmt.Printf("es type is %T\n", es)

	switch s := es.(type) {
	case *catalog.ExchangeJSightSchema:
		ast, err := s.JSchema.GetAST()
		if err != nil {
			fmt.Printf("No AST :(\n")
		} else {
			fmt.Printf("Value: %s \n", ast.Value)
		}
	default:
		fmt.Printf("Not a JSchema\n")
	}
	fmt.Printf("Debug Over\n")
}
