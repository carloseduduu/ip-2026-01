programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
    real si, sr
    leia(si)
    se (si <= 300){
      sr = (si*0.5)+si
      escreva("SALARIO COM REAJUSTE = ",m.arredondar(sr,2))
    } senao{

      sr = (si*0.3)+si
      escreva("SALARIO COM REAJUSTE = ",m.arredondar(sr,2))
    }
    
  }
}
