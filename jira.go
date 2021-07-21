package main

import (
	"github.com/andygrunwald/go-jira"
)

type Issue struct {
	Key       string `csv:"key"`
	Summary   string `csv:"Summary"`
	ParentKey string `csv:"ParentKey"`
	TypeName  string `csv:"TypeName"`
	EpicName  string `csv:"EpicName"`
	EpicKey   string `csv:"EpicKey"`
	Status    string `csv:"Status"`
}

func getIssues(username string, password string, url string, search string) (*[]Issue, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client, err := jira.NewClient(tp.Client(), url)
	if err != nil {
		return nil, err
	}
	opt := &jira.SearchOptions{
		MaxResults: 1000, // Max results can go up to 1000
	}
	chunk, _, err := client.Issue.Search(search, opt)

	if err != nil {
		return nil, err
	}
	var output = make([]Issue, len(chunk))
	for i, s := range chunk {

		output[i] = Issue{
			Key:      s.Key,
			Summary:  s.Fields.Summary,
			TypeName: s.Fields.Type.Name,
			Status:   s.Fields.Status.Name,
		}
		if s.Fields.Parent != nil {
			output[i].ParentKey = s.Fields.Parent.Key
		}
		if s.Fields.Epic != nil {
			output[i].EpicKey = s.Fields.Epic.Key
			output[i].EpicName = s.Fields.Epic.Key
		}

	}

	return &output, nil
}
