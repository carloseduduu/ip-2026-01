package handlers

import (
	"log"
	"net/http"
	"servidorHTTP/app/utils"
)

func DeletePaciente(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		log.Println("Método não suportado DELETEACCOUNTHANDLER")
		return
	}

	cod_sus := request.FormValue("cod_sus")
	nome := request.FormValue("nome")

	// Verifica se o usuário existe no banco de dados
	isValidUser, err := utils.ValidateDelete(cod_sus, nome)
	if err != nil || !isValidUser {
		http.Error(response, "Nome ou Número SUS inválido", http.StatusUnauthorized)
		return
	}

	// Remove o usuário do banco de dados
	err = utils.DeleteUser(cod_sus, nome)
	if err != nil {
		http.Error(response, "Erro ao apagar a conta", http.StatusInternalServerError)
		return
	}

	// Redireciona para a página inicial após o sucesso
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
