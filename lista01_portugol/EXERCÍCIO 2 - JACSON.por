programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    
    inteiro i, jogo
    real pop, geral, arq, cad, ingresso, renda
    real valor_arquibancada = 10
    real valor_cadeiras =20
    real valor_geral =5
    real valor_popular = 1
      
    escreva("Insira a quantidade de jogos e pressione ENTER respectivamente \n",
    "Jogo: ")
    leia (jogo)

    real renda_lista[jogo]

    para(i=0; i < jogo; i++)

    {
      escreva("Quantos ingressos foram vendidos no jogo ",i+1,"?\n","Ingressos: ")
      leia(ingresso)
      escreva("\n")
      escreva("Que porcetagem de pessoas compraram ingresso para a categoria popular?\n","% de Popular = ")
      leia(pop)
      escreva("\n")
      escreva("Que porcetagem de pessoas compraram ingresso para a categoria geral?\n","% de Geral = ")
      leia(geral)
      escreva("\n")
      escreva("Que porcetagem de pessoas compraram ingresso para a categoria arquibancada?\n","% de Arquibancada = ")
      leia(arq)
      escreva("\n")
      escreva("Que porcetagem de pessoas compraram ingresso para a categoria cadeiras?\n","% de Cadeiras = ")
      leia(cad)
      escreva("\n")
      
      renda = ((ingresso*(pop/100))*valor_popular)+((ingresso*(geral/100))*valor_geral)+((ingresso*(arq/100))*valor_arquibancada)+((ingresso*(cad/100))*valor_cadeiras)
      renda_lista[i] = renda
      escreva("\n")
    }

    para (i=0; i < jogo; i++)
    {
      escreva("A RENDA DO JOGO N. ",i+1," É = ", mat.arredondar(renda_lista[i],2),"\n")
    }
  
  }
}
