package main

import (
  "strconv"
  "juego"
  "time"
  "math/rand"
  "fmt"
)

func main(){
  c := juego.Contexto{}

  //c.CargarCartas("cartas.csv")
  //c.MostrarCartasDisponibles()

  heuristicas := []int{juego.HEURISTICA_ID, juego.HEUTISTICA_PUNTOS, juego.HEURISTICA_PROMEDIO_DE_PRODUCE,
  	juego.HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA, juego.HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA,
  	juego.ORNITORRINCO_MAXIMO, juego.CONDOR_ALPINO, juego.SPAGHETTI, juego.CHOCOLATE, juego.KOALA_CHICLOSO_PAPOTEADO,
    juego.CHOCOLATE_CON_PAPAS_FRITAS}

  var mejoresPuntos, mejorHeuristica int
  var cartasJugadas [juego.TURNOS]juego.Carta
  var detalleDePuntos string
  nombreHeuristica := ""
  nombreMejorHeuristica := ""

  var cartasRestantes [juego.INDICE_MAXIMO_CARTA]juego.Carta
  var cartasDisponibles [juego.INDICE_MAXIMO_CARTA]juego.Carta
  var recursosDisponibles [juego.CANTIDAD_RECURSOS]int
  var precioRecursos [juego.CANTIDAD_RECURSOS]int
  var cartasJugablesEnTurno [juego.TURNOS][juego.INDICE_MAXIMO_CARTA]juego.Carta
  var comodinManufacturaJugado, comodinMateriaPrimaJugado bool

  fmt.Println("####### HEURISTICAS #######")
  for _, heuristica := range heuristicas{
  	c.CargarCartas("cartas.csv")
	//c.Resetear()
	c.ComenzarSimulacion(heuristica,0)

	nombreHeuristica = c.ObtenerNombreHeuristica(heuristica)
	fmt.Println("Calculado con heurísitca:", nombreHeuristica, " se obtienen ", c.PuntosTotales, " puntos.")
	if mejoresPuntos <= c.PuntosTotales{
		mejoresPuntos = c.PuntosTotales
    mejorHeuristica = heuristica
		cartasJugadas, cartasRestantes, cartasDisponibles, recursosDisponibles, precioRecursos, comodinManufacturaJugado, comodinMateriaPrimaJugado, cartasJugablesEnTurno = c.GetEstado()
		detalleDePuntos = c.DetalleDePuntos
		nombreMejorHeuristica = nombreHeuristica
	}
  }
  //c.MostrarResultados()
  fmt.Println("\nCalculado con heurísitca:", nombreMejorHeuristica, " se obtienen ", mejoresPuntos, " puntos.")
  c.MostrarResultados(cartasJugadas, detalleDePuntos, mejoresPuntos)

  for i := 0; i< 18 ; i ++ {
    cartas := ""
    for j := 0; j < juego.INDICE_MAXIMO_CARTA; j++ {
      if cartasJugablesEnTurno[i][j].Id != juego.NULL {
        cartas += strconv.Itoa(cartasJugablesEnTurno[i][j].Id) + ","
      }
    }
    fmt.Println("cartas jugables en turno t:", cartas)
  }
  cambioEnTurno := 0
  cartaCambiada := 0
  var cartasJugadasFinales [juego.TURNOS]juego.Carta
  for i:= 0; i < 1; i++ {
      s2 := rand.NewSource(time.Now().UnixNano())
      r2 := rand.New(s2)
      turnoCambio := r2.Intn(juego.TURNOS-1) + 1
      turnoAnteriorCambio := turnoCambio -1
      if turnoCambio == 0 {
        turnoAnteriorCambio = 0
      }
      cartaCambio := r2.Intn(juego.INDICE_MAXIMO_CARTA) //TODO : ESTO SE PUEDE OPTIMIZAR BASTANTE
      for ; cartasJugablesEnTurno[turnoAnteriorCambio][cartaCambio].Id == juego.NULL ; {
        cartaCambio = r2.Intn(juego.INDICE_MAXIMO_CARTA)
      }
      it := juego.Contexto{}
      it.Init(cartasJugadas, cartasRestantes, cartasDisponibles, recursosDisponibles, precioRecursos, comodinManufacturaJugado, comodinMateriaPrimaJugado, cartasJugablesEnTurno)
      it.RejugarUltimaCarta(turnoAnteriorCambio, cartaCambio)
      it.ComenzarSimulacion(mejorHeuristica, turnoCambio)
      if mejoresPuntos <= it.PuntosTotales {
    		mejoresPuntos = it.PuntosTotales
    		//cartasJugadas, cartasRestantes, cartasDisponibles, recursosDisponibles, precioRecursos, comodinManufacturaJugado, comodinMateriaPrimaJugado, cartasJugablesEnTurno = it.GetEstado()
        cartasJugadasFinales = it.CartasJugadas
        detalleDePuntos = it.DetalleDePuntos
    		nombreMejorHeuristica = nombreHeuristica
        cartaCambiada = cartaCambio
        cambioEnTurno = turnoCambio
    }
  }
  fmt.Println("\nCalculado con heurísitca:", nombreMejorHeuristica, " se obtienen ", mejoresPuntos, " puntos.")
  c.MostrarResultados(cartasJugadasFinales, detalleDePuntos, mejoresPuntos)
  fmt.Println("\n Se cambio de turno en ", cambioEnTurno, " y se jugo la carta", cartaCambiada)
}
