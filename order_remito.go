package godoy

import "net/url"

// Data ...
type Data struct {
	Remitos []struct {
		OrderNote     string `json:"order_note"`
		RemitoNumber  string `json:"remito_number"`
		InvoiceNumber string `json:"invoice_number"`
		Quantity      int    `json:"quantity"`
		Delivered     int    `json:"delivered"`
		Pending       int    `json:"pending"`
	} `json:"remitos"`
}

// GetAllOrderRemitos ...
func (g *Godoy) GetAllOrderRemitos(query url.Values) (*Data, error) {
	remitos := Data{}

	err := g.get("/order/remitos", query, &remitos)
	if err != nil {
		return nil, err
	}

	return &remitos, nil
}
