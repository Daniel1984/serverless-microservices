package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
	return "Lambda response", nil
}

func main() {
	fmt.Println("vim-go")
	lambda.Start(handler)
}
