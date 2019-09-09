package main

import (
	"github.com/Jeffail/benthos/v3/lib/serverless/lambda"

	// Add your plugin packages here
	_ "github.com/benthosdev/benthos-plugin-example/cache"
	_ "github.com/benthosdev/benthos-plugin-example/condition"
	_ "github.com/benthosdev/benthos-plugin-example/input"
	_ "github.com/benthosdev/benthos-plugin-example/output"
	_ "github.com/benthosdev/benthos-plugin-example/processor"
)

//------------------------------------------------------------------------------

func main() {
	lambda.Run()
}

//------------------------------------------------------------------------------
