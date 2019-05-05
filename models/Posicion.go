package models

type Posicion struct {
	Id              int
	EquipoId        int
	Puesto          int
	PartidoJugado   int
	PartidoGanado   int
	PartidoEmpatado int
	PartidoPerdido  int
	CantidadPuntaje int
}
