package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"Go", 40}

	tmpMust := template.Must(template.New("CourseTemplate").Parse("Curso: {{.Name}} - Carga Horária {{.Workload}} horas"))

	err := tmpMust.Execute(os.Stdout, course)

	if err != nil {
		panic(err)
	}
}
