programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    
    inteiro qntd, contador
    real c, f
    real temp[1000]
    real fahr[1000]

    escreva("Número de temperaturas em Fahrenheit: \n")
    leia(qntd)
    
    para(contador = 0;contador <= qntd-1; contador++){

      escreva("Medida de temperatura em graus Fahrenheit: \n")
      leia(f)
      fahr[contador] = f
      c = (5*(f-32))/9
      temp[contador] = c
      
    }

    
    para (contador = 0; contador <= qntd-1; contador ++){
      escreva(mat.arredondar(fahr[contador],2)," FAHRENHEIT EQUIVALE A ",mat.arredondar(temp[contador],2), " CELSIUS\n")
    }









  }
}
