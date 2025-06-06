package main

import (
	"fmt"
	"os"

	"example.com/m/v2/openfga"
	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/rego"
	"github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
)

func mainInner() (*client.OpenFgaClient, error) {
	apiUrl := os.Getenv("FGA_API_URL")
	if apiUrl == "" {
		apiUrl = "http://localhost:8080"
	}
	creds := credentials.Credentials{}
	fgaClient, err := client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:               apiUrl,
		StoreId:              os.Getenv("FGA_STORE_ID"), // not needed when calling `CreateStore` or `ListStores`
		AuthorizationModelId: os.Getenv("FGA_MODEL_ID"), // optional, recommended to be set for production
		Credentials:          &creds,
	})
	if err != nil {
		return nil, err
	}

	return fgaClient, nil
}

func main() {
	fgaClient, err := mainInner()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	rego.RegisterBuiltin1(openfga.RegisterCheck("openfga.check", fgaClient))
	rego.RegisterBuiltin1(openfga.RegisterBatchCheck("openfga.batchcheck", fgaClient))
	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
