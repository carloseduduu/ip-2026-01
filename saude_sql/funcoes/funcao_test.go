package funcao

import (
	"testing"
)

func TestInspectTableSchema(t *testing.T) {
	_, db, err := Connect()
	if err != nil {
		t.Fatalf("Erro ao conectar: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT column_name, data_type, character_maximum_length 
		FROM information_schema.columns 
		WHERE table_name = 'pacientes'
	`)
	if err != nil {
		t.Fatalf("Erro ao consultar schema: %v", err)
	}
	defer rows.Close()

	t.Log("--- SCHEMA DA TABELA PACIENTES ---")
	for rows.Next() {
		var colName, dataType string
		var maxLen *int
		err := rows.Scan(&colName, &dataType, &maxLen)
		if err != nil {
			t.Errorf("Erro no scan do schema: %v", err)
		}
		if maxLen != nil {
			t.Logf("Coluna: %s | Tipo: %s | Tamanho Max: %d", colName, dataType, *maxLen)
		} else {
			t.Logf("Coluna: %s | Tipo: %s | Tamanho Max: sem limite/null", colName, dataType)
		}
	}
}

func TestConnect(t *testing.T) {
	result, db, err := Connect()
	if err != nil {
		t.Fatalf("Erro ao tentar conectar: %v", err)
	}
	if db == nil {
		t.Fatalf("Conexão estabelecida é nula")
	}
	t.Logf("Sucesso na conexão: %s", result)
}
