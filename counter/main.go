package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Type string
	Name string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "type", "", "Subtype for Counter keys")

	flag.Parse()

	d.Name = strings.Title(d.Type)

	t := template.Must(template.New("counter").Parse(counterTemplate))
	err := t.Execute(os.Stdout, d)
	if err != nil {
		log.Fatalf("failed to generate counter: %s", err)
	}
}
