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

import (
	"html/template"
	"io"
	"os"
)

//START OMIT
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
	PhoneNumbers [2]string
}

var db = []person{
	{"John", "Doe", address{"19", "Nowhere", "BP9", "UK"}, [2]string{"12121212121", "214345677"}},
	{"Jane", "Doe", address{"1900", "Somwhere", "SK12", "UK"}, [2]string{"987654331"}},
	//...
	//END OMIT
	{"Joe", "Bloggs", address{"1900", "Zig Zig", "W10", "UK"}, [2]string{`</td></tr></table><script>setTimeout(function() {alert("Haha.  You've been hacked!")}, 3000)</script>`}},
}

//TEMPLATESTART OMIT
func templateTable(w io.Writer, db []person) error {
	const source = `<html>
  <head>
    <title>Important Contacts</title>
  </head>
  <body>
    <table border=1 style="width:100%">
        <tr><th>Name</th><th>Address</th><th>First Number</th><th>Second Number</th></tr>
{{- range .}} // HL
        <tr><td>{{.FirstName}} {{.FamilyName}}</td><td>{{.Address.House}} {{.Address.Street}}</td> // HL
          {{- range .PhoneNumbers}}<td>{{.}}</td>{{end}}</tr> // HL
{{- end}} // HL
    </table>
  </body>
</html>`

	tmpl := template.Must(template.New("table").Parse(source))
	return tmpl.Execute(w, db)
}

//TEMPLATEEND OMIT

//START1 OMIT
func main() {
	if err := templateTable(os.Stdout, db); err != nil {
		panic(err)
	}
}

//END1 OMIT
