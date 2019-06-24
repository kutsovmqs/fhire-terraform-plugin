package main

import (
	"io/ioutil"
	"os"
)

const outputFileName = "azuredeploy.parameters.json"
const accountNameInputIndex = 1
const accessPoliciesInputIndex = 2

func main() {
	accountName    := os.Args[accountNameInputIndex]
	accessPolicies := os.Args[accessPoliciesInputIndex:]
	fileName 	   := outputFileName

	var deployParemeters DeployParemeters
	deployParemeters = deployParemeters.Fill(&accountName, &accessPolicies)

	data, _ := deployParemeters.Marshal()

	_ = ioutil.WriteFile(fileName, data, 0644)
}