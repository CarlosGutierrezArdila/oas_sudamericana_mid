package models

import (
	"time"
)

type Partido struct {
	Id                int
	EquipoLocalId     int
	EquipoVisitanteId int
	PuntajeLocal      int
	PuntajeVisitante  int
	Fecha             time.Time
	CiudadId          int
}
