package main

import (
	"fmt"

	"github.com/serverlessworkflow/sdk-go/v3/builder"
	"github.com/serverlessworkflow/sdk-go/v3/validate"
)

func main() {
	build()
	validExample()
}

func build() {
	fmt.Println("builder")

	workflowBuilder := builder.NewWorkflowBuilder()
	documentBuilder := workflowBuilder.Document()
	documentBuilder.SetName("test")
	documentBuilder.SetNamespace("test")
	documentBuilder.SetVersion("1.0.0")

	doBuilder := workflowBuilder.Do()
	callBuilder, _ := doBuilder.AddCall("test")
	callBuilder.SetCall("http")
	withBuilder := callBuilder.With()
	withBuilder.Set("method", "get")
	withBuilder.Set("endpoint", "https://petstore.swagger.io/v2/pet/{petId}")

	err := builder.Validate(workflowBuilder)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("json")
	data, _ := builder.Json(workflowBuilder)
	fmt.Println(string(data))
	fmt.Println("")

	fmt.Println("yaml")
	data, _ = builder.Yaml(workflowBuilder)
	fmt.Println(string(data))
}

func validExample() {
	fmt.Println("./example/example1.yaml")
	err := validate.FromFile("./example/example1.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}

	fmt.Println("")
	fmt.Println("./example/example2.yaml")
	err = validate.FromFile("./example/example2.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
