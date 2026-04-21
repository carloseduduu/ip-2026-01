programa {
  funcao inicio() {
  inteiro n1, n2, n3
  inteiro total = 0

  leia(n1, n2, n3)
    para(inteiro contador = 0;contador<n3;contador++){
      total = total+n1
      n1 = n1+n2
    }
    escreva(total,"\n")
  }  
}
