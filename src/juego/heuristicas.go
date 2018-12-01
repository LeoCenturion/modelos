package juego

/*import (
	"fmt"
)*/

//Calcula el ponderable para un turno

func (c *Contexto) calcularPeso_id(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	peso = float32(unaCarta.Id)
	return
}

func (c *Contexto) calcularPeso_puntos(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	//nombre = "PUNTOS"
	peso = float32(unaCarta.puntos)
	return
}

func (c *Contexto) calcularPeso_promedioDeProduce(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	peso = float32(0)
	w := float32(1.0/CANTIDAD_RECURSOS)
	var pesos [CANTIDAD_RECURSOS]float32
	for j:=0; j<CANTIDAD_RECURSOS; j++{
		pesos[j] = w;
	}
	for i := 0;i < CANTIDAD_RECURSOS;i++ {
		peso += float32(unaCarta.Produce[i]) * pesos[i]
	}
	return
}

func (c *Contexto) calcularPeso_puntosDeCartaQueLibera(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {

	if  cartasRestantes[unaCarta.edificioGratis].Id == -1 {
		peso = 0
	}else{
		peso = float32(cartasRestantes[unaCarta.edificioGratis].puntos)
	}

	return
}

func (c *Contexto) calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	if  cartasRestantes[unaCarta.edificioGratis].Id == -1{
		peso = 0
	}else{
		peso =float32( c.calcularPeso_promedioDeProduce(cartasJugadas,cartasRestantes,cartasRestantes[unaCarta.edificioGratis]))
	}
	return
}


func (c *Contexto) ORNITORRINCO_MAXIMO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	//nombre = "ORNITORRINCO_MAXIMO"
	switch unaCarta.Tipo {
		case RUEDA, ESCRITURA, GEOMETRIA:
			peso = 0.6
		case MILITAR:
			peso = 0.1
		case CIVIL:
			peso = 0.2
		case COMERCIAL:
			peso = 0.4
		default:
			peso = 0.1
	}
	return
}

func (c *Contexto) CONDOR_ALPINO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	//nombre = "CONDOR_ALPINO"
	switch unaCarta.Tipo {
		case RUEDA, ESCRITURA, GEOMETRIA:
			peso = 0.6
		case MILITAR:
			peso = 0.1
		case CIVIL:
			peso = 0.2
		case COMERCIAL:
			peso = 0.4
		default:
			peso = 0.1
	}
	return
}

func (c *Contexto) SPAGHETTI(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
  //nombre = "SPAGHETTI"
  switch unaCarta.Tipo {
    case RUEDA, ESCRITURA, GEOMETRIA:
      peso = float32(1+PUNTOS_CIENTIFICOS_DIFERENTES)
    case MILITAR:
      peso = float32(unaCarta.puntos)
    case CIVIL:
      peso = float32(unaCarta.puntos)
    case COMERCIAL:
      peso = float32(unaCarta.Produce[MONEDA])
    default:
      peso = 1
  }
  return
}

func (c *Contexto) CHOCOLATE(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
  //nombre = "CHOCOLATE"
  switch unaCarta.Tipo {
    case RUEDA, ESCRITURA, GEOMETRIA:
      peso = float32(1+PUNTOS_CIENTIFICOS_DIFERENTES)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case MILITAR:
      peso = float32(unaCarta.puntos)+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case CIVIL:
      peso = float32(unaCarta.puntos)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case COMERCIAL:
      peso = float32(unaCarta.Produce[MONEDA])+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    default:
      peso = MONEDAS_POR_NO_HACER_NADA/PUNTOS_POR_MONEDAS
  }
  return
}

func (c *Contexto) calcularPeso(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, nroHeuristica int) (peso float32) {
	/*nombre := ""
	peso, nombre = c.CHOCOLATE(cartasJugadas, cartasRestantes, unaCarta)
	fmt.Println("Calculado con heurísitca:", nombre)
	return*/
	switch nroHeuristica{
	case HEURISTICA_ID:
		peso = c.calcularPeso_id(cartasJugadas, cartasRestantes, unaCarta)
	case HEUTISTICA_PUNTOS:
		peso = c.calcularPeso_puntos(cartasJugadas, cartasRestantes, unaCarta)
	case HEURISTICA_PROMEDIO_DE_PRODUCE:
		peso = c.calcularPeso_promedioDeProduce(cartasJugadas, cartasRestantes, unaCarta)
	case HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA:
		peso = c.calcularPeso_puntosDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
	case HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA:
		peso = c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
	case ORNITORRINCO_MAXIMO:
		peso = c.ORNITORRINCO_MAXIMO(cartasJugadas, cartasRestantes, unaCarta)
	case CONDOR_ALPINO:
		peso = c.CONDOR_ALPINO(cartasJugadas, cartasRestantes, unaCarta)
	case SPAGHETTI:
		peso = c.SPAGHETTI(cartasJugadas, cartasRestantes, unaCarta)
	case CHOCOLATE:
		peso = c.CHOCOLATE(cartasJugadas, cartasRestantes, unaCarta)
	}
	return
}
