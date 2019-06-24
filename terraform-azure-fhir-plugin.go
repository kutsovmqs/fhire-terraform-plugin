package main

import (
	"os"
)

func main() {
	accountName    := os.Args[1]
	accessPolicies := os.Args[2:]

	fileName := "azuredeploy.parameters.json"

	var data DeployParemeters
	data = data.FillDeployParameters(&accountName, &accessPolicies)

	file, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)

	data.Write(file)
	defer file.Close()
}