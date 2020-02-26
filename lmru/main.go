package main

import (
	"bytes"
	"flag"
	"go/format"
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
	flag.StringVar(&d.Type, "type", "bool", "value type to store in cache")
	flag.StringVar(&d.Package, "pacakge", "main", "package name")
	flag.BoolVar(&d.L, "lru", false, "least recently used?")

	flag.Parse()

	d.Name = strings.Title(d.Type)

	d.Name = "MRU" + d.Name
	if d.L {
		d.Name = "LRU" + d.Name
	}

	t := template.Must(template.New("lmru").Parse(lmruTemplate))

	var b bytes.Buffer
	err := t.Execute(&b, d)
	if err != nil {
		log.Fatalf("error generating type: %s", err)
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated code: %s", err)
	}

	err = os.Stdout.Write(formatted)
	if err != nil {
		log.Fatalf("error writing generated code: %s", err)
	}
}
