programa {
  funcao inicio() {
    
  inteiro conta
  real consumo, valor
  caracter tipo

  escreva("Conta, Consumo e Tipo: \n")
  leia(conta, consumo, tipo)

  se(tipo == 'r'){

    valor = 5 + (0.05*consumo)

  }
  senao se (tipo == 'c' e consumo <= 80){

    valor = 500

  }

  senao se (tipo == 'c'){

    valor = 500 + (consumo-80)*0.25

  }
  senao se (tipo == 'i' e consumo <= 100){

    valor = 800

  }

  senao se (tipo == 'i'){

    valor = 800 + (consumo-100)*0.04

  }

  escreva("CONTA = ",conta,"\n",
          "VALOR DA CONTA = ",valor)

  }
}
