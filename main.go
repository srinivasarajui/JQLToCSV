package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kaduartur/go-cli-spinner/pkg/spinner"
	"github.com/kaduartur/go-cli-spinner/pkg/template"
)

func main() {
	input := readFromCommandLine()
	err := getMissingData(&input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sp := spinner.New("Loading...")
	sp.SetTemplate(template.Arrow)
	sp.Start()

	issues, err := getIssues(input.username, input.password, input.url, input.search)
	if err != nil {
		sp.Error(errors.New("error to load"))
		fmt.Println(err)

		os.Exit(1)
	}
	// Print user input
	WriteToCsv(issues, input.targetFileName)

	if err != nil {
		sp.Error(errors.New("error to save"))
		fmt.Println(err)
		os.Exit(1)
	}
	sp.Success(fmt.Sprintf("Issues exported to %s\n", input.targetFileName))

}
