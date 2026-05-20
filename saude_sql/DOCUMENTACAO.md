# Documentação do Sistema de Saúde (SUS Digital)

Esta documentação descreve o funcionamento do sistema atual baseado em terminal (CLI) desenvolvido em **Go (Golang)** integrado a um banco de dados **PostgreSQL**, explica detalhadamente as funções existentes e apresenta um guia passo a passo para reaproveitar este código na criação de um **site web moderno** (com uma API backend em Go e um frontend em HTML/JavaScript).

---

## 1. Visão Geral do Sistema Atual

Atualmente, o projeto é um aplicativo executado via terminal que realiza operações de **CRUD** (Create, Read, Update, Delete) de pacientes no banco de dados. Ele permite:

1. **Cadastrar Pacientes**: Salvar o número do cartão SUS, nome, data de nascimento, sexo e tipo sanguíneo.
2. **Listar Pacientes**: Visualizar todos os registros ordenados pelo código SUS.
3. **Editar Dados**: Atualizar campos específicos de um paciente no banco de dados utilizando seu código SUS como chave.
4. **Excluir Usuários**: Remover o registro de um paciente a partir do código SUS.

O sistema divide a lógica de manipulação do banco de dados (no pacote `funcao`) da interface de usuário em linha de comando (no arquivo `main.go` da pasta `Backend`).

---

## 2. Estrutura do Banco de Dados

A tabela no PostgreSQL é estruturada conforme o arquivo [pacientes.sql](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/ip-2026-01/saude_sql/database/pacientes.sql):

```sql
CREATE TABLE IF NOT EXISTS pacientes (
    cod_SUS VARCHAR(20) PRIMARY KEY,
    nome VARCHAR(100),
    data_nascimento DATE,
    sexo VARCHAR(10),
    tipo_sanguineo VARCHAR(10) CHECK (tipo_sanguineo IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'))
);
```

### Detalhes dos Campos:
*   `cod_SUS`: Identificador único do paciente (chave primária) de até 20 caracteres.
*   `nome`: Nome completo do paciente (até 100 caracteres).
*   `data_nascimento`: Armazenado no tipo `DATE` do banco de dados.
*   `sexo`: String simples ("Masculino", "Feminino", "Não informado").
*   `tipo_sanguineo`: Armazena a tipagem com validação `CHECK` para aceitar apenas os 8 tipos válidos.

---

## 3. Explicação das Funções ([funcoes/funcao.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/ip-2026-01/saude_sql/funcoes/funcao.go))

O arquivo de funções centraliza toda a comunicação direta com o banco de dados. A string de conexão atual está configurada como constante:
```go
const server = "host=localhost port=5432 user=postgres password=1234 dbname=saude sslmode=disable"
```

Abaixo está a análise detalhada de cada função:

### Struct `Paciente`
```go
type Paciente struct {
    Cod_SUS, Nome, Data_nascimento, Sexo, Tipo_sanguineo string
}
```
*   **Propósito**: Representa a estrutura de dados de um paciente na memória do Go. Os campos iniciam com letra maiúscula para que sejam públicos (exportados), permitindo que outros pacotes (como o `main` ou pacotes de conversão para JSON) os acessem.

---

### `Connect()`
*   **Assinatura**: `func Connect() (string, *sql.DB, error)`
*   **O que faz**:
    1. Abre uma conexão com o banco de dados PostgreSQL usando o driver `"postgres"` (`github.com/lib/pq`).
    2. Testa a conexão usando `db.Ping()`. Se houver falha na rede ou credenciais incorretas, dispara um `panic`.
    3. Retorna uma mensagem de status de sucesso, o ponteiro de conexão `*sql.DB` e o erro (se houver).
*   **Para o Site Web**: Essencial para conectar a API com o banco de dados ao iniciar o servidor.

---

### `Insert()`
*   **Assinatura**: `func Insert(tabela, cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo string) string`
*   **O que faz**:
    1. Abre uma nova conexão com o banco.
    2. Valida o status com `db.Ping()`.
    3. Prepara um comando SQL dinâmico (`INSERT INTO %s ...`) usando `Prepare` e executa usando `Exec`.
    4. Usa parâmetros seguros `$1, $2, $3, $4, $5` para evitar **SQL Injection**.
    5. Retorna uma mensagem de sucesso em formato string.
*   **Para o Site Web**: Será acionada quando o usuário preencher um formulário de cadastro no site e clicar em "Enviar".

