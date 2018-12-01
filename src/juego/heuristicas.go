package juego

import (
		"math/rand"
)

//Calcula el ponderable para un turno

func (c *Contexto) calcularPeso_id(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
	peso = float32(unaCarta.Id)
	return
}

func (c *Contexto) calcularPeso_puntos(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
	//nombre = "PUNTOS"
	peso = float32(unaCarta.puntos)
	return
}

func (c *Contexto) calcularPeso_promedioDeProduce(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
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

func (c *Contexto) calcularPeso_puntosDeCartaQueLibera(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {

	if  cartasRestantes[unaCarta.edificioGratis].Id == -1 {
		peso = 0
	}else{
		peso = float32(cartasRestantes[unaCarta.edificioGratis].puntos)
	}

	return
}

func (c *Contexto) calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
	if  cartasRestantes[unaCarta.edificioGratis].Id == -1{
		peso = 0
	}else{
		peso =float32( c.calcularPeso_promedioDeProduce(cartasJugadas,cartasRestantes,cartasRestantes[unaCarta.edificioGratis], t))
	}
	return
}


func (c *Contexto) ORNITORRINCO_MAXIMO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
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

func (c *Contexto) SPAGHETTI(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
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

func (c *Contexto) CHOCOLATE(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
  //nombre = "CHOCOLATE"
  switch unaCarta.Tipo {
    case RUEDA, ESCRITURA, GEOMETRIA:
      peso = float32(1+PUNTOS_CIENTIFICOS_DIFERENTES)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case MILITAR:
      peso = float32(unaCarta.puntos)+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case CIVIL:
      peso = float32(unaCarta.puntos)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case COMERCIAL:
      peso = float32(unaCarta.Produce[MONEDA])+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    default:
      peso = MONEDAS_POR_NO_HACER_NADA/PUNTOS_POR_MONEDAS
  }
  return
}

func (c *Contexto) CHOCOLATE_CON_PAPAS_FRITAS(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
  //nombre = "CHOCOLATE"
  switch unaCarta.Tipo {
    case RUEDA, ESCRITURA, GEOMETRIA:
      peso = float32(1+PUNTOS_CIENTIFICOS_DIFERENTES)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case MILITAR:
      peso = float32(unaCarta.puntos)+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case CIVIL:
      peso = float32(unaCarta.puntos)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case COMERCIAL:
      peso = float32(unaCarta.Produce[MONEDA])+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    case NO_HACER_NADA:
    	peso = MONEDAS_POR_NO_HACER_NADA/PUNTOS_POR_MONEDAS
    case MARAVILLA:
    	peso = float32(unaCarta.puntos)*2.0+c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)-float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS
    default:
      peso = -float32(unaCarta.monedasNecesarias)/PUNTOS_POR_MONEDAS + c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)+2.0 * c.calcularPeso_promedioDeProduce(cartasJugadas, cartasRestantes, unaCarta,t)
  }
  return
}

func (c *Contexto) KOALA_CHICLOSO_PAPOTEADO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
	var recursosDisponibles [CANTIDAD_RECURSOS]int
	var recursosQueHabilita [CANTIDAD_RECURSOS]int
	necesitanRecurso := [CANTIDAD_RECURSOS]int{}
	recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}

	for _, carta := range cartasRestantes {
		for _ , r := range recursos {
			recursosDisponibles[r] += carta.Produce[r]
		}
	}

	for _ , r := range recursos {
		recursosQueHabilita[r] += unaCarta.Produce[r] * 3
	}

	for _, carta := range cartasRestantes {
		for _ , r := range recursos {
				necesitanRecurso[r] +=carta.requiere[r]
		}
	}

	cantidadRecursosNuevos := 0
	cantidadDeHabilitacionesNuevas := 0
	for _, r :=  range recursos {
		cantidadRecursosNuevos += recursosQueHabilita[r]
	}

	for _, r := range recursos {
		dif := recursosDisponibles[r] - necesitanRecurso[r]
		if dif  < 0 {
			if dif + recursosQueHabilita[r] >= 0 {
				cantidadDeHabilitacionesNuevas++
			}
		}
	}

	peso = float32(cantidadDeHabilitacionesNuevas) * cnn(4 - t) + float32(cantidadRecursosNuevos) * cnn(3-t) * 0.9
	peso += float32(unaCarta.puntos)
	switch unaCarta.Tipo {
		case GEOMETRIA, ESCRITURA, RUEDA:
			peso += float32(t) * 4
	}
	cantidadDeCartasHabilitadas := 0
	for _,carta:=range cartasRestantes{
		if carta.edificioGratis==unaCarta.Id{
			cantidadDeCartasHabilitadas++
		}
	}
	/*if t>=3{
		peso += 3000000.0*float32(cantidadDeCartasHabilitadas)
	}*/
	
  return
}

