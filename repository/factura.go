package repository

import (
	"jesus.tn79/aveonline/model"
)

func (repo Repository) GetFacturas() ([]*model.Factura, error) {
	var facturas []*model.Factura
	rows, err := repo.db.Raw("SELECT id,fecha_crear,pago_total,promocion_id FROM factura").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		factura := &model.Factura{}
		_ = rows.Scan(&factura.ID, &factura.Fecha_Crear, &factura.Pago_Total, &factura.PromocionID)
		facturas = append(facturas, factura)
	}
	for i, item := range facturas {
		promocion, err := repo.GetPromocion(item.PromocionID)
		if err != nil {
			continue
		}
		facturas[i].Promocion = promocion
		medicamentos, err := repo.GetFacturaMedicamentos(item.ID)
		if err != nil {
			continue
		}
		facturas[i].Medicamentos = medicamentos
	}

	return facturas, err
}

func (repo Repository) CreateFactura(data model.FacturaCreateDto) (*string, error) {
	_ = repo.db.Exec(`INSERT INTO factura(fecha_crear,pago_total,promocion_id)VALUES(?,?,?)`, data.Fecha_Crear, data.Pago_Total, data.PromocionID)

	for _, id := range data.MedicamentosID {
		err := repo.CreateFacturaItem(5, id)
		if err != nil {
			continue
		}
	}
	message := "factura creado."
	return &message, nil
}

func (repo Repository) CreateFacturaItem(factura_id, medicamento_id int) error {
	_, err := repo.db.Exec("INSERT INTO factura_items(factura_id,medicamento_id)VALUES(?,?)", factura_id, medicamento_id).Rows()
	if err != nil {
		return err
	}
	return nil
}