---

### `Create_table()`
*   **Assinatura**: `func Create_table(arquivo string)`
*   **O que faz**:
    1. Abre conexão com o banco de dados.
    2. Lê o arquivo de script SQL localizado no caminho fixo `./database/pacientes.sql` (ignora o parâmetro `arquivo`).
    3. Prepara e executa o script no banco para criar a tabela caso ela não exista.
*   **Para o Site Web**: Útil para ser executada apenas uma vez durante a inicialização/deploy do servidor para garantir que a tabela existe no banco.

---

### `List_pacientes()`
*   **Assinatura**: `func List_pacientes() ([]Paciente, error)`
*   **O que faz**:
    1. Abre conexão com o banco.
    2. Executa a query `SELECT * FROM pacientes ORDER BY cod_sus`.
    3. Itera através das linhas retornadas (`list.Next()`), lê as colunas através do `list.Scan()` mapeando para a struct `Paciente`.
    4. Adiciona cada paciente a uma lista do tipo `[]Paciente` e a retorna.
*   **Para o Site Web**: Será usada para enviar a lista de pacientes cadastrados em formato JSON para renderização em uma tabela ou grid na interface web.

---

### `Update()`
*   **Assinatura**: `func Update(tabela, sus, coluna, dado string)`
*   **O que faz**:
    1. Abre conexão com o banco.
    2. Monta uma instrução SQL parametrizada especificando qual coluna será atualizada com base no `cod_sus`.
    3. Executa a query atualizando o valor.
*   **Para o Site Web**: Será usada para o recurso de edição rápida de dados no frontend.

---

### `Delete()`
*   **Assinatura**: `func Delete(tabela, cod_sus string) string`
*   **O que faz**:
    1. Abre conexão com o banco de dados.
    2. Executa um comando `DELETE FROM %s WHERE cod_sus = $1`.
    3. Exclui o registro e retorna uma mensagem de confirmação.
*   **Para o Site Web**: Será acionada por um botão de "Excluir" (geralmente com ícone de lixeira) ao lado do nome do paciente na tabela.

---

## 4. Como Reaproveitar Este Código Para Um Site Web

Para migrar de um aplicativo de terminal para um site web, a arquitetura deve mudar de **CLI (Execução local via Console)** para **Cliente-Servidor (Web API)**:

```
[ Navegador Web ] (Frontend: HTML, CSS, JavaScript)
        │ 
        │ Requisições HTTP (GET, POST, PUT, DELETE) com dados em JSON
        ▼
[ Servidor API Go ] (Backend: main.go adaptado para HTTP)
        │
        │ Comunicação SQL (usando funcao.go)
        ▼
[ Banco de Dados PostgreSQL ]
```

### O que muda no código?
1. **Remover interações de console**: O terminal (`bufio.Scanner`, `fmt.Scanln`, menus de loop `for { switch }`) deve ser substituído por rotas HTTP.
2. **Formato de dados**: O backend receberá e responderá em formato **JSON** em vez de texto puro formatado para o console.
3. **CORS (Cross-Origin Resource Sharing)**: Habilitar cabeçalhos CORS no Go para permitir que o frontend hospedado em outra porta ou domínio consiga consumir a API.

---

## 5. Implementação Prática (Exemplo Completo de Backend e Frontend)

Para reaproveitar suas funções e expor os dados para a web, você pode criar uma API HTTP. Aqui está como fazer usando a biblioteca padrão do Go (`net/http`) de forma simples, direta e sem frameworks externos:

### Backend: Criando a API HTTP (`Backend/main_web.go`)

Crie este arquivo (ou modifique o `main.go` existente) para rodar o servidor HTTP:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	funcao "saude_sql/funcoes"
)

// Estrutura para receber dados de atualização ou criação
type RequestData struct {
	CodSUS         string `json:"cod_sus"`
	Nome           string `json:"nome"`
	DataNascimento string `json:"data_nascimento"` // DD-MM-AAAA ou AAAA-MM-DD
	Sexo           string `json:"sexo"`
	TipoSanguineo  string `json:"tipo_sanguineo"`
	Coluna         string `json:"coluna"`
	Dado           string `json:"dado"`
}

