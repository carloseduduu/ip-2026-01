package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
    "servidorHTTP/app/utils" // Importa funções utilitárias, como criptografia e inserção no banco de dados
    "net/http"              // Usado para lidar com requisições e respostas HTTP
)

// FormHandler é responsável por processar os dados enviados pelo formulário de criação de conta
func CreatePaciente(response http.ResponseWriter, request *http.Request) {
    // Verifica se o método da requisição é POST
    if request.Method != http.MethodPost {
        // Retorna um erro caso o método não seja suportado
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores enviados pelo formulário
    username := request.FormValue("nome") // Nome do usuário
    cod_sus := request.FormValue("cod_sus")       // E-mail do usuário
    data_nascimento := request.FormValue("data_nascimento") // Data de nascimento do usuário
    sexo := request.FormValue("sexo") //Sexo do paciente
    tipo_sanguineo := request.FormValue("tipo_sanguineo") //Tipo sanguíneo
    
    // Insere os dados do usuário no banco de dados
    err := utils.Insert(cod_sus,username, data_nascimento, sexo, tipo_sanguineo)
    if err != nil {
        // Retorna um erro caso ocorra falha ao salvar os dados no banco de dados
        http.Error(response, "Erro ao salvar os dados no banco de dados", http.StatusInternalServerError)
        return
    }

    // Redireciona o usuário para a página inicial após o sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}