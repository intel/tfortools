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

