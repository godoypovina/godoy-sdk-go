package godoy

import "net/url"

// Remito ...
type PedidoVentaDetalle struct {
	Detalles []struct {
		Codigo          string `json:"codigo"`
		Nombre          string `json:"nombre"`
		Cantidad        string `json:"cantidad"`
		CantidadResiduo string `json:"cantidad_residuo"`
		Marca           string `json:"marca"`
		Posicion        string `json:"posicion"`
		NumeroPedido    string `json:"numero_pedido"`
		Residuo         string `json:"residuo"`
	} `json:"detalles"`
	Total int `json:"total"`
}

// GetPedidoVentaDetalle ...
func (g *Godoy) GetPedidoVentaDetalle(number string, query url.Values) (*PedidoVentaDetalle, error) {
	detalles := PedidoVentaDetalle{}

	err := g.get("/pedidos/"+number+"/detalle", query, &detalles)
	if err != nil {
		return nil, err
	}

	return &detalles, nil
}

func (g *Godoy) GetPedidoVentaDetalleUnigestion(number string, query url.Values) (*PedidoVentaDetalle, error) {
	detalles := PedidoVentaDetalle{}

	err := g.get("/uni/pedidos/"+number+"/detalle", query, &detalles)
	if err != nil {
		return nil, err
	}

	return &detalles, nil
}
