package repository

import (
	"jesus.tn79/aveonline/model"
)

func (repo Repository) GetMedicamentos() ([]*model.Medicamento, error) {
	var medicamentos []*model.Medicamento
	rows, err := repo.db.Raw("SELECT id,nombre,precio,ubicacion FROM medicamento").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		medicamento := &model.Medicamento{}
		_ = rows.Scan(&medicamento.ID, &medicamento.Nombre, &medicamento.Precio, &medicamento.Ubicacion)
		medicamentos = append(medicamentos, medicamento)
	}
	return medicamentos, err
}

func (repo Repository) GetFacturaMedicamentos(factura_id int) ([]*model.Medicamento, error) {
	var medicamentos []*model.Medicamento
	rows, err := repo.db.Raw(`
			SELECT  m.id, m.nombre, m.precio, m.ubicacion
			FROM factura_items fi
			INNER JOIN medicamento m ON fi.medicamento_id=m.id
			WHERE fi.factura_id=?`, factura_id).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		medicamento := &model.Medicamento{}
		_ = rows.Scan(&medicamento.ID, &medicamento.Nombre, &medicamento.Precio, &medicamento.Ubicacion)
		medicamentos = append(medicamentos, medicamento)
	}
	return medicamentos, err
}

func (repo Repository) CreateMedicamento(data model.Medicamento) (*string, error) {
	sql := repo.db.Exec("INSERT INTO medicamento(nombre,precio,ubicacion)VALUES(?,?,?)", data.Nombre, data.Precio, data.Ubicacion)
	if sql.Error != nil {
		return nil, sql.Error
	}
	message := "medicamento creado."
	return &message, nil
}
