package main

import (
	"encoding/json"
)

const schema = "https://schema.management.azure.com/schemas/2015-01-01/deploymentParameters.json#"
const contentVersion = "1.0.0.0"

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

func (data *DeployParemeters) Fill(accountName *string, accessPolicies *[]string) DeployParemeters {
	accessPoliciesValues := make([]AccessPoliciesValue, len(*accessPolicies))

	for i := 0 ; i < len(*accessPolicies); i++ {
		accessPoliciesValues[i] = AccessPoliciesValue{ObjectId: (*accessPolicies)[i]}
	}

	return DeployParemeters{
		Schema: schema,
		ContentVersion: contentVersion,
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

func (deployParemeters *DeployParemeters) Marshal() ([]uint8, error) {
	return json.MarshalIndent(deployParemeters, "", " ")
}