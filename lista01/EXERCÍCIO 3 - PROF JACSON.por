programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {

    inteiro n1, n2, n3

    leia(n1,n2,n3)
    //limitação de dígito maior que 2 dígitos
    se (n1>=10 ou n2>=10 ou n3>=10){
      escreva("DÍGITO INVÁLIDO!")
    }
    
    senao{

    // Organizo os valores por casas decimais
      inteiro concat = (n1*100)+(n2*10)+n3

      escreva(concat,",",mat.potencia(concat,2))
    }

  }
}
