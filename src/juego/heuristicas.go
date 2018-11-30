package juego

//Calcula el ponderable para un turno

func (c *Contexto) calcularPeso_id(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	peso = float32(unaCarta.Id)
	return
}

func (c *Contexto) calcularPeso_puntos(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
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

func (c *Contexto) calcularPeso_masBarata(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	peso = float32(0)
	w := float32(1.0/CANTIDAD_RECURSOS)
	var pesos [CANTIDAD_RECURSOS]float32
	for j:=0; j<CANTIDAD_RECURSOS; j++{
		pesos[j] = w;
	}
	for i := 0;i < CANTIDAD_RECURSOS;i++ {
		peso += float32(1.0/(float32(unaCarta.requiere[i]) * pesos[i]))
	}
	
	return
}

func (c *Contexto) calcularPeso(cartasJugadas [TURNOS]Carta, cartasRestantes [INDICE_MAXIMO_CARTA]Carta, unaCarta Carta) (peso float32) {
	return c.calcularPeso_puntos(cartasJugadas, cartasRestantes, unaCarta)
}
