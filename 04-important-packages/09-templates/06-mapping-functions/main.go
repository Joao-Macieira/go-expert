package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := []string{
			"header.html",
			"content.html",
			"footer.html",
		}

		tmp := template.New("content.html")
		tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
		tmp = template.Must(tmp.ParseFiles(templates...))

		err := tmp.Execute(w, Courses{
			{"Go", 40},
			{"Typescript", 40},
			{"NestJs", 60},
			{"ReactJs", 30},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}
