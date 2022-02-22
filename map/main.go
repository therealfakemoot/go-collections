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
	Key     string // key type
	Value   string // value type
	Name    string // name of the generated type, e.g. StringIntOrderedMap
	Package string
}

func main() {
	var d data
	flag.StringVar(&d.Key, "key", "bool", "value type to store in cache")
	flag.StringVar(&d.Value, "value", "bool", "value type to store in cache")
	flag.StringVar(&d.Package, "package", "main", "package name")
	flag.StringVar(&d.Name, "name", "", "name of the generated type, e.g. StringIntOrderedMap")
	// flag.BoolVar(&d.L, "lru", false, "least recently used?")

	flag.Parse()

	d.Name = "Ordered" + strings.Title(d.Key+d.Value) + "Map"

	t := template.Must(template.New("orderedmap").Parse(orderedMapTemplate))

	var b bytes.Buffer
	err := t.Execute(&b, d)
	if err != nil {
		log.Fatalf("error generating type: %s", err)
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated code: %s", err)
	}

	_, err = os.Stdout.Write(formatted)
	if err != nil {
		log.Fatalf("error writing generated code: %s", err)
	}
}
