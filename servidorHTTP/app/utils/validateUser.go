package utils

import (
	"log"
)

// Valida as credenciais do usuário no banco de dados

func ValidateUser(cod_sus string) (bool, error) {
	//query := `SELECT COUNT(*) FROM saude WHERE cod_sus = $1 AND password = $2`
	query := `SELECT COUNT(*) FROM pacientes WHERE cod_sus = $1`
	var count int
	err := DB.QueryRow(query, cod_sus).Scan(&count)
	if err != nil {
		log.Printf("Erro ao validar usuário no banco de dados: %v", err)
		return false, err
	}
	return count > 0, nil
}

func ValidateDelete(cod_sus, nome string) (bool, error) {
	query := `SELECT COUNT (*) FROM pacientes WHERE cod_sus = $1 AND nome = $2`
	var count int
	err := DB.QueryRow(query, cod_sus, nome).Scan(&count)
	if err != nil || count <= 0 {
		DB.Close()
		log.Println("Erro ao validar NOME E NÚMERO SUS")
		return false, err
	}

	return count > 0, nil
}
