# Documentação Técnica: Servidor HTTP e Sistema de Gerenciamento de Contas em GoLang

Esta documentação descreve de forma abrangente e detalhada a arquitetura, a estrutura de arquivos e o funcionamento do site/servidor desenvolvido em **Go (Golang)** integrado a um banco de dados **PostgreSQL**.

O sistema consiste em uma aplicação web funcional com suporte completo a operações de **CRUD** (Create, Read, Update, Delete) de contas de usuários, utilizando autenticação, criptografia de senhas, consultas dinâmicas preparadas para evitar injeção de SQL e renderização dinâmica de páginas com o mecanismo de templates nativo do Go.

---

## 1. Arquitetura Geral do Sistema

A aplicação adota o modelo **Cliente-Servidor clássico com renderização híbrida**:

```
[ Navegador do Usuário (Cliente) ]
       │ ▲
       │ │ 1. Requisição HTTP (GET / POST) com dados de formulários
       ▼ │ 4. Resposta HTTP (Arquivos estáticos ou HTML dinâmico renderizado)
[ Servidor HTTP Go (Backend) ]
       │ ▲
       │ │ 2. Consultas SQL parametrizadas (via database/sql + lib/pq)
       ▼ │ 3. Respostas das consultas (Tabelas de dados)
[ Banco de Dados PostgreSQL ]
```

1. **Frontend (Cliente)**: Páginas estáticas HTML e estilos CSS que hospedam formulários interativos. Para áreas autenticadas (como a exibição do perfil), é utilizado um template HTML processado no servidor.
2. **Backend (Servidor)**: Um serviço HTTP em Go que gerencia o roteamento de URLs, processa os dados enviados via formulários (codificados como `application/x-www-form-urlencoded`), executa validações de regras de negócios e faz a criptografia de senhas.
3. **Persistência (Banco de Dados)**: Um banco PostgreSQL que armazena os registros dos usuários de forma segura.

---

## 2. Estrutura do Projeto

O código está organizado em diretórios com responsabilidades bem definidas:

```
servidorHTTP/
├── .env                        # Variáveis de ambiente com credenciais do banco
├── docker-compose.yml.example  # Exemplo de configuração de container do Postgres
├── go.mod                      # Definição do módulo Go e dependências
├── go.sum                      # Checksums das dependências instaladas
├── app/                        # Diretório principal da lógica backend Go
│   ├── main.go                 # Ponto de entrada do sistema e definição de rotas
│   ├── handlers/               # Controladores que tratam as requisições HTTP
│   │   ├── helloHandler.go     # Handler de teste de rota simples
│   │   ├── formHandler.go      # Processa a criação de novas contas
│   │   ├── loginHandler.go     # Processa autenticação e exibe o perfil
│   │   ├── updateAccountHandler.go # Processa atualizações parciais de dados
│   │   └── deleteAccountHandler.go # Processa a exclusão da conta
│   └── utils/                  # Utilitários de dados, banco de dados e segurança
│       ├── DB.go               # Função de inserção direta no banco
│       ├── connectToDB.go      # Inicializador da conexão com o PostgreSQL
│       ├── deleteUser.go       # Função SQL para deletar usuários
│       ├── encrypt.go          # Algoritmo de criptografia de senhas
│       ├── getUserByEmail.go   # Estrutura do Usuário e busca por e-mail
│       ├── updateUser.go       # SQL dinâmico e parametrizado para atualizações
│       └── validateUser.go     # Validação de credenciais (e-mail + hash)
└── static/                     # Arquivos estáticos do frontend
    ├── index.html              # Página inicial da aplicação (Sumário)
    ├── profile.html            # Template Go HTML do perfil do usuário
    ├── forms/                  # Formulários HTML específicos
    │   ├── createAccount.html  # Formulário de cadastro
    │   ├── login.html          # Formulário de login
    │   ├── updateAccount.html  # Formulário de atualização cadastral
    │   └── deleteAccount.html  # Formulário de exclusão de conta
    └── styles/                 # Estilizações CSS individuais para cada página
        ├── index.style.css
        ├── profile.style.css
        ├── createAccount.style.css
        ├── login.style.css
        ├── updateAccount.style.css
        └── deleteAccount.style.css
```

---

## 3. Banco de Dados

### 3.1. Variáveis de Ambiente (`.env`)
O servidor busca as credenciais de acesso ao PostgreSQL a partir de um arquivo `.env` localizado na raiz do projeto:

