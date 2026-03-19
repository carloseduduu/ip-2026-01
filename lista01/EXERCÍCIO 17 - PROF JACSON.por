programa {
  funcao inicio() {
    inteiro n1, n2, limite
    leia(n1, limite)
    se (n1%2 == 0){

    para(inteiro i = 0;i<limite;i++){
      escreva(n1,"\n")
      n1 = n1+2
    }
    }senao{
      escreva("O PRIMEIRO NUMERO NAO E PAR")
    }
  }
}
