programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
  inteiro hr, min, seg
  leia(hr, min, seg)

  seg = (hr*3600)+(min*60)+seg
  escreva("O TEMPO EM SEGUNDOS E = ",seg,"\n")

  }  
}
