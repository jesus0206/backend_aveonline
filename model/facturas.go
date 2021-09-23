package model

// Promocion exported
type Factura struct {
	ID           int            `json:"id"`
	Fecha_Crear  string         `json:"fecha_crear"`
	Pago_Total   string         `json:"pago_total"`
	Promocion    *Promocion     `json:"promocion"`
	PromocionID  int            `json:"-"`
	Medicamentos []*Medicamento `json:"medicacmentos"`
}

// FacturaSimularDto exported
type FacturaSimularDto struct {
	Fecha_Compra   string `json:"fecha_compra"`
	MedicamentosID []int  `json:"id_medicacmentos"`
}

// FacturaDto exported
type FacturaDto struct {
	Fecha_Inicio string `json:"fecha_inicio"`
	Fecha_Fin    string `json:"fecha_fin"`
}

// FacturaCreateDto exported
type FacturaCreateDto struct {
	ID             int     `json:"id"`
	Fecha_Crear    string  `json:"fecha_crear"`
	Pago_Total     float32 `json:"pago_total"`
	PromocionID    int     `json:"id_promocion"`
	MedicamentosID []int   `json:"id_medicamentos"`
}
