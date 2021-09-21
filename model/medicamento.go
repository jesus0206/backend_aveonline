package model

// Medicamento exported
type Medicamento struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Precio    float32 `json:"precio"`
	Ubicacion string  `json:"ubicacion"`
}