func main() {
	// 1. Conecta ou cria a tabela no início
	_, _, err := funcao.Connect()
	if err != nil {
		fmt.Println("Erro ao conectar com o banco:", err)
		return
	}
	funcao.Create_table("pacientes") // Garante que a tabela existe
	fmt.Println("Banco conectado e tabela verificada.")

	// 2. Define as rotas HTTP da API
	http.HandleFunc("/api/pacientes", handlePacientes)

	// 3. Inicia o servidor HTTP na porta 8080
	fmt.Println("Servidor HTTP rodando na porta 8080 (http://localhost:8080)...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao rodar servidor:", err)
	}
}

// Configura cabeçalhos CORS para permitir chamadas do frontend
func setupCORS(w *http.ResponseWriter, req *http.Request) bool {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	
	if req.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return true
	}
	return false
}

func handlePacientes(w http.ResponseWriter, r *http.Request) {
	if setupCORS(&w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// LISTAR PACIENTES (GET http://localhost:8080/api/pacientes)
		lista, err := funcao.List_pacientes()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Falha ao listar pacientes"})
			return
		}
		json.NewEncoder(w).Encode(lista)

	case http.MethodPost:
		// CADASTRAR PACIENTE (POST http://localhost:8080/api/pacientes)
		var p RequestData
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Dados inválidos"})
			return
		}

		// Reutiliza a função Insert do seu pacote
		msg := funcao.Insert("pacientes", p.CodSUS, p.Nome, p.DataNascimento, p.Sexo, p.TipoSanguineo)
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"mensagem": msg})

	case http.MethodPut:
		// ATUALIZAR PACIENTE (PUT http://localhost:8080/api/pacientes)
		var p RequestData
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Dados inválidos"})
			return
		}

		// Reutiliza a função Update do seu pacote
		funcao.Update("pacientes", p.CodSUS, p.Coluna, p.Dado)
		
		json.NewEncoder(w).Encode(map[string]string{"mensagem": "Dado atualizado com sucesso!"})

	case http.MethodDelete:
		// DELETAR PACIENTE (DELETE http://localhost:8080/api/pacientes?cod_sus=XXXX)
		codSUS := r.URL.Query().Get("cod_sus")
		if codSUS == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"erro": "Código SUS é obrigatório"})
			return
		}

		// Reutiliza a função Delete do seu pacote
		msg := funcao.Delete("pacientes", codSUS)
		
		json.NewEncoder(w).Encode(map[string]string{"mensagem": msg})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
