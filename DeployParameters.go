package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func (data *DeployParemeters) FillDeployParameters(accountName *string, accessPolicies *[]string) DeployParemeters {
	accessPoliciesValues := make([]AccessPoliciesValue, len(*accessPolicies))

	for i := 0 ; i < len(*accessPolicies); i++ {
		accessPoliciesValues[i] = AccessPoliciesValue{ObjectId: (*accessPolicies)[i]}
	}

	return DeployParemeters{
		Schema: "https://schema.management.azure.com/schemas/2015-01-01/deploymentParameters.json#",
		ContentVersion: "1.0.0.0",
		Parameters: Parameters{
			AccountName{
				Value: *accountName,
			},
			AccessPolicies{
				Value: accessPoliciesValues,
			},
		},
	}
}

func (data *DeployParemeters) Write(writer io.Writer) (int, error) {
	thread, _ := json.MarshalIndent(data, "", " ")
	if len(thread) == 0 {
		fmt.Println("Empty input parameters.")
		return 0, nil
	}
	writer.Write(thread)
	return len(thread), nil
}