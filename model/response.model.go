package model

// A Response exported
type Response struct {
	Error   string `json:"error"`
	Mensaje string `json:"mensaje"`
}

// ResponseError response
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseEmpty response
type ResponseEmpty struct {
	Data string `json:"data"`
}

// MedicamentosResponse response
type MedicamentosResponse struct {
	Data []*Medicamento `json:"data"`
}

// // OficinaSala model response
// type OficinaSala struct {
// 	Oficina Oficina `json:"oficina"`
// 	Salas   []*Sala `json:"salas"`
// }

// // ResponseOficinaSala model response
// type ResponseOficinaSala struct {
// 	Body OficinaSala `json:"body"`
// }

// // ResponseTokenUser model response
// type ResponseTokenUser struct {
// 	Body string `json:"body"`
// }
