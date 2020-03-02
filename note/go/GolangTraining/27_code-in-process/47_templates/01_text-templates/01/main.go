package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//tpl, err := template.ParseFiles("tpl.gohtml")
	html := `E:\work\learn\go\git\GolangTraining\27_code-in-process\47_templates\01_text-templates\01\`+`tpl.gohtml`
	tpl, err := template.ParseFiles(html)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
