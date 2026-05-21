package utils

import (
	"log"
)

func DeleteUser(cod_sus, nome string) error {
	query := `DELETE FROM pacientes WHERE cod_sus = $1 AND nome = $2`
	_, err := DB.Exec(query, cod_sus, nome)
	if err != nil {
		log.Printf("Erro ao apagar usuário do banco de dados: %v", err)
		return err
	}
	log.Println("Paciente apagado com sucesso!")
	return nil
}
