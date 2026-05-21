package handlers

import (
	"fmt"
	"net/http"
	"servidorHTTP/app/utils"
	//"text/template"
)

func UpdatePacienteHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores do formulário
	cod_sus := request.FormValue("cod_sus")
	nome := request.FormValue("nome")
	data_nascimento := request.FormValue("bornDate")
	sexo := request.FormValue("sexo")
	tipo_sanguineo := request.FormValue("tipo_sanguineo")

	fmt.Println(cod_sus)

	// Verifica se a senha atual está correta
	isValidUser, err := utils.ValidateUser(cod_sus)
	if err != nil || !isValidUser {
		http.Error(response, "Número do SUS inválido!(Não pode ser alterado!)", http.StatusUnauthorized)
		return
	}

	// Cria um mapa para armazenar os campos a serem atualizados
	updates := make(map[string]string)

	if nome != "" {
		updates["nome"] = nome
	}
	if data_nascimento != "" {
		updates["data_nascimento"] = data_nascimento
	}
	if sexo != "" {
		updates["sexo"] = sexo
	}
	if tipo_sanguineo != "" {
		updates["tipo_sanguineo"] = tipo_sanguineo //utils.Encrypt(newPassword) // Criptografa a nova senha
	}

	// Atualiza os campos informados no banco de dados
	err = utils.UpdatePaciente(cod_sus, updates)
	if err != nil {
		http.Error(response, "Erro ao atualizar os dados no banco de dados", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	/*
		var num_sus string

		if email == "" {
			consultationEmail = currentEmail
		} else {
			consultationEmail = email
		}

		user, err := utils.GetUserByEmail(consultationEmail)
		if err != nil {
			http.Error(response, "Erro ao buscar informações do usuário", http.StatusInternalServerError)
			return
		}

		// Carrega o template do perfil
		tmpl, err := template.ParseFiles("static/profile.html")
		if err != nil {
			http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
			return
		}

		// Renderiza o template com os dados do usuário
		err = tmpl.Execute(response, user)
		if err != nil {
			http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
			return
		}
	*/
	// Redireciona para a página de perfil ou outra página de sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
