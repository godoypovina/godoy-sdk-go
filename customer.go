package godoy

import "net/url"

// Customers ...
type Customers struct {
	Customers []struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		ProtheusCode string `json:"protheus_code"`
		Direccion    string `json:"address"`
		City         string `json:"city"`
		ZipCode      string `json:"zip_code"`
		State        string `json:"state"`
		Cuit         string `json:"cuit"`
		Code         string `json:"code"`
		SellerName   string `json:"seller_name"`
	} `json:"customers"`
	Total int `json:"total"`
}

// GetAllCustomers returns list of customers
func (g *Godoy) GetAllCustomers(query url.Values) (*Customers, error) {
	customers := Customers{}

	err := g.get("/customers", query, &customers)
	if err != nil {
		return nil, err
	}

	return &customers, nil
}
