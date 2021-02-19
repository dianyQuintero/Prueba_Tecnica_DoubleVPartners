package models
import (
    "fmt"
    "net/http"
)
type Item struct {
    ID int `json:"id"`
    Usuario string `json:"usuario"`
    FechaCreacion string `json:"fecha_creacion"`
	FechaActualizacion string `json:"fecha_actualizacion"`
	Estado_Abierto string `json:"estado_abierto"`
}
type ItemList struct {
    Items []Item `json:"items"`
}
func (i *Item) Bind(r *http.Request) error {
    if i.Usuario == "" {
        return fmt.Errorf("por favor proporcione un usuario")
    }
    return nil
}
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}