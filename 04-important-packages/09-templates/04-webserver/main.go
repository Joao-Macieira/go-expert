package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpMust := template.Must(template.New("template.html").ParseFiles("template.html"))

		err := tmpMust.Execute(w, Courses{
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
