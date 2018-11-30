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
  recursosDisponibles, precioRecursos [CANTIDAD_RECURSOS]int
  escudos [TURNOS]int
  puntosTotales int
  comodinMateriaPrimaJugado, comodinManufacturaJugado bool
}

//Carga de un csv todas las cartas disponibles
func (c *Contexto) CargarCartas(archivo string) {
	for i, _ := range c.cartasJugadas {
		c.cartasJugadas[i].Id = NULL
	}

	for i, _ := range c.cartasDisponibles {
		c.cartasDisponibles[i].Id = NULL
	}

  for i, _ := range c.precioRecursos {
    c.precioRecursos[i] = PRECIO_INICIAL_RECURSO
  }

  c.recursosDisponibles[MONEDA] = 3
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
  if t >= 7 && t <= 12 {
    era = 2
  } else if t >= 13 {
    era = 3
  }
  return

}

//Obtiene la carta que se debería jugar en el turno t, dado el estado de cartasJugadas.
func (c *Contexto) SimularTurno(t int, cartasJugadas [TURNOS]Carta) (cartaJugada Carta){
  pesoMaximo := float32(-1) //TODO: Esto se puede hacer acumulativo entre turnos y tener memoria
  eraActual := c.eraEnTurno(t)
  for _, carta := range c.cartasRestantes {
    if carta.Id != NULL && carta.SePuedeJugar(c.recursosDisponibles, cartasJugadas, eraActual, c.precioRecursos, c.comodinMateriaPrimaJugado, c.comodinManufacturaJugado) {
      pesoCarta := c.calcularPeso(cartasJugadas, c.cartasRestantes, carta)
      if pesoCarta > pesoMaximo {
        pesoMaximo = pesoCarta
        cartaJugada = carta
      }
    }
  }
  return
}

//Calcula los puntos segun todas las cartasJugadas
func (c *Contexto) calcularPuntos() {
  c.puntosTotales = 0
  /*for _, carta := range c.cartasJugadas {
    c.puntosTotales += carta.Id
  }*/

  escudos:=[3]int{0,0,0}
  puntosCiviles:=0
  cantidadGeometria:=0
  cantidadEscritura:=0
  cantidadRueda:=0
  for i, cartaJugada := range c.cartasJugadas{
    switch cartaJugada.Tipo{
    case CIVIL:
      //puntos civiles
      puntosCiviles+=cartaJugada.puntos
    case GEOMETRIA:
      //cantidad de cartas científicas de cada tipo
      cantidadGeometria+=1
    case RUEDA:
      cantidadRueda+=1
  case ESCRITURA:
    cantidadEscritura+=1
  case MILITAR:
    //cantidad de escudos que tiene en las batallas
      if i<BATALLA1{
        escudos[0]+=cartaJugada.puntos
        escudos[1]+=cartaJugada.puntos
        escudos[2]+=cartaJugada.puntos
      } else if i<BATALLA2{
        escudos[1]+=cartaJugada.puntos
        escudos[2]+=cartaJugada.puntos
      } else{
        escudos[2]+=cartaJugada.puntos
      }
    }
  }

  //puntos militares
  puntosMilitares := 0
  puntosContrincante:=[3]int{CONTRINCANTE1, CONTRINCANTE2, CONTRINCANTE3}
  for i, _ := range puntosContrincante{
    if escudos[i]<puntosContrincante[i]{
      puntosMilitares-=1
  } else if escudos[i]>puntosContrincante[i]{
    puntosMilitares+=puntosContrincante[i]
  }
  }

  puntosCientificosIguales := cantidadGeometria*cantidadGeometria + cantidadRueda*cantidadRueda + cantidadEscritura*cantidadEscritura
  puntosCientificosDiferentes:=cantidadEscritura
  cantidadesCientificas:=[3]int{cantidadEscritura,cantidadRueda,cantidadGeometria}
  for _,cantidad := range cantidadesCientificas{
    if cantidad<puntosCientificosDiferentes{
      puntosCientificosDiferentes = cantidad
    }
  }
  puntosCientificos := puntosCientificosIguales*PUNTOS_CIENTIFICOS_IGUALES + puntosCientificosDiferentes
  puntosMonedas := c.recursosDisponibles[MONEDA]/3

  puntosComerciales := c.CalcularPuntosComerciales()

  c.puntosTotales=puntosMilitares+puntosCiviles+puntosMonedas+puntosCientificos + puntosComerciales
  fmt.Println("### PUNTOS ###")
  fmt.Println("Puntos militares:", puntosMilitares)
  fmt.Println("Puntos civiles:", puntosCiviles)
  fmt.Println("Puntos Monedas:", puntosMonedas)
  fmt.Println("Puntos cientificos:", puntosCientificos)
  fmt.Println("Puntos comerciales:", puntosComerciales)
}


