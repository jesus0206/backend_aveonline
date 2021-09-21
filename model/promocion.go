package model

// Promocion exported
type Promocion struct {
	ID           int     `json:"id"`
	Descripcion  string  `json:"descripcion"`
	Porcentaje   float32 `json:"porcentaje"`
	Fecha_Inicio string  `json:"fecha_inicio"`
	Fecha_Fin    string  `json:"fecha_fin"`
}