*   `DB_USER`: Usuário do PostgreSQL (Ex: `"postgres"`).
*   `DB_PASSWORD`: Senha de acesso ao banco (Ex: `1234`).
*   `DB_NAME`: Nome do banco de dados (Ex: `"saude"`).
*   `DB_HOST`: Endereço onde o banco está rodando (Ex: `"localhost"`).
*   `DB_PORT`: Porta do serviço de banco de dados (Ex: `"5432"`).

### 3.2. Estrutura da Tabela (`users`)
O sistema gerencia as contas de usuários a partir de uma tabela no banco de dados chamada `users` estruturada da seguinte forma:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    born_date DATE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

*   **`username`**: Guarda o nome completo do usuário.
*   **`email`**: Identificador único de autenticação (não podem existir dois cadastros com o mesmo e-mail).
*   **`born_date`**: Data de nascimento.
*   **`password`**: Armazena a senha do usuário em formato de hash (MD5 criptografado de 32 caracteres).

---

## 4. Inicialização do Servidor e Roteamento (`app/main.go`)

O arquivo [main.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/main.go) é responsável por dar a partida na aplicação:

1.  **Conexão ao Banco de Dados**: Dispara `utils.ConnectToDB()`, que faz o carregamento do arquivo `.env` via biblioteca `godotenv`, realiza o `sql.Open` com o driver `github.com/lib/pq` e efetua um `DB.Ping()` para testar a conectividade ativa. Caso haja falhas na conexão, o servidor interrompe sua execução exibindo um erro fatal.
2.  **Serviço de Arquivos Estáticos**:
    ```go
    fileserver := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileserver)
    ```
    Isso diz ao servidor Go que qualquer requisição que não dê match com as rotas dinâmicas registradas deve ser pesquisada dentro do diretório `./static`. Assim, arquivos como `/forms/login.html` ou `/styles/login.style.css` são servidos automaticamente para o navegador do cliente.
3.  **Registro de Endpoints**:
    Associa caminhos de URL para seus respectivos controladores (Handlers):
    *   `/hello` ➔ `handlers.HelloHandler`
    *   `/form` ➔ `handlers.FormHandler`
    *   `/login` ➔ `handlers.LoginHandler`
    *   `/updateAccount` ➔ `handlers.UpdateAccountHandler`
    *   `/deleteAccount` ➔ `handlers.DeleteAccountHandler`
4.  **Auto-detecção de IP de Rede**:
    O servidor interroga as interfaces de rede da máquina de desenvolvimento utilizando o pacote `net` para encontrar o endereço de IP local (IPv4) ativo. Caso o encontre, imprime no console um link amigável de acesso rápido apontando para a rede local (Ex: `http://192.168.0.15:3000/`). Se não houver rede ativa, ele adota por padrão o localhost (`127.0.0.1`).
5.  **Abertura de Escuta (Listening)**:
    O servidor começa a escutar conexões no socket `"0.0.0.0:3000"` permitindo tráfego vindo tanto do próprio computador local quanto de outros aparelhos na mesma rede local.

---

## 5. Fluxo Detalhado das Rotas e Handlers

### 5.1. Rota Raiz `/` (Arquivos Estáticos)
*   **Método HTTP**: `GET`
*   **Funcionamento**: Direciona para o arquivo `static/index.html`. É o painel inicial do site que serve como hub de navegação para os quatro formulários principais.

---

### 5.2. Rota `/hello` (`HelloHandler`)
*   **Método HTTP**: Apenas `GET`
*   **Arquivo**: [helloHandler.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/handlers/helloHandler.go)
*   **Funcionamento**: Rota simples de verificação. Valida se a URL é estritamente `/hello` e se o método é `GET`. Se sim, escreve a string de texto `"Hello You"` no corpo de resposta.

---

### 5.3. Cadastro de Contas `/form` (`FormHandler`)
*   **Método HTTP**: Apenas `POST`
*   **Arquivo**: [formHandler.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/handlers/formHandler.go)
*   **Passo a Passo**:
    1.  Recebe parâmetros enviados pelo formulário em `static/forms/createAccount.html`: `username`, `email`, `bornDate`, e `password`.
    2.  Invoca `utils.Encrypt(password)` para obter o hash MD5 da senha.
    3.  Chama a rotina `utils.InsertUser(...)` que executa uma query de inserção:
        ```sql
        INSERT INTO users (username, email, born_date, password) VALUES ($1, $2, $3, $4)
        ```
    4.  Se a inserção for bem-sucedida (o e-mail inserido não for duplicado), redireciona o navegador do cliente de volta à página inicial `/index.html` utilizando o status HTTP `303 See Other`. Caso falhe, responde com erro `500 Internal Server Error`.

