package handlers

import (
	"html/template"
	"log"
	"net/http"
	"servidorHTTP/app/utils"
)

// Struct que representa os dados do seu banco de dados
type Pacientes struct {
	Cod_sus         string
	Nome            string
	Data_nascimento string
	Sexo            string
	Tipo_sanguineo  string
}

func ListPacientes(w http.ResponseWriter, r *http.Request) {
	// Simulação de dados vindos do banco de dados
	lista, err := utils.DB.Query("SELECT cod_sus, nome, data_nascimento, sexo, tipo_sanguineo FROM pacientes")
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes no banco", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	defer lista.Close()
	//Obtive os dados do banco de dados em lista, agora criar uma slice que guarde tudo organizado

	var listaPacientes []Pacientes

	for lista.Next() {
		var p Pacientes
		err := lista.Scan(&p.Cod_sus, &p.Nome, &p.Data_nascimento, &p.Sexo, &p.Tipo_sanguineo)
		if err != nil {
			http.Error(w, "Erro ao ler bando de dados (listPacienteHandler)", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		listaPacientes = append(listaPacientes, p)
	}

	// Carrega o arquivo HTML (ou string)
	tmpl, err := template.ParseFiles("static/forms/listPaciente.html")
	if err != nil {
		http.Error(w, "Erro ao carregar o template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Executa o template passando os dados
	err = tmpl.Execute(w, listaPacientes)
	if err != nil {
		http.Error(w, "Erro ao renderizar a página", http.StatusSeeOther)
		return
	}
}
