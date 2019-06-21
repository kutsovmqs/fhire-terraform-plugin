package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DeployParemeters struct {
	Schema string `json:"$schema"`
	ContentVersion string `json:"contentVersion"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	AccountName `json:"accountName"`
	AccessPolicies `json:"accessPolicies"`
}

type AccountName struct {
	Value string `json:"value"`
}

type AccessPolicies struct {
	Value [] AccessPoliciesValue `json:"value"`
}

type AccessPoliciesValue struct {
	ObjectId string `json:"objectId"`
}

func main() {
	accountName := os.Args[1]
	accessPolicies := make([]AccessPoliciesValue, len(os.Args[2:]))

	for i := 0 ; i < len(os.Args[2:]); i++ {
		accessPolicies[i] = AccessPoliciesValue{ObjectId:os.Args[i+2]}
	}

	data := DeployParemeters {
		Schema:			"https://schema.management.azure.com/schemas/2015-01-01/deploymentParameters.json#",
		ContentVersion: "1.0.0.0",
		Parameters: Parameters{
			AccountName{
				Value: accountName,
			},
			AccessPolicies{
				Value:accessPolicies,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("azuredeploy.parameters.json", file, 0644)
}