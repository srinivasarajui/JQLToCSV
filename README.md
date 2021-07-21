# JQLToCSV

## for runing in development

go run \*.go

## Building for windows

env GOOS=windows GOARCH=amd64 go build -o JqlToCsv.exe .

JqlToCsv.exe -url http://jira.ultria.net:8080/ -username srini.raju@ultria.com -password password -filename sample.csv -search "project = Ultria-CLM "
