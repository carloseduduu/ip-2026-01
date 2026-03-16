programa {
  funcao inicio() {
    real nota

    leia(nota)
    se(nota == 9 e nota <= 10){
      escreva("NOTA = ",nota," CONCEITO = A")
    }
    senao se(nota >= 7.5 e nota <= 8.9){
      escreva("NOTA = ",nota, " CONCEITO = B")
    }
    senao se(nota >= 6 e nota <= 7.4){
      escreva("NOTA = ",nota, " CONCEITO = C")
    }
    senao se(nota >= 0 e nota <= 5.9){
      escreva("NOTA = ",nota, " CONCEITO = D")
    }
    senao{
      escreva("VALOR INVÁLIDO")
    }
  }
}
