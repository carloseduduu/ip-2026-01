programa {
  inclua biblioteca Util --> u
  funcao inicio() {
    real a, b, c, d, det
    real matriz[2][2]
    leia(matriz[0][0],
         matriz[0][1],
         matriz[1][0],
         matriz[1][1])
    a = matriz[0][0]
    b = matriz[0][1]
    c = matriz[1][0]
    d = matriz[1][1]

    det = (a*d)-(b*c)
    escreva("O VALOR DO DETERMINANTE E = ",det)
  }
}
