programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
  real soma = 0
  inteiro k

  leia(k)

  se(k<1){
    escreva("Numero invalido!\n")
  } 
    senao{
      para (inteiro i=1;i<k;i++){
      soma = soma + (1/i)}
      escreva(m.arredondar(soma,6))
    }
    
  }  
}
