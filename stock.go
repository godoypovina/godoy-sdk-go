package godoy

import "strconv"

// Stock ...
type Stock struct {
	Stock struct {
		Code      string `gorm:"column:codigo" json:"code"`
		Store     int    `gorm:"column:deposito" json:"store"`
		Available int    `gorm:"column:disponible" json:"disponible"`
	} `json:"stock,omitempty"`
}

// GetStock ...
func (g *Godoy) GetStock(code string, store int) (*Stock, error) {
	stock := Stock{}

	err := g.get(`/stock/`+code+"/"+strconv.Itoa(store), nil, &stock)
	if err != nil {
		return nil, err
	}

	return &stock, nil
}
