package juego

const CANTIDAD_RECURSOS=8


//Constantes para acceder a columnas csvs, y los recursos
const LADRILLO = 0
const CEMENTO = 1
const ORO = 2
const MADERA = 3
const CERAMICA = 4
const TELA = 5
const PAPIRO = 6
const MONEDA = 7
const OFFSET_REQUIERE = 8
const ID = 16
const TIPO = 17
const ERA = 18
const PUNTOS = 19
const EDIFICIO_GRATIS = 20
const CARTA_REQUERIDA = 21
const NOMBRE = 22

//Tipos de cartas
const GEOMETRIA = 0
const RUEDA = 1
const ESCRITURA = 2
const CIVIL = 3
const MATERIA_PRIMA = 4
const MANUFACTURA = 5
const COMERCIAL = 6
const MILITAR = 7
const MARAVILLA = 8
const NO_HACER_NADA = 55

//misc
const NINGUNA = 0
const TURNOS=18
const INDICE_MAXIMO_CARTA = 91
const NULL = -1
const ID_CARTA_COMODIN = 61
const CUALQUIERA = -1
const PRECIO_INICIAL_RECURSO = 2
const MARKETPLACE = 75
const WEST_TRADING_POST = 74
const HAVEN = 88
const CHAMBER = 89
const LIGHTHOUSE = 90
const VINEYARD = 81
const BAZAR = 82
const CARAVANSERY = 83
const FORUM = 80
//turno de batallas
const BATALLA1 = 6
const BATALLA2 = 12
const BATALLA3 = 18

//puntos militares del contrincante para las batallas
const CONTRINCANTE1 = 1
const CONTRINCANTE2 = 3
const CONTRINCANTE3 = 5

//puntos científicos por cada grupo de 3 tipos distintos
const PUNTOS_CIENTIFICOS_DIFERENTES = 7

const PUNTOS_POR_MONEDAS = 3
const MONEDAS_POR_NO_HACER_NADA = 3

//identificadores de las heuristicas
const HEURISTICA_ID = 0
const HEUTISTICA_PUNTOS = 1
const HEURISTICA_PROMEDIO_DE_PRODUCE = 2
const HEURISTICA_PUNTOS_DE_CARTA_QUE_LIBERA = 3
const HEURISTICA_PROMEDIO_PRODUCE_DE_CARTA_QUE_LIBERA = 4
const ORNITORRINCO_MAXIMO = 5
const CONDOR_ALPINO = 6
const SPAGHETTI = 7
const CHOCOLATE = 8

func recursoEsMateriaPrima (recurso int) (esMateriaPrima bool) {
  esMateriaPrima = recurso < CERAMICA
  return
}
