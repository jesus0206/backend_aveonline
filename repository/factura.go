package repository

import (
	"errors"
	"strings"

	"jesus.tn79/aveonline/model"
)

func (repo Repository) GetFacturas(query model.FacturaDto) ([]*model.Factura, error) {
	var facturas []*model.Factura
	rows, err := repo.db.Raw("SELECT id,fecha_crear,pago_total,promocion_id FROM factura WHERE (fecha_crear BETWEEN  ? AND ?)", query.Fecha_Inicio, query.Fecha_Fin).Rows()
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
	var factura_id *int
	rows, err := repo.db.Raw(`INSERT INTO factura(promocion_id,pago_total)VALUES(?,?)RETURNING id`, data.PromocionID, data.Pago_Total).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		_ = rows.Scan(&factura_id)
	}
	if factura_id == nil {
		return nil, errors.New("no se genero el id de la factura.")
	}
	for _, id := range data.MedicamentosID {
		err := repo.create_factura_tem(*factura_id, id)
		if err != nil {
			continue
		}
	}
	err = repo.updatePagoTotalFacturaItem(*factura_id)
	if err != nil {
		return nil, err
	}
	message := "factura creado."
	return &message, nil
}

func (repo Repository) SimularFactura(fecha, medicamentos_id string) (*float32, error) {
	precio := repo.precio_medicamentos(medicamentos_id)

	promocion, err := repo.GetPromocionPorFecha(fecha)
	if err != nil {
		return nil, err
	}
	result := *precio - (*precio * promocion.Porcentaje / 100)
	return &result, nil
}

func (repo Repository) create_factura_tem(factura_id, medicamento_id int) error {
	_, err := repo.db.Exec("INSERT INTO factura_items(factura_id,medicamento_id)VALUES(?,?)", factura_id, medicamento_id).Rows()
	if err != nil {
		return err
	}
	return nil
}

func (repo Repository) updatePagoTotalFacturaItem(factura_id int) error {
	var pago_total *float32
	rows, err := repo.db.Raw(`
		select sum(m.precio) as total 
		from factura f
		inner join factura_items fi on f.id=fi.factura_id
		inner join medicamento m on m.id=fi.medicamento_id
		where f.id=?`, factura_id).Rows()
	if err != nil {
		return err
	}
	for rows.Next() {
		_ = rows.Scan(&pago_total)
	}
	if pago_total == nil {
		return errors.New("el pago total es null")
	}
	rows, err = repo.db.Raw(`
		update factura
		set pago_total=?
		where id=?`, *pago_total, factura_id).Rows()
	if err != nil {
		return err
	}
	return nil
}

func (repo Repository) precio_medicamentos(id_medicamentos string) *float32 {
	var costo *float32
	sql := `select sum(m.precio) as costo 
		from medicamento m
		where m.id in ({id_medicamento})`
	sql = strings.ReplaceAll(sql, "{id_medicamento}", strings.Join(strings.Split(id_medicamentos, ","), ","))
	rows, err := repo.db.Raw(sql).Rows()
	if err != nil {
		return nil
	}
	for rows.Next() {
		_ = rows.Scan(&costo)
	}
	return costo
}
