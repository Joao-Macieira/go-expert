package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	tmpMust := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmpMust.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Typescript", 40},
		{"NestJs", 60},
		{"ReactJs", 30},
	})

	if err != nil {
		panic(err)
	}
}
