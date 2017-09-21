package main

import "text/template"
import "os"

type address struct {
	House    string
	Street   string
	PostCode string
	Country  string
}

type person struct {
	FirstName    string
	FamilyName   string
	Address      address
	PhoneNumbers []string
}

func main() {
	source :=
		//START OMIT
`{
  "phone_numbers" : [
  {{- range $i, $val := .PhoneNumbers -}}
   {{if ne $i 0 }},{{end -}} // HL
   "{{$val}}"
   {{- end}}]
}
`
p := person{PhoneNumbers: []string{"11111111111", "2222222222"}}
	//END OMIT
	tmpl := template.Must(template.New("table").Parse(source))
	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
