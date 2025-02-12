package main

import (
	"net/http"
	"os"
	"text/template"
	"time"
)

type Pessoa struct {
	Nome           string
	Curso          string
	DataNascimento time.Time
}

type Pessoas []Pessoa

func main() {
	matheus := Pessoa{
		"Matheus José",
		"Sistemas de informação",
		time.Date(1995, time.April, 4, 0, 0, 0, 0, time.UTC),
	}

	pamella := Pessoa{
		"Pâmella Christina",
		"Medicina",
		time.Date(1998, time.December, 18, 0, 0, 0, 0, time.UTC),
	}

	// Fazendo toda a operação, separando o Nome do template do valor a ser parseado
	pessoaTemplate := template.New("PessoaTemplate")

	html, err := pessoaTemplate.Parse("O(A) usuário(a) {{.Nome}}, possui {{.RetornaIdade}} anos, cursa(ou) {{.Curso}} \n")

	if err != nil {
		panic(err)
	}

	err = html.Execute(os.Stdout, matheus)

	if err != nil {
		panic(err)
	}

	err = html.Execute(os.Stdout, pamella)

	if err != nil {
		panic(err)
	}

	// Fazendo toda a operação de forma automática
	must := template.Must(template.New("PersonTemplate").Parse("The person {{.Nome}}, has {{ .RetornaIdade }} years old and study {{.Curso}} \n"))

	err = must.Execute(os.Stdout, matheus)

	if err != nil {
		panic(err)
	}

	err = must.Execute(os.Stdout, pamella)

	if err != nil {
		panic(err)
	}

	// Fazendo o parse de um arquivo html
	pessoas := Pessoas{
		matheus,
		pamella,
	}
	html = template.Must(template.New("templateFile.html").ParseFiles("templateFile.html"))

	err = html.Execute(os.Stdout, pessoas)

	if err != nil {
		panic(err)
	}

	// Fazendo através de um http server

	http.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		pessoas := Pessoas{
			matheus,
			pamella,
		}
		http := template.Must(template.New("templateFile.html").ParseFiles("templateFile.html"))
		err = http.Execute(writter, pessoas)

		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)

}

// Função que calcula a idade
func (p Pessoa) RetornaIdade() int {
	var anoAtual = time.Now()
	const horasNoAno = 8760
	return int(anoAtual.Sub(p.DataNascimento).Hours() / horasNoAno)
}
