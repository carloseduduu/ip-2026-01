programa {
  inclua biblioteca Matematica --> m
  funcao inicio() {
    real h, aresta, raiz = m.raiz(3,2), Ab, volume
    leia(h,aresta)

    Ab = (3*(aresta*aresta)*raiz)/2
    volume = (1/3)*Ab*h
    escreva("O VOLUME DA PIRAMIDE E = ",m.arredondar(volume,2)," METROS CUBICOS")
  }
}
