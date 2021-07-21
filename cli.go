package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/manifoldco/promptui"
)

const usage = `Description: Get Jira Issues to a csv file.

Usage: %s [options]

Options:
`

type inputData struct {
	username       string
	password       string
	url            string
	search         string
	targetFileName string
}

func validateURL(input string) error {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("Invalid URL")
	}
	return nil
}

func validateLength(size int) promptui.ValidateFunc {
	return func(input string) error {
		if len(input) <= size {
			return fmt.Errorf("The input should be minium of %d", size)
		}
		return nil
	}
}

// promptPlain prompts user for an input that is echo-ed on terminal.
func prompt(question string, validate promptui.ValidateFunc, hideInput bool) (string, error) {

	prompt := promptui.Prompt{
		Label:    question,
		Validate: validate,
	}
	if hideInput {
		prompt.Mask = '*'
	}
	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

func getMissingData(input *inputData) error {
	if len(input.url) < 1 {
		url, err := prompt("What is your JIRA url?", validateURL, false)
		if err != nil {
			return err
		}
		input.url = url
	}
	if len(input.username) < 1 {
		username, err := prompt("What is your JIRA username?", validateLength(1), false)
		if err != nil {
			return err
		}
		input.username = username
	}
	if len(input.password) < 1 {
		password, err := prompt("What is your JIRA password?", validateLength(1), true)
		if err != nil {
			return err
		}
		input.password = password
	}
	if len(input.search) < 1 {
		search, err := prompt("What is your JQL?", validateLength(1), false)
		if err != nil {
			return err
		}
		input.search = search
	}
	if len(input.targetFileName) < 1 {
		targetFileName, err := prompt("What is your output filename ?", validateLength(1), false)
		if err != nil {
			return err
		}
		input.targetFileName = targetFileName
	}
	return nil
}

func readFromCommandLine() inputData {
	var input inputData
	// Collect user input
	flag.StringVar(&input.url, "url", input.url, "The Jira URL")
	flag.StringVar(&input.username, "username", input.username, "The username of Jira")
	flag.StringVar(&input.password, "password", input.password, "The password of Jira")
	flag.StringVar(&input.search, "search", input.search, "The JQL to search")
	flag.StringVar(&input.targetFileName, "filename", input.targetFileName, "The csv file name where details to be exported")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	return input
}
