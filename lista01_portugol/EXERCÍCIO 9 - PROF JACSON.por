programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
    real a, b, c, delta
    leia(a,b,c)
    delta = (b*b) - 4.00*a*c
    escreva("O VALOR DE DELTA É = ",m.arredondar(delta,2),"\n")
  }
}
