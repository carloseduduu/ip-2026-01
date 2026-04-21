programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
  
  real raio, altura, m2, custo
  const real pi = 3.14159
  leia(raio, altura)

  m2 = (2*(pi*(raio*raio)))+(2*pi*raio*altura)
  custo = m2*100.00
  escreva("O VALOR DO CUSTO É = ",mat.arredondar(custo,2),"\n")

  }
}
