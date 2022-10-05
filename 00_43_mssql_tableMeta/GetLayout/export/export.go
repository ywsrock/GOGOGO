package export

import (
	"fmt"
	"getlayout/model"
	"html/template"
	"log"
	"os"
)

/**
*
* Scan result into Map
*
**/
func ScanInMap(results []map[string]interface{}) {
	for _, v := range results {
		for k, val := range v {
			fmt.Printf("%s:%s \n", k, fmt.Sprint(val))
		}
	}
}

/**
*
* Export by stdout
*
**/
func OutputStdout(results []model.Result) {
	for _, v := range results {
		fmt.Printf("%s - %s - %s \n", v.TableName, fmt.Sprint(v.ColumnName), v.MaxLength)
	}
}

/**
*
* Export by text template
*
**/
func OutputTextTemp(r *[]model.Result) {
	// output template
	tmpl := `
		TableName ---- ColumnName ---- DataType ---- MaxLength
		{{ range $k, $v := . }}
			{{ $v.TableName}} ---- {{ $v.ColumnName}} ---- {{ $v.DataType }} ---- {{$v.MaxLength}}
		{{ end }}
		`
	// template instance from template format tmpl
	t, _ := template.New("").Parse(tmpl)
	// parse template
	if err := t.Execute(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
