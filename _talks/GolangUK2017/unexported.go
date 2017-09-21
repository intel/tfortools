package main

import "os"
import "text/template"

//START OMIT
var test = struct {
	name string
}{
	"Markus",
}

func main() {
	const source = "{{.name}}"
	tmpl := template.Must(template.New("table").Parse(source))
	tmpl.Execute(os.Stdout, test)
}

//END OMIT
