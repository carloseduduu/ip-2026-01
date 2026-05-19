package funcao

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const server = "host=localhost port=5432 user=postgres password=1234 dbname=saude sslmode=disable"

func Connect() (string, *sql.DB, error) {

	db, err := sql.Open("postgres", server)
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		panic(err)
	}

	result := ("Connected to " + "saude")

	return result, db, err
}

func Insert(tabela, cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo string) string {
	db, err := sql.Open("postgres", server)
	if err != nil {
		db.Close()
		return "ERRO!"
	}

	err = db.Ping() //Verifica se há conexão estável com o servidor
	if err != nil {
		db.Close()
		panic(err)
	}

	command := fmt.Sprintf("INSERT INTO %s (cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo) VALUES ($1, $2, $3, $4, $5)", tabela)

	query, err := db.Prepare(command)
	_, err = query.Exec(cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo)
	if err != nil {
		db.Close()
		panic(err)
	}

	return "DADOS INSERIDOS COM SUCESSO!"
}

func Create_table(arquivo string) {
	db, err := sql.Open("postgres", server)
	if err != nil {
		db.Close()
		panic(err)
	}

	err = db.Ping() //Verifica se há conexão estável com o servidor
	if err != nil {
		db.Close()
		panic(err)
	}

	dados, err := os.ReadFile("./database/pacientes.sql")
	conteudo := string(dados)
	query, err := db.Prepare(conteudo)
	if err != nil {
		panic(err)
	}
	_, err = query.Exec()
	if err != nil {
		panic(err)
	}
}

type Paciente struct {
	Cod_SUS, Nome, Data_nascimento, Sexo, Tipo_sanguineo string
}

func List_pacientes() ([]Paciente, error) {
	db, err := sql.Open("postgres", server)
	if err != nil {
		panic(err)
	}
	list, err := db.Query("SELECT * FROM pacientes ORDER BY cod_sus")
	if err != nil {
		panic(err)
	}
	var lista []Paciente
	defer db.Close()
	defer list.Close()
	for list.Next() {
		var p Paciente
		list.Scan(&p.Cod_SUS, &p.Nome, &p.Data_nascimento, &p.Sexo, &p.Tipo_sanguineo)

		lista = append(lista, p)
	}
	return lista, nil
}

func Update(tabela, sus, coluna, dado string) {
	db, err := sql.Open("postgres", server)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	up := fmt.Sprintf("UPDATE %s SET %s = $1 WHERE cod_sus = $2", tabela, coluna)
	query, err := db.Prepare(up)
	if err != nil {
		panic(err)
	}
	defer query.Close()
	_, err = query.Exec(dado, sus)
	if err != nil {
		panic(err)
	}

}

func Delete(tabela, cod_sus string) string {
	db, err := sql.Open("postgres", server)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	line := fmt.Sprintf("DELETE FROM %s WHERE cod_sus = $1", tabela)
	query, err := db.Prepare(line)
	if err != nil {
		panic(err)
	}
	_, err = query.Exec(cod_sus)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Paciente SUS: %s DELETADO com sucesso!", cod_sus)
}
