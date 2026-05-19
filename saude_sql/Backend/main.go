package main

import (
	"bufio"
	f "fmt"
	"os"
	funcao "saude_sql/funcoes"
	"strings"
	"time"
)

func main() {

	result, _, err := funcao.Connect()

	if err != nil {
		f.Print("ERRO NA CONEXÃO!")
	} else {
		f.Println(result + "✅")
	}
	var opcao string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for {

			f.Println("\nSEJA BEM VINDO AO SUS - DIGITAL!")
			f.Println("1 - CADASTRO PACIENTE")
			f.Println("2 - LISTA DE PACIENTES")
			f.Println("3 - EDITAR DADOS")
			f.Println("4 - EXCLUIR USUÁRIO")
			f.Println("5 - SAIR DO SISTEMA")
			scanner.Scan() //Função scan valores com espaço
			opcao = strings.TrimSpace(scanner.Text())

			switch opcao {
			case "1":
				var cod_SUS, nome, data_nascimento, sexo, tipo_sanguineo string
				f.Println("Número SUS")
				scanner.Scan()
				cod_SUS = strings.TrimSpace(scanner.Text())
				f.Println("Nome do paciente")
				scanner.Scan()
				nome = strings.TrimSpace(scanner.Text())
				f.Println("Data de Nascimento(DD-MM-AAAA)")
				scanner.Scan()
				data_nascimento = strings.TrimSpace(scanner.Text())
				data_nascimento_formatada, _ := time.Parse("02-01-2006", data_nascimento)
				f.Println("Sexo: ")
				f.Println("1 - Masculino")
				f.Println("2 - Feminino")
				f.Println("3 - Não informar")
				scanner.Scan()
				opcao = strings.TrimSpace(scanner.Text())

				switch opcao {
				case "1":
					sexo = "Masculino"
				case "2":
					sexo = "Feminino"
				case "3":
					sexo = "Não informado"
				default:
					sexo = "null"
				}

				f.Println("Selecione seu tipo sanguíneo:")
				f.Println("1 - A+")
				f.Println("2 - A-")
				f.Println("3 - B+")
				f.Println("4 - B-")
				f.Println("5 - AB+")
				f.Println("6 - AB-")
				f.Println("7 - O+")
				f.Println("8 - O-")

				scanner.Scan()
				opcao = strings.TrimSpace(scanner.Text())

				switch opcao {
				case "1":
					tipo_sanguineo = "A+"
				case "2":
					tipo_sanguineo = "A-"
				case "3":
					tipo_sanguineo = "B+"
				case "4":
					tipo_sanguineo = "B-"
				case "5":
					tipo_sanguineo = "AB+"
				case "6":
					tipo_sanguineo = "AB-"
				case "7":
					tipo_sanguineo = "O+"
				case "8":
					tipo_sanguineo = "O-"
				default:
					tipo_sanguineo = ""
				}
				funcao.Insert("pacientes", cod_SUS, nome, data_nascimento_formatada.String(), sexo, tipo_sanguineo)
			case "2":
				lista, err := funcao.List_pacientes()

				if err != nil {
					f.Println("Erro de Leitura!")
					//panic(err)
					break
				}

				f.Println("---LISTA PACIENTES---")
				for _, v := range lista {
					f.Printf("SUS: %s\tNome: %s\tData Nasc.: %s\tSexo: %s\tTipo Sanguíneo: %s\t\n", v.Cod_SUS, v.Nome, v.Data_nascimento, v.Sexo, v.Tipo_sanguineo)
				}
			case "3":
				f.Println("Qual o Código SUS do paciente?")
				scanner.Scan()
				cod_sus := scanner.Text()
				f.Println("Qual valor alterar?")
				f.Println("1 - Número SUS")
				f.Println("2 - Nome")
				f.Println("3 - Data Nascimento")
				f.Println("4 - Sexo")
				f.Println("5 - Tipo Sanguíneo")
				scanner.Scan()
				opcao := strings.TrimSpace(scanner.Text())
				var dado string
				switch opcao {
				case "1":
					f.Println("Insira o novo código: ")
					scanner.Scan()
					dado = scanner.Text()
					funcao.Update("pacientes", cod_sus, "cod_sus", dado)
				case "2":
					f.Println("Insira o novo nome: ")
					scanner.Scan()
					dado = scanner.Text()
					funcao.Update("pacientes", cod_sus, "nome", dado)
				case "3":
					f.Println("Insira a nova Data de Nascimento: ")
					scanner.Scan()
					dado = scanner.Text()
					funcao.Update("pacientes", cod_sus, "data_nascimento", dado)
				case "4":
					f.Println("Selecione o novo sexo")
					f.Println("1 - Masculino")
					f.Println("2 - Feminino")
					f.Println("3 - Não informado")
					scanner.Scan()
					opcao := scanner.Text()
					switch opcao {
					case "1":
						funcao.Update("pacientes", cod_sus, "sexo", "Masculino")
					case "2":
						funcao.Update("pacientes", cod_sus, "sexo", "Feminino")
					case "3":
						funcao.Update("pacientes", cod_sus, "sexo", "Não informado")
					default:
						funcao.Update("pacientes", cod_sus, "sexo", "null")
					}
				case "5":
					f.Println("Insira o novo tipo sanguíneo: ")
					f.Println("1 - A+")
					f.Println("2 - A-")
					f.Println("3 - B+")
					f.Println("4 - B-")
					f.Println("5 - AB+")
					f.Println("6 - AB-")
					f.Println("7 - O+")
					f.Println("8 - O-")

					scanner.Scan()
					opcao = strings.TrimSpace(scanner.Text())

					switch opcao {
					case "1":
						dado = "A+"
					case "2":
						dado = "A-"
					case "3":
						dado = "B+"
					case "4":
						dado = "B-"
					case "5":
						dado = "AB+"
					case "6":
						dado = "AB-"
					case "7":
						dado = "O+"
					case "8":
						dado = "O-"
					default:
						dado = ""
					}
					funcao.Update("pacientes", cod_sus, "tipo_sanguineo", dado)
				default:

				}
			case "4":
				f.Println("Qual o código SUS a ser excluído: ")
				scanner.Scan()
				dado := scanner.Text()
				f.Println(funcao.Delete("pacientes", dado))

			case "5":
				f.Println("Saindo do sistema...")
			}
			if opcao == "5" {
				break
			}
		}
		if opcao == "5" {
			break
		}
	}
}
