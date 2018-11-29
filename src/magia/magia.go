package main

import (
  "juego"
)
func main() {
  c := juego.Contexto{}
  c.CargarCartas("test.csv")
  c.MostrarCartasDisponibles()
  c.ComenzarSimulacion()
  c.MostrarResultados()
}
