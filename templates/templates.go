package templates

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func RunTemplates() {
	curso := Curso{"GO", 40}

	tmp := template.New("CursoTemplate")

	tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}")

	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}

func RunTemplateMust() {
	curso := Curso{"GO", 40}

	tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))

	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}

type Cursos []Curso

func RunTemplateWithExternalFile() {
	tmp := template.Must(template.New("template.html").ParseFiles("./templates/template.html"))

	err := tmp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 20},
	})

	if err != nil {
		panic(err)
	}
}

func RunTemplateWithWebServer() {
	http.HandleFunc("/courses", func(w http.ResponseWriter, request *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("./templates/template.html"))

		err := tmp.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 60},
			{"Python", 20},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}

func RunTemplateWithCompoundTemplates() {
	templates := []string{
		"./templates/header.html",
		"./templates/content.html",
		"./templates/footer.html",
	}

	tmp := template.Must(template.New("content.html").ParseFiles(templates...))

	err := tmp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 20},
	})

	if err != nil {
		panic(err)
	}
}

func RunTemplateWithMappedFunctions() {
	templates := []string{
		"./templates/header.html",
		"./templates/content.html",
		"./templates/footer.html",
	}

	tmp := template.New("content.html")

	tmp.Funcs(template.FuncMap{
		"Upping": strings.ToUpper,
		"MyFunc": MyFunc,
	})

	tmp = template.Must(tmp.ParseFiles(templates...))

	err := tmp.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 60},
		{"Python", 20},
	})

	if err != nil {
		panic(err)
	}
}

func MyFunc(value int) int {
	return value * 5
}
