package godoy

import "fmt"

type ProtheusPedidoVentaDetalle struct {
	Nro        int     `json:"nro"`
	Cod        string  `json:"cod"`
	Cant       int     `json:"cant"`
	Dto        float64 `json:"dto"`
	CliID      string  `json:"cli_id"`
	CliLj      string  `json:"cli_lj"`
	Fecha      string  `json:"fecha"`
	Cond       string  `json:"cond"`
	Entrega    string  `json:"entrega"`
	Vendedor   string  `json:"vendedor"`
	TipoRemito string  `json:"tipoRemito"`
	TipoVenta  string  `json:"tipoVenta"`
	Precio     float64 `json:"precio"`
	Deposito   int     `json:"deposito"`
}

// RequestNewPedido ...
type RequestNewPedidoProtheus struct {
	Data []ProtheusPedidoVentaDetalle `json:"data"`
}

// NewPedidoResponse ...
type NewPedidoProtheusResponse struct {
	Data struct {
		Pedidos []string `json:"pedidos"`
	} `json:"data"`
	Endpoint string `json:"endpoint"`
}

// GetStock ...
func (g *Godoy) CreatePedido(data interface{}) (*NewPedidoProtheusResponse, error) {
	response := NewPedidoProtheusResponse{}

	err := g.post("/pedidos/venta", data, &response)
	if err != nil {
		return nil, err
	}

	fmt.Println(response)

	return &response, nil
}
