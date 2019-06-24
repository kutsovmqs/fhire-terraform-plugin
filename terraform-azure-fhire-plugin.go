package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "terraform-azure-fhire-plugin"
	app.Usage = "Create terraform azure fhire plugin"
	app.Action = func(c *cli.Context) error {
		accountName := os.Args[1]
		accessPolicies := os.Args[2:]
		fileName := "azuredeploy.parameters.json"

		var data DeployParemeters
		data = data.FillDeployParameters(&accountName, &accessPolicies)

		f2, _ := os.OpenFile(fileName, os.O_WRONLY, 0644)
		data.Write(f2)
		defer f2.Close()

		fmt.Println("done")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}