---

### 5.4. Autenticação e Perfil `/login` (`LoginHandler`)
*   **Método HTTP**: Apenas `POST`
*   **Arquivo**: [loginHandler.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/handlers/loginHandler.go)
*   **Passo a Passo**:
    1.  O usuário envia seu `email` e `password` através do formulário em `static/forms/login.html`.
    2.  O handler encripta a senha enviada para comparação de hashes.
    3.  A função `utils.ValidateUser(email, encryptedPassword)` é acionada. Ela executa um count na base de dados buscando a correspondência:
        ```sql
        SELECT COUNT(*) FROM users WHERE email = $1 AND password = $2
        ```
    4.  Se o resultado for zero (credenciais inválidas), retorna o status `401 Unauthorized` com a mensagem `"Credenciais inválidas"`.
    5.  Se as credenciais forem válidas, a aplicação aciona `utils.GetUserByEmail(email)` para buscar os dados de cadastro (`username`, `email`, `born_date`) e mapeá-los na struct `User`.
    6.  O handler lê o arquivo de template dinâmico [profile.html](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/static/profile.html) localizado na pasta estática e o processa utilizando o pacote `text/template`:
        *   Substitui `{{.Username}}`, `{{.Email}}` e `{{.BornDate}}` pelos valores reais do registro do usuário no banco.
    7.  Retorna o HTML renderizado personalizado para o navegador do usuário.

---

### 5.5. Atualização de Dados `/updateAccount` (`UpdateAccountHandler`)
*   **Método HTTP**: Apenas `POST`
*   **Arquivo**: [updateAccountHandler.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/handlers/updateAccountHandler.go)
*   **Passo a Passo**:
    1.  O formulário em `static/forms/updateAccount.html` envia os dados que o usuário quer alterar (`username`, `email`, `bornDate`, `newPassword`), bem como as credenciais de segurança do usuário atual (`currentEmail` e `currentPassword`).
    2.  Verifica se a senha atual está correta efetuando a validação contra o banco com `utils.ValidateUser(currentEmail, utils.Encrypt(currentPassword))`. Se as credenciais atuais forem inválidas, bloqueia a operação retornando `401 Unauthorized`.
    3.  **Criação Dinâmica da Query**:
        Para que o usuário não seja obrigado a digitar todos os campos (e para evitar que campos em branco limpem o banco de dados), o handler cria um mapa temporário na memória `updates := make(map[string]string)` e adiciona apenas as chaves cujos valores enviados no formulário não estão vazios:
        *   Se `username` não for vazio ➔ insere no mapa de alteração.
        *   Se `email` não for vazio ➔ insere no mapa de alteração.
        *   Se `bornDate` não for vazio ➔ insere no mapa de alteração.
        *   Se `newPassword` não for vazio ➔ aplica hash MD5 e insere no mapa de alteração.
    4.  Chama a função `utils.UpdateUser(currentEmail, updates)` que reconstrói a query de forma dinâmica (veja seção 6.6).
    5.  Após a atualização, o handler busca as informações atualizadas no banco. Ele utiliza o novo e-mail informado (caso tenha mudado) ou o e-mail atual para buscar o registro da struct atualizada com `utils.GetUserByEmail(email)`.
    6.  Parseia e executa o template `static/profile.html` de forma a renderizar o perfil atualizado do usuário na tela instantaneamente.

---

