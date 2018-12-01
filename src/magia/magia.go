package main

import (
  "juego"
  "fmt"
)

//identificadores de las heuristicas
const HEURISTICA_ID = 0
const HEUTISTICA_PUNTOS = 1
const HEURISTICA_PROMEDIO_DE_PRODUCE = 2
const HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA = 3
const HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA = 4
const ORNITORRINCO_MAXIMO = 5
const CONDOR_ALPINO = 6
const SPAGHETTI = 7
const CHOCOLATE = 8
const KOALA_CHICLOSO_PAPOTEADO = 9
func main(){
  c := juego.Contexto{}

  //c.CargarCartas("cartas.csv")
  //c.MostrarCartasDisponibles()

  heuristicas := []int{HEURISTICA_ID, HEUTISTICA_PUNTOS, HEURISTICA_PROMEDIO_DE_PRODUCE,
  	HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA, HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA,
  	ORNITORRINCO_MAXIMO, CONDOR_ALPINO, SPAGHETTI, CHOCOLATE, KOALA_CHICLOSO_PAPOTEADO}

  var mejorHeuristica int
  var mejoresPuntos int

	fmt.Println("#######HEURISTICAS#######")
  for _, heuristica := range heuristicas{
  	c.CargarCartas("cartas.csv")
	//c.Resetear()
	c.ComenzarSimulacion(heuristica)
	fmt.Println("Calculado con heurísitca:", heuristica, " se obtienen ", c.PuntosTotales, " puntos.")
	if mejoresPuntos <= c.PuntosTotales{
		mejoresPuntos = c.PuntosTotales
		mejorHeuristica = heuristica
	}
  }
  //c.MostrarResultados()
  fmt.Println("\nCalculado con heurísitca:", mejorHeuristica, " se obtienen ", mejoresPuntos, " puntos.")
  c.CargarCartas("cartas.csv")
  c.ComenzarSimulacion(mejorHeuristica)
  c.MostrarResultados()
}
