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
  CartasJugadas [TURNOS]Carta
  cartasDisponibles, cartasRestantes [INDICE_MAXIMO_CARTA]Carta
	cartasJugablesEnTurno [TURNOS][INDICE_MAXIMO_CARTA]Carta
  recursosDisponibles, precioRecursos [CANTIDAD_RECURSOS]int
  PuntosTotales int
  comodinMateriaPrimaJugado, comodinManufacturaJugado bool
  DetalleDePuntos string
}

func (c Contexto) GetEstado() ([TURNOS]Carta, [INDICE_MAXIMO_CARTA]Carta,[INDICE_MAXIMO_CARTA]Carta, [CANTIDAD_RECURSOS]int, [CANTIDAD_RECURSOS]int, bool, bool, [TURNOS][INDICE_MAXIMO_CARTA]Carta) {
	return c.CartasJugadas, c.cartasDisponibles, c.cartasRestantes, c.recursosDisponibles, c.precioRecursos, c.comodinManufacturaJugado, c.comodinMateriaPrimaJugado, c.cartasJugablesEnTurno
}

func (c *Contexto) Init(cartasJugadas [TURNOS]Carta, cartasDisponibles, cartasRestantes[INDICE_MAXIMO_CARTA]Carta, recursosDisponibles, precioRecursos [CANTIDAD_RECURSOS]int,comodinManufacturaJugado, comodinMateriaPrimaJugado bool, cartasJugablesEnTurno [TURNOS][INDICE_MAXIMO_CARTA]Carta) {
	c.CartasJugadas = cartasJugadas
	c.cartasDisponibles = cartasDisponibles
	c.cartasRestantes = cartasRestantes
	c.recursosDisponibles = recursosDisponibles
	c.precioRecursos = precioRecursos
	c.comodinManufacturaJugado = comodinManufacturaJugado
	c.comodinMateriaPrimaJugado = comodinMateriaPrimaJugado
	c.cartasJugablesEnTurno = cartasJugablesEnTurno
}

func (c *Contexto) Resetear(){
	for i, _ := range c.CartasJugadas {
		c.CartasJugadas[i].Id = NULL
	}

	for i := 0; i < TURNOS; i++ {
		for j := 0; j < INDICE_MAXIMO_CARTA; j ++ {
			carta := Carta{}
			carta.Id = NULL
			c.cartasJugablesEnTurno[i][j] = carta
		}
	}

	for i, _ := range c.cartasDisponibles {
		c.cartasDisponibles[i].Id = NULL
	}

	for i, _ := range c.precioRecursos {
		c.precioRecursos[i] = PRECIO_INICIAL_RECURSO
		c.recursosDisponibles[i] = 0
	}

	c.PuntosTotales = 0
	c.comodinMateriaPrimaJugado = false
	c.comodinManufacturaJugado = false
	c.recursosDisponibles[MONEDA] = MONEDAS_POR_NO_HACER_NADA
}

//Carga de un csv todas las cartas disponibles
func (c *Contexto) CargarCartas(archivo string) {
	c.Resetear()

	data, _ := ioutil.ReadFile(archivo)
	r := csv.NewReader(strings.NewReader(string(data)))

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	/*fmt.Println("### Cabecera ### ")
	fmt.Println(records[0])*/
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
	produce[MONEDA] = MONEDAS_POR_NO_HACER_NADA
	cartaNoHacerNada.Init(NO_HACER_NADA,NO_HACER_NADA, CUALQUIERA, 0, 0,0, "No hacer nada", produce, requiere)
	c.cartasDisponibles[NO_HACER_NADA] = cartaNoHacerNada
	c.cartasRestantes = c.cartasDisponibles //TODO: esto tal vez copia la referencia del array, por ahora no nos jode igual por que no se si vamos a usar cartas disponibles y restantes
}

//Determina la era a la que pertence el turno t
func (c Contexto) eraEnTurno(t int) (era int){
	era = 1
	if t+1 >= TURNO_QUE_EMPIEZA_ERA_2 && t+1 <= TURNO_QUE_FINALIZA_ERA_2 {
		era = 2
	} else if t+1 >= TURNO_QUE_EMPIEZA_ERA_3 {
		era = 3
	}
	return

}

