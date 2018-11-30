package main

import (
  "juego"
)
func main(){
  c := juego.Contexto{}
  c.CargarCartas("cartas.csv")
  c.MostrarCartasDisponibles()
  c.ComenzarSimulacion()
  c.MostrarResultados()
}
