programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real fah, celsius, polegada, mm

    leia(fah,polegada)
    
    celsius = ((5*fah-160)/9)
    mm = polegada*25.4

    escreva("O VALOR EM CELSIUS = ", mat.arredondar(celsius, 2),"\n")
    escreva("A QUANTIDADE DE CHUVA É = ",mat.arredondar(mm,2))
  }
}
