package main

import (
	"os"
	"text/template"
)

type Message struct {
	Name   string
	Host   string
	Time   string
	Intact bool
}

func main() {
	messages := []Message{
		{
			Name:   "Johannes",
			Host:   "lachs-buttern.jetzt",
			Time:   "in 10 Sekunden",
			Intact: true,
		},
		{
			Name:   "Ruben",
			Host:   "lack-saufen.jetzt",
			Time:   "vor 5 Minuten",
			Intact: true,
		},
	}
	var tmplFile = "messages.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, messages)
	if err != nil {
		panic(err)
	}
}
