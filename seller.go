package godoy

import "net/url"

// Seller ...
type Seller struct {
	Sellers []struct {
		Code  string `json:"codigo"`
		Name  string `json:"nombre"`
		Email string `json:"email"`
	} `json:"vendedores"`
	Total int `json:"total"`
}

// GetAllSellers returns list of sellers
func (g *Godoy) GetAllSellers(query url.Values) (*Seller, error) {
	sellers := Seller{}

	err := g.get("/vendedores", query, &sellers)
	if err != nil {
		return nil, err
	}

	return &sellers, nil
}
