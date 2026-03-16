programa {
  funcao inicio() {
    real resto3, resto5
    inteiro num_int
    leia(num_int)
    resto3 = num_int % 3
    resto5 = num_int % 5
    se (resto3 == 0 e resto5 == 0){
      escreva("O NUMERO É DIVISIVEL")
    }
    senao{
      escreva("O NUMERO NÃO É DIVISIVEL")
    }


  }
}
