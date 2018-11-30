package juego

import (
	"fmt"
)

//Calcula el ponderable para un turno

func (c *Contexto) calcularPeso_id(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	peso = float32(unaCarta.Id)
	return
}

func (c *Contexto) calcularPeso_puntos(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32, nombre string) {
	nombre = "PUNTOS"
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

func (c *Contexto) calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	if  cartasRestantes[unaCarta.edificioGratis].Id == -1{
		peso = 0
	}else{
		peso =float32( c.calcularPeso_promedioDeProduce(cartasJugadas,cartasRestantes,cartasRestantes[unaCarta.edificioGratis]))
	}
	return
}


func (c *Contexto) ORNITORRINCO_MAXIMO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32, nombre string) {
	nombre = "ORNITORRINCO_MAXIMO"
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

func (c *Contexto) CONDOR_ALPINO(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32, nombre string) {
	nombre = "CONDOR_ALPINO"
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

func (c *Contexto) SPAGHETTI(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32, nombre string) {
  nombre = "ORNITORRINCO_MAXIMO"
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

func (c *Contexto) CHOCOLATE(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32, nombre string) {
  nombre = "ORNITORRINCO_MAXIMO"
  switch unaCarta.Tipo {
    case RUEDA, ESCRITURA, GEOMETRIA:
      peso = float32(1+PUNTOS_CIENTIFICOS_DIFERENTES)+c.calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
    case MILITAR:
      peso = float32(unaCarta.puntos)+c.calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
    case CIVIL:
      peso = float32(unaCarta.puntos)+c.calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
    case COMERCIAL:
      peso = float32(unaCarta.Produce[MONEDA])+c.calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
    default:
      peso = 1+c.calcularPeso_promedoProduceDeCartaQueLibera(cartasJugadas, cartasRestantes, unaCarta)
  }
  return
}

func (c *Contexto) calcularPeso(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	nombre := ""
	peso, nombre = c.CHOCOLATE(cartasJugadas, cartasRestantes, unaCarta)
	fmt.Println("Calculado con heurísitca:", nombre)
	return
}
