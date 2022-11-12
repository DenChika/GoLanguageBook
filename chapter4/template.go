package main

import (
	"chapter4/github"
	"html/template"
	"log"
	"os"
)

/*var reportText = template.Must(template.New("Issuelisttext").
Funcs(template.FuncMap{"DaysAgo": github.DaysAgo}).
Parse(github.TemplText))*/

var reportHTML = template.Must(template.New("Issuelisthtml").
	Parse(github.TemplHTML))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	/*if err := reportText.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}*/
	if err := reportHTML.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
