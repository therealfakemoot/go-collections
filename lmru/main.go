package main

import (
	"flag"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Type string
	Name string
	L    bool // whether this is an mru or lru
}

func main() {
	var d data
	flag.StringVar(&d.Type, "valueType", "bool", "valueType to store in cache")
	flag.BoolVar(&d.L, "lru", false, "least recently used?")

	flag.Parse()

	d.Name = strings.Title(d.Type)

	t := template.Must(template.New("lmru").Parse(lmruTemplate))
	t.Execute(os.Stdout, d)
}
