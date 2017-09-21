package main

import "text/template"
import "os"


type address struct {
	House string
	Street string
	PostCode string
	Country string
}

type person struct {
	FirstName string
	FamilyName string
	Address address
	PhoneNumbers []string
}

func main() {
	source :=
//START OMIT
`{{ "Who's this guy and where does he live" }}
{{.FirstName}}
{{with .Address -}}
  The {{ .Country }} is where {{$.FirstName}} lives.  {{/* . refers to Address instance */}}
{{end}}`

p := person{ FirstName: "Markus", Address : address{Country: "UK"}}
	//END OMIT
	tmpl := template.Must(template.New("table").Parse(source))
	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