### 5.6. Exclusão de Contas `/deleteAccount` (`DeleteAccountHandler`)
*   **Método HTTP**: Apenas `POST`
*   **Arquivo**: [deleteAccountHandler.go](file:///c:/Users/Carlos/Desktop/Github%20-%20Meus%20Reposit%C3%B3rios/servidorHTTP/app/handlers/deleteAccountHandler.go)
*   **Passo a Passo**:
    1.  Recebe `email` e `password` do formulário em `static/forms/deleteAccount.html`.
    2.  Executa validação com `utils.ValidateUser` para assegurar que apenas o proprietário da conta (ou alguém de posse da senha) possa excluí-la. Se for inválido, bloqueia com `401 Unauthorized`.
    3.  Chama `utils.DeleteUser(email)` que remove permanentemente a linha da tabela através do comando:
        ```sql
        DELETE FROM users WHERE email = $1
        ```
    4.  Após a exclusão com sucesso do banco de dados, redireciona o cliente para a página inicial com status `303 See Other`.

---

## 6. Lógica das Funções Utilitárias (`app/utils/`)

### 6.1. `connectToDB.go`
Responsável por orquestrar a inicialização de conexão. Carrega variáveis via `godotenv.Load()`, lê os dados do ambiente, formata a string de conexão desativando SSL (`sslmode=disable`), inicializa a variável global `DB *sql.DB` e executa um `Ping()` de integridade.

### 6.2. `encrypt.go`
Encapsula o algoritmo de dispersão MD5. A função `Encrypt` recebe a string de texto puro, escreve seus bytes em um criador de hash (`md5.New()`) e converte os bytes resultantes em uma string legível codificada em hexadecimal com `hex.EncodeToString(...)`.

### 6.3. `DB.go`
Fornece suporte para o cadastro inicial. Executa o comando de `INSERT` passando os parâmetros dinamicamente. Caso ocorra erro (como violação de restrição única do e-mail), repassa o erro para o handler correspondente tratar.

### 6.4. `validateUser.go`
Realiza uma query do tipo `SELECT COUNT(*)` comparando o e-mail e a senha criptografada passados. Retorna um valor booleano (`true`/`false`) informando se o usuário é autêntico.

### 6.5. `getUserByEmail.go`
Puxa os dados de exibição do usuário. Define a struct `User` que contém apenas `Username`, `Email` e `BornDate`. Faz um scan na query para preencher os dados de memória e retorná-los para renderização no template de perfil.

### 6.6. `updateUser.go` (Construção Dinâmica de Queries)
Esta é uma das partes mais robustas do backend. Como o formulário de atualização permite que o usuário altere apenas os campos que desejar, a query SQL não pode ser estática. 

A função `UpdateUser` monta o comando SQL dinamicamente usando um mapa de atualizações, prevenindo injeções de SQL por meio de placeholders dinâmicos:

```go
func UpdateUser(currentEmail string, updates map[string]string) error {
    setClauses := []string{}
    values := []interface{}{}
    i := 1

    for column, value := range updates {
        // Ex: "username = $1", "email = $2", etc.
        setClauses = append(setClauses, fmt.Sprintf("%s = $%d", column, i))
        values = append(values, value)
        i++
    }

    values = append(values, currentEmail)
    // Constrói ex: "UPDATE users SET username = $1, email = $2 WHERE email = $3"
    query := fmt.Sprintf("UPDATE users SET %s WHERE email = $%d", strings.Join(setClauses, ", "), i)

    _, err := DB.Exec(query, values...)
    return err
}
```

---

## 7. Mecanismos de Segurança Implementados

1.  **Proteção contra Injeção de SQL (SQL Injection)**:
    O sistema **não** faz concatenação direta de strings nas consultas executadas (`"WHERE email = '" + email + "'"`). Em vez disso, todas as consultas utilizam placeholders de argumentos do PostgreSQL (`$1`, `$2`, `$3`, etc.) processados por prepared statements nativos do `database/sql`. Os dados fornecidos pelo usuário são enviados separadamente e tratados como valores literais no banco, anulando a possibilidade de injeção de scripts maliciosos.
2.  **Proteção de Senhas em Trânsito/Banco**:
    As senhas são convertidas para hash MD5 no servidor no instante em que as requisições são recebidas, garantindo que o banco de dados nunca armazene as senhas originais dos usuários em texto claro.
3.  **Autorização para Alterações Destrutivas/Modificadoras**:
    Para atualizar dados sensíveis (como e-mail e senha) ou deletar a conta, as funções de handler sempre validam as credenciais atuais de e-mail e senha antes de delegar a instrução às funções do banco de dados, protegendo contra acessos forjados.

---

## 8. Frontend e Interface com o Usuário

A interface frontend utiliza HTML5 semântico e CSS customizado separado para manter uma boa legibilidade e organização:

*   **Páginas Individuais de Formulário**: Cada formulário (`createAccount.html`, `login.html`, `updateAccount.html`, `deleteAccount.html`) é uma página independente ligada a um arquivo de folha de estilos CSS próprio na pasta `styles`. Isso facilita modificações visuais e evita poluição de estilos globais.
*   **Go Templates**: O perfil (`profile.html`) utiliza a renderização nativa de templates do Go. Ele funciona como uma página dinâmica: quando renderizado, insere os dados da struct de forma limpa na estrutura do HTML e retorna a página totalmente preenchida para o cliente.
