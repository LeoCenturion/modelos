package juego

import (
  "fmt"
  "strings"
  "strconv"
)

type Carta struct  {
  Id, Tipo, era, puntos, edificioGratis, cartaRequerida, monedasNecesarias int
  Produce, requiere [CANTIDAD_RECURSOS]int
  Nombre string
}

//Inicializar los parametros de la carta
func (c *Carta) Init(id, tipo, era, puntos, edificioGratis, cartaRequerida int, nombre string, produce, requiere [CANTIDAD_RECURSOS]int){
  c.Id = id
  c.Tipo = tipo
  c.era = era
  c.puntos = puntos
  c.edificioGratis = edificioGratis
  c.Nombre = nombre
  c.Produce = produce
  c.requiere = requiere
  c.cartaRequerida = cartaRequerida
}

//Determina si la carta se puede jugar según los parámetros

func (c *Carta) SePuedeJugar(recursosDisponibles  [CANTIDAD_RECURSOS]int, cartasJugadas [TURNOS]Carta, eraActual int, precioRecurso [CANTIDAD_RECURSOS]int, comodinMateriaPrimaJugado, comodinManufacturaJugado bool) bool {
  c.monedasNecesarias = 0

  if !(c.era == eraActual || c.era == CUALQUIERA) { return false }

  cartaRequeridaJugada := c.cartaRequerida == NINGUNA
  if c.edificioGratis != NINGUNA || !cartaRequeridaJugada {
    for _, e := range cartasJugadas {
      if e.Id == c.edificioGratis { return true }
      if e.Id == c.cartaRequerida {
        cartaRequeridaJugada = true
        break
      }
   }
   if !cartaRequeridaJugada { return false }
  }
  tieneRecursosSuficientes := false
  maxPrecioRecurso := 0
  maxEsMateriaPrima := false
  for i, r := range recursosDisponibles {
    dif := r - c.requiere[i]
    if dif < 0 {
      c.monedasNecesarias += (-dif) * (precioRecurso[i])
      if precioRecurso[i] > maxPrecioRecurso {
        maxPrecioRecurso = precioRecurso[i]
        maxEsMateriaPrima = recursoEsMateriaPrima(r)
      }
    }
  }

  if c.monedasNecesarias <= recursosDisponibles[MONEDA] {
    tieneRecursosSuficientes = true
  } else if maxEsMateriaPrima {
    if comodinMateriaPrimaJugado && c.monedasNecesarias - maxPrecioRecurso <= recursosDisponibles[MONEDA] {
      tieneRecursosSuficientes = true
      c.monedasNecesarias -= maxPrecioRecurso
    }
  } else {
    if comodinManufacturaJugado && c.monedasNecesarias - maxPrecioRecurso <= recursosDisponibles[MONEDA] {
      tieneRecursosSuficientes = true
      c.monedasNecesarias -= maxPrecioRecurso
    }
  }
/*  if comodinJugado { c.monedasNecesarias -= maxPrecioRecurso }*/
  return tieneRecursosSuficientes
}

//Muestra la información de la carta
func (c Carta) MostrarCarta() (informacion string){
  informacion = ""
  informacion += strconv.Itoa(c.Id) + "|" + strconv.Itoa(c.Tipo) + "|" + strconv.Itoa(c.era) + "|" + strconv.Itoa(c.puntos) +
                                      "|" + strconv.Itoa(c.edificioGratis) + "|" + strconv.Itoa(c.cartaRequerida) + "|" +
                                      c.Nombre + "|" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.Produce)), ","), "[]") + "|" +
                                      strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.requiere)), ","), "[]")
  return
}