func (c *Contexto) CalcularPuntosComerciales() (puntos int){
  //TODO: NO ES OPTIMO PODRIA METERSE EN OTRO LADO; PERO BUE
  seJugoHaven := false
  seJugoChamber := false
  seJugoLighthouse := false
  for _, carta := range c.cartasJugadas {
    switch carta.Id {
      case HAVEN:
        seJugoHaven = true
      case CHAMBER:
        seJugoChamber = true
      case LIGHTHOUSE:
        seJugoLighthouse = true
    }
  }
  puntosPorMateriasPrimasAlFinal := 0
  puntosPorManufacturasAlFinal := 0
  puntosPorComercialesAlFinal := 0
  for _, carta := range c.cartasJugadas {
    switch carta.Tipo {
      case MATERIA_PRIMA:
        if seJugoHaven { puntosPorMateriasPrimasAlFinal++ }
      case MANUFACTURA:
        if seJugoChamber { puntosPorManufacturasAlFinal += 2 }
      case COMERCIAL:
        if seJugoLighthouse { puntosPorComercialesAlFinal += 2 }
    }
  }
  return puntosPorMateriasPrimasAlFinal + puntosPorManufacturasAlFinal + puntosPorComercialesAlFinal

}

func (c *Contexto) jugarCarta(cartaJugada Carta) {
  if cartaJugada.Id != NO_HACER_NADA {
    c.cartasRestantes[cartaJugada.Id].Id = NULL
  }
  if cartaJugada.Id == CARAVANSERY {
    c.comodinMateriaPrimaJugado = true
  }
  if cartaJugada.Id == FORUM {
    c.comodinManufacturaJugado = true
  }
  if cartaJugada.Id == MARKETPLACE {
    c.precioRecursos[CERAMICA] = 1
    c.precioRecursos[TELA] = 1
    c.precioRecursos[PAPIRO] = 1
  }
  if cartaJugada.Id == WEST_TRADING_POST {
    c.precioRecursos[LADRILLO] = 1
    c.precioRecursos[CEMENTO] = 1
    c.precioRecursos[ORO] = 1
    c.precioRecursos[MADERA] = 1
  }
  //ESTO se puede optimizar para que solo se llama
  materiasPrimasJugadas := 0
  manufacturasJugadas := 0
  comercialesJugadas := 0
  for _, unaCarta := range c.cartasJugadas {
    switch unaCarta.Tipo {
      case MATERIA_PRIMA:
        materiasPrimasJugadas++
      case MANUFACTURA:
        manufacturasJugadas++
      case COMERCIAL:
        comercialesJugadas++
    }
  }
  if cartaJugada.Id == VINEYARD || cartaJugada.Id == HAVEN {
    c.recursosDisponibles[MONEDA] = materiasPrimasJugadas
  }
  if cartaJugada.Id == BAZAR || cartaJugada.Id == CHAMBER {
    c.recursosDisponibles[MONEDA] = 2 * manufacturasJugadas
  }
  if cartaJugada.Id == LIGHTHOUSE {
    c.recursosDisponibles[MONEDA] = comercialesJugadas
  }

  c.recursosDisponibles[MONEDA] -= cartaJugada.monedasNecesarias
  recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}
  for r, _ := range recursos {
    c.recursosDisponibles[r] += cartaJugada.Produce[r]
  }
}

//Realiza la heurística de construcción
func (c *Contexto) ComenzarSimulacion() {
  for t:= 0; t < TURNOS;t++ {
    cartaJugada := c.SimularTurno(t, c.cartasJugadas)
    c.cartasJugadas[t] = cartaJugada
    c.jugarCarta(cartaJugada)
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
