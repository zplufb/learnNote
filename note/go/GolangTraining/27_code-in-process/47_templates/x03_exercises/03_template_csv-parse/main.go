package main

import (

	"html/template"
	"log"
	"net/http"
	"github.com/GoesToEleven/GolangTraining/27_code-in-process/47_templates/x03_exercises/03_template_csv-parse/parse"
)

func main() {
	// parse csv
	records := parse.Parse("table.csv")

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// function
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// execute template
		err = tpl.Execute(res, records)
		if err != nil {
			log.Fatalln(err)
		}
	})

	// create server
	http.ListenAndServe(":9000", nil)
}
