package models

// DocentePlanTrabajo representa la informacion de un docente con una solicitud
// de revision de soporte de alguna actividad de su plan de trabajo
type DocentePlanTrabajo struct {
	ID                    int
	PrimerNombre          string
	SegundoNombre         string
	PrimerApellido        string  	
	SegundoApellido       string
}