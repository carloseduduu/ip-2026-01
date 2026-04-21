programa {
  inclua biblioteca Matematica --> mat
  funcao inicio(){
    
    real n1, n2, n3, media
    escreva("Insira cada nota individualmente e pressione ENTER: \n")
    leia (n1,n2,n3)

    escreva ("Suas notas foram: \n",
    n1,"\n", n2,"\n", n3,"\n")

    media = (n1+n2+n3)/3

   escreva ("Sua média final foi de: \n", "MÉDIA = ",mat.arredondar (media, 1),"\n")
    se (media >= 6)
    {
      escreva ("APROVADO")
    }
    senao
    {
      escreva ("REPROVADO")
    }





  
  }
}