//Obtiene la carta que se debería jugar en el turno t, dado el estado de cartasJugadas.
func (c *Contexto) SimularTurno(t int, cartasJugadas [TURNOS]Carta, nroHeuristica int) (cartaJugada Carta){
	pesoMaximo := float32(-1) //TODO: Esto se puede hacer acumulativo entre turnos y tener memoria
	eraActual := c.eraEnTurno(t)
	for _, carta := range c.cartasRestantes {
		isIdNull := carta.Id == NULL
		sePuedeJugar := carta.SePuedeJugar(c.recursosDisponibles, cartasJugadas, eraActual, c.precioRecursos, c.comodinMateriaPrimaJugado, c.comodinManufacturaJugado)
		fmt.Println("se puede jugar", sePuedeJugar)
		fmt.Println("is null", isIdNull)
		if  !isIdNull && sePuedeJugar  {
			c.cartasJugablesEnTurno[t][carta.Id] = carta
			pesoCarta := c.calcularPeso(cartasJugadas, c.cartasRestantes, carta, nroHeuristica, t)
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
	c.PuntosTotales = 0
	/*for _, carta := range c.CartasJugadas {
    c.PuntosTotales += carta.Id
  }*/

	escudos:=[3]int{0,0,0}
	puntosCiviles:=0
	cantidadGeometria:=0
	cantidadEscritura:=0
	cantidadRueda:=0

	seJugoHaven := false
	seJugoChamber := false
	seJugoLighthouse := false

	puntosPorMateriasPrimasAlFinal := 0
	puntosPorManufacturasAlFinal := 0
	puntosPorComercialesAlFinal := 0

	for i, cartaJugada := range c.CartasJugadas{
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
		case MATERIA_PRIMA:
			if seJugoHaven { puntosPorMateriasPrimasAlFinal++ }
		case MANUFACTURA:
			if seJugoChamber { puntosPorManufacturasAlFinal += 2 }
		case COMERCIAL:
			if seJugoLighthouse { puntosPorComercialesAlFinal += 2 }
		}
		switch cartaJugada.Id {
		case HAVEN:
			seJugoHaven = true
		case CHAMBER:
			seJugoChamber = true
		case LIGHTHOUSE:
			seJugoLighthouse = true
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
	for _,cantidad := range cantidadesCientificas {
		if cantidad<puntosCientificosDiferentes{
			puntosCientificosDiferentes = cantidad
		}
	}
	puntosCientificos := puntosCientificosIguales + puntosCientificosDiferentes*PUNTOS_CIENTIFICOS_DIFERENTES

  	puntosMonedas := c.recursosDisponibles[MONEDA]/3

	puntosComerciales := puntosPorMateriasPrimasAlFinal + puntosPorManufacturasAlFinal + puntosPorComercialesAlFinal

	c.PuntosTotales=puntosMilitares+puntosCiviles+puntosMonedas+puntosCientificos + puntosComerciales

	c.DetalleDePuntos = fmt.Sprintf("\nPuntos Militares: %d \nPuntos civiles: %d\nPuntos monedas: %d\nPuntos cientificos totales: %d\nPuntos cientificos iguales: %d\nPuntos cientificos diferentes: %d\nCantidad de cartas de tipo escritura: %d\nCantidad de cartas de tipo rueda: %d\nCantidad de cartas de tipo geometria: %d\nPuntos comerciales: %d\n", puntosMilitares, puntosCiviles, puntosMonedas, puntosCientificos, puntosCientificosIguales, puntosCientificosDiferentes*PUNTOS_CIENTIFICOS_DIFERENTES, cantidadEscritura, cantidadRueda, cantidadGeometria, puntosComerciales)
	/*fmt.Println("### PUNTOS ###")
	fmt.Println("Puntos militares:", puntosMilitares)
	fmt.Println("Puntos civiles:", puntosCiviles)
	fmt.Println("Puntos Monedas:", puntosMonedas)
	fmt.Println("Puntos cientificos totales:", puntosCientificos)
	fmt.Println("Puntos cientificos iguales:", puntosCientificosIguales)
	fmt.Println("Puntos cientificos diferentes:", puntosCientificosDiferentes*PUNTOS_CIENTIFICOS_DIFERENTES)
	fmt.Println("Cantidad de cartas de tipo escritura:", cantidadEscritura)
	fmt.Println("Cantidad de cartas de tipo rueda:", cantidadRueda)
	fmt.Println("Cantidad de cartas de tipo geometria:", cantidadGeometria)
	fmt.Println("Puntos comerciales:", puntosComerciales)*/
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
  for _, unaCarta := range c.CartasJugadas {
		if unaCarta.Id != NULL {
	  switch unaCarta.Tipo {
      case MATERIA_PRIMA:
        materiasPrimasJugadas++
      case MANUFACTURA:
        manufacturasJugadas++
      case COMERCIAL:
        comercialesJugadas++
    }
	}
  }
  if cartaJugada.Id == VINEYARD || cartaJugada.Id == HAVEN {
    c.recursosDisponibles[MONEDA] += materiasPrimasJugadas
  }
  if cartaJugada.Id == BAZAR || cartaJugada.Id == CHAMBER {
    c.recursosDisponibles[MONEDA] += 2 * manufacturasJugadas
  }
  if cartaJugada.Id == LIGHTHOUSE {
    c.recursosDisponibles[MONEDA] += comercialesJugadas
  }

  c.recursosDisponibles[MONEDA] -= cartaJugada.monedasNecesarias
  recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}
  for r, _ := range recursos {
    c.recursosDisponibles[r] += cartaJugada.Produce[r]
  }
}

//Realiza la heurística de construcción
func (c *Contexto) ComenzarSimulacion(nroHeuristica, turno int) {
  for t := turno; t < TURNOS;t++ {
	  cartaJugada := c.SimularTurno(t, c.CartasJugadas, nroHeuristica)
	  if cartaJugada.Id == NULL  {  log.Fatal("se pudre todo" )}
//	  fmt.Println("Carta jugada", cartaJugada)
    c.CartasJugadas[t] = cartaJugada
    c.jugarCarta(cartaJugada)
    //fmt.Println("Monedas turno", t,":", c.recursosDisponibles[MONEDA])
  }
  c.calcularPuntos()
}

func (c *Contexto) RejugarUltimaCarta(turnoCambio, cartaCambio int) {
	c.deshacerJugadasHasta(turnoCambio)
	c.CartasJugadas[turnoCambio] = c.cartasJugablesEnTurno[turnoCambio][cartaCambio]
	c.jugarCarta(c.CartasJugadas[turnoCambio])
}

func (c *Contexto) deshacerJugadasHasta(turnoCambio int) {
	for t := TURNOS -1 ; t >= turnoCambio; t-- {
		c.deshacerJugadaCarta(c.CartasJugadas[t])
	}
}

func (c *Contexto) deshacerJugadaCarta(cartaJugada Carta) {
	c.cartasRestantes[cartaJugada.Id].Id = cartaJugada.Id
  if cartaJugada.Id == CARAVANSERY {
    c.comodinMateriaPrimaJugado = false
  }
  if cartaJugada.Id == FORUM {
    c.comodinManufacturaJugado = false
  }
  if cartaJugada.Id == MARKETPLACE {
    c.precioRecursos[CERAMICA] = PRECIO_INICIAL_RECURSO
    c.precioRecursos[TELA] = PRECIO_INICIAL_RECURSO
    c.precioRecursos[PAPIRO] = PRECIO_INICIAL_RECURSO
  }
  if cartaJugada.Id == WEST_TRADING_POST {
    c.precioRecursos[LADRILLO] = PRECIO_INICIAL_RECURSO
    c.precioRecursos[CEMENTO] = PRECIO_INICIAL_RECURSO
    c.precioRecursos[ORO] = PRECIO_INICIAL_RECURSO
    c.precioRecursos[MADERA] = PRECIO_INICIAL_RECURSO
  }
  //ESTO se puede optimizar para que solo se llama
  materiasPrimasJugadas := 0
  manufacturasJugadas := 0
  comercialesJugadas := 0
  for _, unaCarta := range c.CartasJugadas {
		if unaCarta.Id != NULL {
		switch unaCarta.Tipo {
      case MATERIA_PRIMA:
        materiasPrimasJugadas++
      case MANUFACTURA:
        manufacturasJugadas++
      case COMERCIAL:
        comercialesJugadas++
    }
	}
  }
  if cartaJugada.Id == VINEYARD || cartaJugada.Id == HAVEN {
    c.recursosDisponibles[MONEDA] -= materiasPrimasJugadas
  }
  if cartaJugada.Id == BAZAR || cartaJugada.Id == CHAMBER {
    c.recursosDisponibles[MONEDA] -= 2 * manufacturasJugadas
  }
  if cartaJugada.Id == LIGHTHOUSE {
    c.recursosDisponibles[MONEDA] -= comercialesJugadas
  }

  c.recursosDisponibles[MONEDA] += cartaJugada.monedasNecesarias
  recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}
  for r, _ := range recursos {
    c.recursosDisponibles[r] -= cartaJugada.Produce[r]
  }
}

//devuelve el nombre de la heuristica de la cual se pasa el nro por parametro
func (c *Contexto) ObtenerNombreHeuristica(heuristica int) string{
	switch heuristica{
	case HEURISTICA_ID:
		return "HEURISTICA_ID"
	case HEUTISTICA_PUNTOS:
		return "HEUTISTICA_PUNTOS"
	case HEURISTICA_PROMEDIO_DE_PRODUCE:
		return "HEURISTICA_PROMEDIO_DE_PRODUCE"
	case HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA:
		return "HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA"
	case HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA:
		return "HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA"
	case ORNITORRINCO_MAXIMO:
		return "ORNITORRINCO_MAXIMO"
	case CONDOR_ALPINO:
		return "CONDOR_ALPINO"
	case SPAGHETTI:
		return "SPAGHETTI"
	case CHOCOLATE:
		return "CHOCOLATE"
	case KOALA_CHICLOSO_PAPOTEADO:
		return "KOALA_CHICLOSO_PAPOTEADO"
	case CHOCOLATE_CON_PAPAS_FRITAS:
		return "CHOCOLATE_CON_PAPAS_FRITAS"
	}
	return "NO SE QUE HEURISTICA ES!!!"
}

//devuelve el nombre del tipo del nro que se pasa por parametro
func (c Contexto) obtenerNombreTipo(tipo int) string{
	switch tipo{
	case GEOMETRIA:
		return "Cientifica tipo Geometria"
	case RUEDA:
		return "Cientifica tipo Rueda"
	case ESCRITURA:
		return "Cientifica tipo Escritura"
	case CIVIL:
		return "Civil"
	case MATERIA_PRIMA:
		return "Materia Prima"
	case MANUFACTURA:
		return "Manufactura"
	case COMERCIAL:
		return "Comercial"
	case MILITAR:
		return "Militar"
	case MARAVILLA:
		return "Maravilla"
	case NO_HACER_NADA:
		return "No hacer nada"
	}
	return "NO SE DE QUE TIPO ES!!!"
}

//muestra los resultados de la simulación (las cartas que se deciden jugar)
func (c Contexto) MostrarResultados(cartasJugadas [TURNOS]Carta, detalleDePuntos string, puntosTotales int) {
	fmt.Println("### Resultados ###")
	nombreTipo := ""
	for i, e := range cartasJugadas {
		nombreTipo = c.obtenerNombreTipo(e.Tipo)
		fmt.Println("Carta jugada en turno", i+1, ":",e.Nombre,"-",nombreTipo,"-", e.Id)
	}
	fmt.Println(detalleDePuntos)
	fmt.Println("Puntos obtenidos:", puntosTotales)
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
