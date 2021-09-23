package repository

import (
	"fmt"

	"jesus.tn79/aveonline/model"
)

func (repo Repository) GetPromociones() ([]*model.Promocion, error) {
	var promociones []*model.Promocion
	rows, err := repo.db.Raw("SELECT id, descripcion, porcentaje, fecha_inicio, fecha_fin FROM promocion").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		promocion := &model.Promocion{}
		_ = rows.Scan(&promocion.ID, &promocion.Descripcion, &promocion.Porcentaje, &promocion.Fecha_Inicio, &promocion.Fecha_Fin)
		promociones = append(promociones, promocion)
	}
	return promociones, err
}

func (repo Repository) GetPromocion(promocion_id int) (*model.Promocion, error) {
	rows, err := repo.db.Raw(`SELECT id, descripcion, porcentaje, fecha_inicio, fecha_fin FROM promocion WHERE id=?`, promocion_id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	promocion := &model.Promocion{}
	for rows.Next() {
		_ = rows.Scan(&promocion.ID, &promocion.Descripcion, &promocion.Porcentaje, &promocion.Fecha_Fin, &promocion.Fecha_Fin)
	}
	return promocion, err
}

func (repo Repository) GetPromocionPorFecha(fecha string) (*model.Promocion, error) {
	rows, err := repo.db.Raw(`SELECT id, descripcion, porcentaje, fecha_inicio, fecha_fin FROM promocion WHERE ? BETWEEN fecha_inicio and fecha_fin`, fecha).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	promocion := &model.Promocion{}
	for rows.Next() {
		_ = rows.Scan(&promocion.ID, &promocion.Descripcion, &promocion.Porcentaje, &promocion.Fecha_Fin, &promocion.Fecha_Fin)
	}
	return promocion, err
}

func (repo Repository) CreatePromocion(data model.Promocion) (*string, error) {
	tx := repo.db.Exec("INSERT INTO promocion(descripcion,porcentaje,fecha_inicio,fecha_fin)VALUES(?,?,?,?)", data.Descripcion, data.Porcentaje, data.Fecha_Inicio, data.Fecha_Fin)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	message := "promoción creado."
	return &message, nil
}
