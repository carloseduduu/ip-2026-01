package utils

import (
	"log"

	_ "github.com/lib/pq"
)

func Insert(cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo string) error {

	err := DB.Ping() //Verifica se há conexão estável com o servidor
	if err != nil {
		DB.Close()
		panic(err)
	}

	query := `INSERT INTO pacientes (cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo) VALUES ($1, $2, $3, $4, $5)`

	_, err = DB.Exec(query, cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo)

	if err != nil {
		log.Printf("Erro ao inserir usuário no banco de dados: %v", err)
		return err
	}

	log.Println("Usuário inserido com sucesso!")

	return nil
}
