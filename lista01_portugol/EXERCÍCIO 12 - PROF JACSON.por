programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
  inteiro horas, total, valor, total2

  leia(horas)
  total = horas/3 // Tenho quantos valores inteiros de 3h vai pagar
  total2 = horas % 3 // Tenho quantas horas vão sobrar(resto) para pagar 5 reais
  valor = (total*10)+(total2*5)
  escreva("O VALOR A PAGAR E = ",m.arredondar(valor,2))
  
  }
}
