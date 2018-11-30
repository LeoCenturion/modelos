package juego

import (
  "fmt"
  "strings"
  "strconv"
)

type Carta struct  {
  Id, Tipo, era, puntos, edificioGratis, cartaRequerida int
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
func (c *Carta) SePuedeJugar(recursosDisponibles  [CANTIDAD_RECURSOS]int, cartasJugadas [TURNOS]Carta, eraActual int, precioRecurso [CANTIDAD_RECURSOS]int, comodinJugado bool) bool {
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

  monedasNecesarias := 0
  tieneRecursosSuficientes := false
  maxPrecioRecurso := 0
  for i, r := range recursosDisponibles {
    dif := r - c.requiere[i]
    if dif < 0 {
      monedasNecesarias += (-dif) * (precioRecurso[i])
      if precioRecurso[i] > maxPrecioRecurso {
        maxPrecioRecurso = precioRecurso[i]
      }
    }
  }

  if monedasNecesarias <= recursosDisponibles[MONEDA] {
    tieneRecursosSuficientes = true
  } else if comodinJugado && monedasNecesarias - maxPrecioRecurso <= recursosDisponibles[MONEDA] {
    tieneRecursosSuficientes = true
  }

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
