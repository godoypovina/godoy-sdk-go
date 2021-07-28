package godoy

import "net/url"

// Remito ...
type ResponseRemitoDetalle struct {
	Detalles []struct {
		Codigo       string `json:"codigo"`
		Nombre       string `json:"nombre"`
		Cantidad     string `json:"cantidad"`
		NumeroPedido string `json:"numero_pedido"`
		Posicion     string `json:"posicion"`
	} `json:"detalles"`
	Total int `json:"total"`
}

// GetDetail ...
func (g *Godoy) GetRemitoDetalle(number string, query url.Values) (*ResponseRemitoDetalle, error) {
	detalles := ResponseRemitoDetalle{}

	err := g.get("/remitos/"+number+"/detalle", query, &detalles)
	if err != nil {
		return nil, err
	}

	return &detalles, nil
}
