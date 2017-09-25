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
	"fmt"
	"io"
	"os"
)

//START OMIT
type address struct {
	House, Street, PostCode, Country string
}

type person struct {
	FirstName, FamilyName string
	Address               address
	PhoneNumbers          [2]string
}

var db = []person{
	{"John", "Doe", address{"19", "Nowhere", "BP9", "UK"}, [2]string{"12121212121", "214345677"}},
	{"Jane", "Doe", address{"1900", "Somewhere", "SK12", "UK"}, [2]string{"987654331", "172846281"}},
	{"Joe", "Bloggs", address{"1900", "Zig Zig", "W10", "UK"},
		[2]string{`</td></tr></table><script>setTimeout(function() {alert("Haha.  You've been hacked!")}, 3000)</script>`}},
	//END OMIT
}

//FPRINTSTART OMIT
func fprintTable(w io.Writer, db []person) error {
	const header = `<html>
  <head>
    <title>Important Contacts</title>
  </head>
  <body>
    <table border=1 style="width:100%%">
      <tr><th>Name</th><th>Address</th><th>First Number</th><th>Second Number</th></tr>
`
	if _, err := fmt.Fprintf(w, header); err != nil {
		return err
	}
	for _, p := range db {
		if _, err := fmt.Fprintf(w, "      <tr><td>%s %s</td>", p.FirstName, p.FamilyName); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, "<td>%s %s</td>", p.Address.House, p.Address.Street); err != nil {
			return err
		}
		for _, num := range p.PhoneNumbers {
			if _, err := fmt.Fprintf(w, "<td>%s</td>", num); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(w, "</tr>\n"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(w, "    </table>\n  </body>\n</html>\n"); err != nil {
		return err
	}

	return nil
}

//FPRINTEND OMIT

//START1 OMIT
func main() {
	if err := fprintTable(os.Stdout, db); err != nil {
		panic(err)
	}
}

//END1 OMIT
