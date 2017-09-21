package main

import "os"
import "strings"
import "text/template"

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
`{{.FirstName}}
{{trim .FirstName}}
{{trim .FirstName | title}}
`

p := person{ FirstName : "     shamrock   "}
funcMap := template.FuncMap{
	"title": strings.Title,
	"trim": strings.TrimSpace,
}
tmpl := template.Must(template.New("table").Funcs(funcMap).Parse(source)) // HL
err := tmpl.Execute(os.Stdout, p)
	//END OMIT
	if err != nil {
		panic(err)
	}
}