```

---

### Frontend: Criando a Interface Web (`index.html`)

Este é um exemplo de interface moderna que consome a API acima em tempo real. Ele contém um formulário de cadastro e uma tabela que exibe os pacientes com botões para excluir.

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>SUS Digital - Gestão de Pacientes</title>
    <!-- Fonte Premium do Google Fonts -->
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;700&display=swap" rel="stylesheet">
    <style>
        :root {
            --primary: #10b981;
            --primary-hover: #059669;
            --bg-dark: #0f172a;
            --bg-card: #1e293b;
            --text-main: #f8fafc;
            --text-muted: #94a3b8;
            --danger: #ef4444;
            --danger-hover: #dc2626;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: 'Outfit', sans-serif;
        }

        body {
            background-color: var(--bg-dark);
            color: var(--text-main);
            min-height: 100vh;
            padding: 2rem;
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        header {
            margin-bottom: 2rem;
            text-align: center;
        }

        h1 {
            font-size: 2.5rem;
            color: var(--primary);
            font-weight: 700;
            letter-spacing: -0.5px;
            margin-bottom: 0.5rem;
        }

        .subtitle {
            color: var(--text-muted);
        }

        .container {
            width: 100%;
            max-width: 1000px;
            display: grid;
            grid-template-columns: 1fr;
            gap: 2rem;
        }

        @media (min-width: 768px) {
            .container {
                grid-template-columns: 350px 1fr;
            }
        }

        .card {
            background: var(--bg-card);
            border-radius: 16px;
            padding: 1.5rem;
            box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.3);
            border: 1px solid rgba(255, 255, 255, 0.05);
        }

        .card h2 {
            font-size: 1.25rem;
            margin-bottom: 1.25rem;
            border-bottom: 2px solid rgba(255, 255, 255, 0.05);
            padding-bottom: 0.5rem;
        }

        .form-group {
            margin-bottom: 1rem;
        }

        label {
            display: block;
            margin-bottom: 0.35rem;
            font-size: 0.875rem;
            color: var(--text-muted);
        }

        input, select {
            width: 100%;
            padding: 0.75rem;
            border-radius: 8px;
            border: 1px solid rgba(255, 255, 255, 0.1);
            background: rgba(15, 23, 42, 0.6);
            color: #fff;
            outline: none;
            transition: all 0.2s ease;
        }

        input:focus, select:focus {
            border-color: var(--primary);
            box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
        }

        button {
            width: 100%;
            padding: 0.75rem;
            border: none;
            border-radius: 8px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.2s ease;
        }

        .btn-primary {
            background-color: var(--primary);
            color: #fff;
        }

        .btn-primary:hover {
            background-color: var(--primary-hover);
        }

        .btn-danger {
            background-color: var(--danger);
            color: #fff;
            padding: 0.4rem 0.8rem;
            font-size: 0.85rem;
            border-radius: 6px;
            width: auto;
        }

        .btn-danger:hover {
            background-color: var(--danger-hover);
        }

        /* Tabela customizada */
        .table-container {
            overflow-x: auto;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            text-align: left;
        }

        th, td {
            padding: 0.85rem 1rem;
            border-bottom: 1px solid rgba(255, 255, 255, 0.05);
        }

        th {
            color: var(--text-muted);
            font-weight: 600;
            font-size: 0.875rem;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }

        tr:hover {
            background-color: rgba(255, 255, 255, 0.02);
        }

        .badge-blood {
            background: rgba(16, 185, 129, 0.1);
            color: var(--primary);
            padding: 0.25rem 0.5rem;
            border-radius: 4px;
            font-weight: 700;
            font-size: 0.8rem;
        }
    </style>
</head>
<body>

    <header>
        <h1>SUS Digital</h1>
        <p class="subtitle">Gestão Integrada de Fichas Médicas de Pacientes</p>
    </header>

    <div class="container">
        <!-- FORMULÁRIO DE CADASTRO -->
        <div class="card">
            <h2>Cadastrar Paciente</h2>
            <form id="pacienteForm">
                <div class="form-group">
                    <label for="cod_sus">Número SUS</label>
                    <input type="text" id="cod_sus" required placeholder="Ex: 1234567890">
                </div>
                <div class="form-group">
                    <label for="nome">Nome do Paciente</label>
                    <input type="text" id="nome" required placeholder="Ex: João da Silva">
                </div>
                <div class="form-group">
                    <label for="data_nascimento">Data de Nascimento</label>
                    <input type="date" id="data_nascimento" required>
                </div>
                <div class="form-group">
                    <label for="sexo">Sexo</label>
                    <select id="sexo">
                        <option value="Masculino">Masculino</option>
                        <option value="Feminino">Feminino</option>
                        <option value="Não informado">Não informado</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="tipo_sanguineo">Tipo Sanguíneo</label>
                    <select id="tipo_sanguineo">
                        <option value="A+">A+</option>
                        <option value="A-">A-</option>
                        <option value="B+">B+</option>
                        <option value="B-">B-</option>
                        <option value="AB+">AB+</option>
                        <option value="AB-">AB-</option>
                        <option value="O+">O+</option>
                        <option value="O-">O-</option>
                    </select>
                </div>
                <button type="submit" class="btn-primary">Adicionar Paciente</button>
            </form>
        </div>

        <!-- LISTAGEM DE PACIENTES -->
        <div class="card">
            <h2>Pacientes Cadastrados</h2>
            <div class="table-container">
                <table>
                    <thead>
                        <tr>
                            <th>SUS</th>
                            <th>Nome</th>
                            <th>Nascimento</th>
                            <th>Sexo</th>
                            <th>Sangue</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody id="tabelaCorpo">
                        <!-- Linhas adicionadas dinamicamente via JS -->
                    </tbody>
                </table>
            </div>
        </div>
    </div>

    <script>
        const API_URL = 'http://localhost:8080/api/pacientes';

        // Carrega pacientes ao iniciar a página
        document.addEventListener('DOMContentLoaded', carregarPacientes);

        // Submissão do Formulário de Cadastro
        document.getElementById('pacienteForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            
            // Pega os campos da página
            const cod_sus = document.getElementById('cod_sus').value;
            const nome = document.getElementById('nome').value;
            
            // Converte a data do HTML (AAAA-MM-DD) para DD-MM-AAAA caso a sua função
            // time.Parse("02-01-2006") no backend exija esse formato.
            const inputDate = document.getElementById('data_nascimento').value;
            const [ano, mes, dia] = inputDate.split('-');
            const data_nascimento = `${dia}-${mes}-${ano}`;

            const sexo = document.getElementById('sexo').value;
            const tipo_sanguineo = document.getElementById('tipo_sanguineo').value;

            try {
                const response = await fetch(API_URL, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                        cod_sus,
                        nome,
                        data_nascimento,
                        sexo,
                        tipo_sanguineo
                    })
                });

                if (response.ok) {
                    alert('Paciente inserido com sucesso!');
                    document.getElementById('pacienteForm').reset();
                    carregarPacientes();
                } else {
                    const err = await response.json();
                    alert('Erro ao cadastrar: ' + err.erro);
                }
            } catch (error) {
                console.error('Erro na requisição:', error);
                alert('Erro ao se conectar ao servidor backend.');
            }
        });

        // Função para buscar pacientes no backend
        async function carregarPacientes() {
            const corpo = document.getElementById('tabelaCorpo');
            corpo.innerHTML = '<tr><td colspan="6" style="text-align: center;">Carregando registros...</td></tr>';
            
            try {
                const response = await fetch(API_URL);
                const pacientes = await response.json();
                
                corpo.innerHTML = '';
                if (!pacientes || pacientes.length === 0) {
                    corpo.innerHTML = '<tr><td colspan="6" style="text-align: center; color: var(--text-muted);">Nenhum paciente cadastrado.</td></tr>';
                    return;
                }

                pacientes.forEach(p => {
                    // Formata a exibição de data de nascimento se necessário
                    const dataFormatada = p.Data_nascimento ? p.Data_nascimento.substring(0, 10) : '-';

                    const tr = document.createElement('tr');
                    tr.innerHTML = `
                        <td><strong>${p.Cod_SUS}</strong></td>
                        <td>${p.Nome}</td>
                        <td>${dataFormatada}</td>
                        <td>${p.Sexo}</td>
                        <td><span class="badge-blood">${p.Tipo_sanguineo}</span></td>
                        <td>
                            <button class="btn-danger" onclick="excluirPaciente('${p.Cod_SUS}')">Excluir</button>
                        </td>
                    `;
                    corpo.appendChild(tr);
                });
            } catch (error) {
                console.error('Erro ao listar pacientes:', error);
                corpo.innerHTML = '<tr><td colspan="6" style="text-align: center; color: var(--danger);">Não foi possível carregar a lista de pacientes. Certifique-se de que o backend está rodando.</td></tr>';
            }
        }

        // Função para deletar paciente
        async function excluirPaciente(codSus) {
            if (!confirm(`Deseja realmente excluir o paciente com cartão SUS ${codSus}?`)) {
                return;
            }

            try {
                const response = await fetch(`${API_URL}?cod_sus=${codSus}`, {
                    method: 'DELETE'
                });

                if (response.ok) {
                    alert('Paciente excluído com sucesso!');
                    carregarPacientes();
                } else {
                    alert('Erro ao excluir paciente.');
                }
            } catch (error) {
                console.error('Erro ao deletar:', error);
                alert('Erro na comunicação com o servidor.');
            }
        }
    </script>
</body>
</html>
```

---

## 6. Próximos Passos Recomendados

Para que a transição seja perfeita e o sistema se torne robusto para o ambiente de produção:

1.  **Gerenciamento de Configuração**:
    Evite senhas e portas chumbadas no código (`const server = ...`). Use variáveis de ambiente (pacote `os.Getenv` no Go) ou um arquivo `.env` para carregar as credenciais do banco.
2.  **Pool de Conexões**:
    Nas funções atuais do pacote `funcao`, a conexão com o banco é aberta (`sql.Open`) e testada (`db.Ping`) dentro de quase todas as chamadas. Para a web, o ideal é abrir a conexão **uma única vez** na inicialização do servidor (`main`) e passar a instância de `*sql.DB` já ativa para as funções ou armazená-la de forma global compartilhada. Isso economiza recursos computacionais significativos e evita esgotar o número máximo de conexões permitidas pelo PostgreSQL.
3.  **Biblioteca de Roteamento**:
    Para APIs mais complexas, considere usar frameworks e roteadores rápidos em Go como o [Fiber](https://gofiber.io/) ou [Gin](https://gin-gonic.com/) que já simplificam o parse de JSON, validação de campos e aplicação de middlewares de segurança (CORS, Rate Limit).