func cnn(a int) float32 {
	if a < 0 {
		return float32(0)
	} else {
		return float32(a)
	}
}

func (c *Contexto) CONDOR_ALPINO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, t int) (peso float32) {
	recursosQueHabilita := 0
	necesitanRecurso := [CANTIDAD_RECURSOS]int{}
	recursos := [CANTIDAD_RECURSOS]int{LADRILLO, CEMENTO, ORO, MADERA, CERAMICA, TELA, PAPIRO, MONEDA}

	for _ , r := range recursos {
		recursosQueHabilita += unaCarta.Produce[r]
	}

	for _, carta := range cartasRestantes {
		for _ , r := range recursos {
			if carta.Produce[r] > 0{
				necesitanRecurso[r]++
			}
		}
	}

	prom := float32(0)
	for _, e := range necesitanRecurso {
		prom += float32(e)
	}

	prom /= CANTIDAD_RECURSOS

	switch unaCarta.Tipo {
		case RUEDA, ESCRITURA, GEOMETRIA:
			peso = 0.3 * float32(t)
		default:
			ruido := rand.Float32()
			if recursosQueHabilita != 0 {
				peso = (float32(recursosQueHabilita) - ruido * prom)/float32(recursosQueHabilita)
				if peso < 0 {
					 peso = ruido * ruido * 0.3 * (3-float32(t))
					 if peso < 0 {
						 peso = ruido *  0.05
					 }
				}
			} else {
				peso =ruido *  0.01 * float32(t)
			}
	}

	return
}

func (c *Contexto) calcularPeso(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta, nroHeuristica int, t int) (peso float32) {
	/*nombre := ""
	peso, nombre = c.CHOCOLATE(cartasJugadas, cartasRestantes, unaCarta)
	fmt.Println("Calculado con heurísitca:", nombre)
	return*/
	switch nroHeuristica{
	case HEURISTICA_ID:
		peso = c.calcularPeso_id(cartasJugadas, cartasRestantes, unaCarta,t)
	case HEUTISTICA_PUNTOS:
		peso = c.calcularPeso_puntos(cartasJugadas, cartasRestantes, unaCarta,t)
	case HEURISTICA_PROMEDIO_DE_PRODUCE:
		peso = c.calcularPeso_promedioDeProduce(cartasJugadas, cartasRestantes, unaCarta,t)
	case HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA:
		peso = c.calcularPeso_puntosDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)
	case HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA:
		peso = c.calcularPeso_promedioProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta,t)
	case ORNITORRINCO_MAXIMO:
		peso = c.ORNITORRINCO_MAXIMO(cartasJugadas, cartasRestantes, unaCarta,t)
	case CONDOR_ALPINO:
		peso = c.CONDOR_ALPINO(cartasJugadas, cartasRestantes, unaCarta,t)
	case SPAGHETTI:
		peso = c.SPAGHETTI(cartasJugadas, cartasRestantes, unaCarta,t)
	case CHOCOLATE:
		peso = c.CHOCOLATE(cartasJugadas, cartasRestantes, unaCarta,t)
	case KOALA_CHICLOSO_PAPOTEADO:
		peso = c.KOALA_CHICLOSO_PAPOTEADO(cartasJugadas, cartasRestantes, unaCarta, t)
	case CHOCOLATE_CON_PAPAS_FRITAS:
		peso = c.CHOCOLATE_CON_PAPAS_FRITAS(cartasJugadas, cartasRestantes, unaCarta, t)
	}
	return
}
