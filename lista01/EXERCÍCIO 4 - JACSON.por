programa {

  inclua biblioteca Matematica --> mat

  funcao inicio() {
    
    real salario, kw
    
    leia(salario,kw)

    // 100 kw = salario*((70)/100)
    // 1 kw = salario*(0.7/100)

    real custo_kw = salario*(0.7/100)
    real consumo = kw*custo_kw
    real desconto = consumo-(consumo*0.10)

    escreva("Custo por kW: R$ ", mat.arredondar(custo_kw,5),"\n",
            "Custo do consumo: R$ ", mat.arredondar(consumo,2),"\n",
            "Custo com desconto: R$ ", mat.arredondar(desconto,2))


  }
}
