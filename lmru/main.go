package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

type data struct {
	Type    string
	Name    string
	L       bool // whether this is an mru or lru
	Package string
}

func main() {
	var d data
	flag.StringVar(&d.Type, "valueType", "bool", "valueType to store in cache")
	flag.StringVar(&d.Package, "pacakge", "main", "package name")
	flag.BoolVar(&d.L, "lru", false, "least recently used?")

	flag.Parse()

	d.Name = strings.Title(d.Type)

	d.Name = "MRU" + d.Name
	if d.L {
		d.Name = "LRU" + d.Name
	}

	t := template.Must(template.New("lmru").Parse(lmruTemplate))
	err := t.Execute(os.Stdout, d)
	if err != nil {
		log.Fatalf("error generating type: %s", err)
	}
}
