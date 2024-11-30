package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type DSA struct {
	Generic     string     `json:"generic"`
	TypeGeneric string     `json:"typeGeneric"`
	Ty          string     `json:"type"`
	Methods     []Method   `json:"methods"`
	Properties  []Property `json:"properties"`
	Fn          string     `json:"fn"`
	Args        string     `json:"args"`
	Ret         string     `json:"return"`
}
type Function struct {
	Name string
	Ty   string `json:"type"`
	Fn   string `json:"fn"`
	Args string `json:"args"`
	Ret  string `json:"return"`
}

type Structure struct {
	Name        string
	Letter      string
	Generic     string     `json:"generic"`
	TypeGeneric string     `json:"typeGeneric"`
	Ty          string     `json:"type"`
	Methods     []Method   `json:"methods"`
	Properties  []Property `json:"properties"`
}

type Method struct {
	Name string `json:"name"`
	Args string `json:"args"`
	Ret  string `json:"return"`
}

type Property struct {
	Name  string `json:"name"`
	Ty    string `json:"type"`
	Scope string `json:"scope"`
}

//go:embed data.json
var content string

func main() {
	var types map[string]DSA
	err := json.Unmarshal([]byte(content), &types)
	if err != nil {
		fmt.Println("sucks to suck", err)
		return
	}

	structs := []Structure{}
	fns := []Function{}
	for key, val := range types {

		switch val.Ty {
		case "struct":
			v := Structure{
				Name:        key,
				TypeGeneric: val.TypeGeneric,
				Letter:      string(strings.ToLower(key)[0]),
				Generic:     val.Generic,
				Ty:          val.Ty,
				Methods:     val.Methods,
				Properties:  val.Properties,
			}
			structs = append(structs, v)
		case "fn":
			v := Function{
				Name: key,
				Ty:   val.Ty,
				Fn:   val.Fn,
				Args: val.Args,
				Ret:  val.Args,
			}
			fns = append(fns, v)
		}
	}

	var structTmpl = `package katamachine
// Code is generated but you can edit it to get the tests passing
{{if .Generic }}
import (
	"cmp"
)
{{end}}
{{ $StructName := .Name }}
{{ $StructLetter := .Letter }}
{{ $TypeGeneric := .TypeGeneric }}

type {{.Name}}{{.Generic}} struct {
{{- range $key, $val := .Properties}}
	{{ $val.Name }} {{$val.Ty}}
{{- end}}
}
{{ printf "\n" }}
{{- range $key, $val := .Methods}}
func ({{ $StructLetter }} *{{$StructName}}{{$TypeGeneric}}) {{$val.Name}}({{$val.Args}}) {{if .Ret}}({{.Ret}}){{end}}{
	{{if .Ret }}return nil{{end}}
}
{{ printf "\n" }}
{{- end }}`
	template := template.Must(template.New("").Parse(structTmpl))

	for _, st := range structs {

		fileName := strings.ToLower(st.Name) + "_gen.go"
		if _, err := os.Stat(fileName); err == nil {
			fmt.Println("file", fileName, "already exists. Skipping...")
			continue
		}
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		template.Execute(file, st)
	}
}
