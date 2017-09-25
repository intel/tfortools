//
// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
