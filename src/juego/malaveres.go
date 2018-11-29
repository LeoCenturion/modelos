package juego

import (
  "strconv"
  "fmt"
  "encoding/csv"
  "io/ioutil"
  "strings"
  "log"
)


type Contexto struct {
  cartasJugadas [TURNOS]Carta
  cartasDisponibles, cartasRestantes [INDICE_MAXIMO_CARTA]Carta
  monedas [TURNOS]int
  recursosDisponibles [CANTIDAD_RECURSOS]int
  escudos [TURNOS]int
  puntosTotales int
  comodinJugado bool
}

//Carga de un csv todas las cartas disponibles
func (c *Contexto) CargarCartas(archivo string) {
  for i, _ := range c.cartasJugadas {
    c.cartasJugadas[i].Id = NULL
  }

  for i, _ := range c.cartasDisponibles {
    c.cartasDisponibles[i].Id = NULL
  }

  data, _ := ioutil.ReadFile(archivo)
  r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()

  if err != nil {
		log.Fatal(err)
	}

  fmt.Println("### Cabecera ### ")
  fmt.Println(records[0])
  records = records[1:]
  recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}

  for _, row := range records {
    nuevaCarta := Carta{}
    var produce [CANTIDAD_RECURSOS]int
    var requiere [CANTIDAD_RECURSOS]int
    for _, recurso := range recursos {
      produce[recurso], _ = strconv.Atoi(row[recurso])
      requiere[recurso], _ = strconv.Atoi(row[OFFSET_REQUIERE + recurso])
    }
    id, _ := strconv.Atoi(row[ID])
    tipo, _ := strconv.Atoi(row[TIPO])
    era, _ := strconv.Atoi(row[ERA])
    puntos, _ := strconv.Atoi(row[PUNTOS])
    edificioGratis, _ := strconv.Atoi(row[EDIFICIO_GRATIS])
    cartaRequerida, _ := strconv.Atoi(row[CARTA_REQUERIDA])
    nombre := row[NOMBRE]
    nuevaCarta.Init(id, tipo, era, puntos,edificioGratis, cartaRequerida, nombre, produce, requiere)
    c.cartasDisponibles[id] = nuevaCarta
  }
  cartaNoHacerNada := Carta{}
  produce :=  [CANTIDAD_RECURSOS]int{}
  requiere :=  [CANTIDAD_RECURSOS]int{}
  produce[MONEDA] = 3
  cartaNoHacerNada.Init(NO_HACER_NADA,NO_HACER_NADA, CUALQUIERA, 0, 0,0, "No hacer nada", produce, requiere)
  c.cartasDisponibles[NO_HACER_NADA] = cartaNoHacerNada
  c.cartasRestantes = c.cartasDisponibles //TODO: esto tal vez copia la referencia del array, por ahora no nos jode igual por que no se si vamos a usar cartas disponibles y restantes
}

//Determina la era a la que pertence el turno t
func (c Contexto) eraEnTurno(t int) (era int){
  era = 1
  if t >=7 && t <= 12 {
    era = 2
  } else if t >= 13 {
    era = 3
  }
  return
}

//Obtiene la carta que se debería jugar en el turno t, dado el estado de cartasJugadas.
func (c *Contexto) SimularTurno(t int, cartasJugadas [TURNOS]Carta) (cartaJugada Carta){
  pesoMaximo := float32(0) //TODO: Esto se puede hacer acumulativo entre turnos y tener memoria
  eraActual := c.eraEnTurno(t)
  for _, carta := range c.cartasRestantes {
    if carta.Id != NULL && carta.SePuedeJugar(c.recursosDisponibles, cartasJugadas, eraActual,c.comodinJugado) {
      pesoCarta := c.calcularPeso(cartasJugadas, c.cartasRestantes, carta)
      if pesoCarta > pesoMaximo {
        pesoMaximo = pesoCarta
        cartaJugada = carta
      }
    }
  }
  return
}

//Calcula el ponderable para un turno
func (c *Contexto) calcularPeso(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
  peso = float32(unaCarta.Id)
  return
}

//Calcula los puntos segun todas las cartasJugadas
func (c *Contexto) calcularPuntos() {
  c.puntosTotales = 0
  for _, carta := range c.cartasJugadas {
    c.puntosTotales += carta.Id
  }
}

//Realiza la heurística de construcción
func (c *Contexto) ComenzarSimulacion() {
  recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}
  for t:= 0; t < TURNOS;t++ {
    cartaJugada := c.SimularTurno(t, c.cartasJugadas)
    c.cartasJugadas[t] = cartaJugada
    if cartaJugada.Id != NO_HACER_NADA {
      c.cartasRestantes[cartaJugada.Id].Id = NULL
    }
    if cartaJugada.Id == ID_CARTA_COMODIN {
      c.comodinJugado = true
    }

    for r, _ := range recursos {
      c.recursosDisponibles[r] += cartaJugada.Produce[r]
    }
  }
  c.calcularPuntos()
}

//muestra los resultados de la simulación (las cartas que se deciden jugar)
func (c Contexto) MostrarResultados() {
  fmt.Println("### Resultados ###")
  for i, e := range c.cartasJugadas {
    fmt.Println("Carta jugada en turno", i+1, ":",e.Nombre)
  }
  fmt.Println("Puntos obtenidos:", c.puntosTotales)
}

//Muestra todas las cartas con toda la información cargada del csv
func (c Contexto) MostrarCartasDisponibles() {
  fmt.Println("### Cartas cargadas del csv ###")
  for _, carta := range c.cartasDisponibles {
    if carta.Id != NULL {
      fmt.Println(carta.MostrarCarta())
    }
  }
}
