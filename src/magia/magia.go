package main

import (
  "juego"
  "fmt"
)

func main(){
  c := juego.Contexto{}

  //c.CargarCartas("cartas.csv")
  //c.MostrarCartasDisponibles()

  heuristicas := []int{juego.HEURISTICA_ID, juego.HEUTISTICA_PUNTOS, juego.HEURISTICA_PROMEDIO_DE_PRODUCE, 
  	juego.HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA, juego.HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA, 
  	juego.ORNITORRINCO_MAXIMO, juego.CONDOR_ALPINO, juego.SPAGHETTI, juego.CHOCOLATE}

  var mejorHeuristica int
  var mejoresPuntos int
  var cartasJugadas [juego.TURNOS]juego.Carta
  var detalleDePuntos string
  nombreHeuristica := ""
  nombreMejorHeuristica := ""

  fmt.Println("####### HEURISTICAS #######")
  for _, heuristica := range heuristicas{
  	c.CargarCartas("cartas.csv")
	//c.Resetear()
	c.ComenzarSimulacion(heuristica)

	nombreHeuristica = c.ObtenerNombreHeuristica(heuristica)
	fmt.Println("Calculado con heurísitca:", nombreHeuristica, " se obtienen ", c.PuntosTotales, " puntos.")
	if mejoresPuntos <= c.PuntosTotales{
		mejoresPuntos = c.PuntosTotales
		mejorHeuristica = heuristica
		cartasJugadas = c.CartasJugadas
		detalleDePuntos = c.DetalleDePuntos
		nombreMejorHeuristica = nombreHeuristica
	}
  }
  //c.MostrarResultados()
  fmt.Println("\nCalculado con heurísitca:", nombreMejorHeuristica, " se obtienen ", mejoresPuntos, " puntos.")
  c.MostrarResultados(cartasJugadas, detalleDePuntos, mejoresPuntos)
}
