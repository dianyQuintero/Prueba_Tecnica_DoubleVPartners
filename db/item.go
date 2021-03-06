package db
import (
    "database/sql"
    "github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners/models"
)
func (db Database) GetAllItems() (*models.ItemList, error) {
    list := &models.ItemList{}
    rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
    if err != nil {
        return list, err
    }
    for rows.Next() {
		var item models.Item
        err := rows.Scan(&item.ID, &item.Usuario, &item.FechaCreacion, &item.FechaActualizacion, &item.Estado_Abierto)
        if err != nil {
            return list, err
        }
        list.Items = append(list.Items, item)
    }
    return list, nil
}
func (db Database) AddItem(item *models.Item) error {
    var id int
	var fecha_creacion string
	var estado_abierto string
	var fecha_actualizacion string
    query := `INSERT INTO items (usuario) VALUES ($1) RETURNING id, fecha_creacion, fecha_actualizacion, estado_abierto`
    err := db.Conn.QueryRow(query, item.Usuario).Scan(&id, &fecha_creacion, &fecha_actualizacion , &estado_abierto)
    if err != nil {
        return err
    }
    item.ID = id
	item.FechaCreacion = fecha_creacion
	item.FechaActualizacion= fecha_actualizacion
	item.Estado_Abierto = estado_abierto
    return nil
}
func (db Database) GetItemById(itemId int) (models.Item, error) {
    item := models.Item{}
    query := `SELECT * FROM items WHERE id = $1;`
    row := db.Conn.QueryRow(query, itemId)
    switch err := row.Scan(&item.ID, &item.Usuario, &item.FechaCreacion, &item.FechaActualizacion, &item.Estado_Abierto); err {
    case sql.ErrNoRows:
        return item, ErrNoMatch
    default:
        return item, err
    }
}
func (db Database) DeleteItem(itemId int) error {
    query := `DELETE FROM items WHERE id = $1;`
    _, err := db.Conn.Exec(query, itemId)
    switch err {
    case sql.ErrNoRows:
        return ErrNoMatch
    default:
        return err
    }
}
func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
    item := models.Item{}
    query := `UPDATE items SET usuario=$1, fecha_actualizacion=$2, estado_abierto=$3 WHERE id=$4 RETURNING id, usuario, fecha_creacion, fecha_actualizacion, estado_abierto;`
    err := db.Conn.QueryRow(query, itemData.Usuario, itemData.FechaActualizacion, itemData.Estado_Abierto, itemId).Scan(&item.ID, &item.Usuario, &item.FechaCreacion, &item.FechaActualizacion, &item.Estado_Abierto)
    if err != nil {
        if err == sql.ErrNoRows {
            return item, ErrNoMatch
        }
        return item, err
    }
    return item, nil
}